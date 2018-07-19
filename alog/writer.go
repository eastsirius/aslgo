// 日志输出器
package alog

import (
	"fmt"
	"time"
	"os"
	"sync"
	"path/filepath"
	"io/ioutil"
	"strings"
)

type Writer interface {
	WriteLog(log *LogItem) error
	Start()
	Stop()
}


type ConsoleWriter struct {
}

func NewConsoleWriter() *ConsoleWriter {
	return &ConsoleWriter{}
}

func (writer *ConsoleWriter) WriteLog(log *LogItem) error {
	fmt.Println(log.Log)
	return nil
}

func (writer *ConsoleWriter) Start() {
}

func (writer *ConsoleWriter) Stop() {
}


type FileNamer interface {
	GetName(tm *time.Time) string
}

type fileInfo struct {
	path string
	size uint64
	modTime time.Time
}

type basicFileNamer struct {
}

func (fn *basicFileNamer) GetName(tm *time.Time) string {
	return tm.Format("2006-01-02") + ".txt"
}

type FileWriter struct {
	// 日志文件命名器
	Namer FileNamer

	// 日志目录
	Path string

	// 秒单位日志保存时间，0代表永久保存
	SavedTime uint64

	// 字节单位最大日志总大小，0代表不做限制
	MaxSize uint64

	// 秒单位文件刷新周期
	FlushTime uint64

	// 秒单位文件监控周期
	MonitorTimer uint64

	file *os.File
	fileName string
	fileLock sync.Locker
	quitCtrl, quitEvent chan int
}

func NewFileWriter() *FileWriter {
	writer := &FileWriter{}
	writer.Namer = &basicFileNamer{}
	writer.FlushTime = 1
	writer.MonitorTimer = 60
	writer.fileLock = &sync.Mutex{}
	writer.quitCtrl = make(chan int, 1)
	writer.quitEvent = make(chan int, 1)

	return writer
}

func (writer *FileWriter) WriteLog(log *LogItem) error {
	fileName := writer.Namer.GetName(&log.Time)

	writer.fileLock.Lock()
	defer writer.fileLock.Unlock()

	if fileName != writer.fileName {
		err := writer.reopenFile(fileName)
		if err != nil {
			return err
		}
	}
	_, err := writer.file.WriteString(log.Log + "\n")
	if err != nil {
		return err
	}

	if writer.FlushTime == 0 {
		if writer.file != nil {
			writer.file.Sync()
		}
	}

	return nil
}

func (writer *FileWriter) Start() {
	go writer.workerProc()
}

func (writer *FileWriter) Stop() {
	writer.quitCtrl <- 0
	<- writer.quitEvent
	writer.fileLock.Lock()
	writer.closeFile()
	writer.fileLock.Unlock()
}

func (writer *FileWriter) reopenFile(fileName string) error {
	filePath := writer.Path + "/" + fileName
	filePath = strings.Replace(filePath, "\\", "/", -1)
	index := strings.LastIndex(filePath, "/")
	if index > 0 {
		dirPath := filePath[:index]
		err := os.MkdirAll(dirPath, 0777)
		if err != nil {
			return err
		}
	}

	file, err := os.OpenFile(filePath, os.O_APPEND | os.O_CREATE | os.O_SYNC | os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	writer.file = file
	writer.fileName = fileName

	return nil
}

func (writer *FileWriter) closeFile() {
	if writer.file != nil {
		writer.file.Close()
		writer.file = nil
	}
	writer.fileName = ""
}

func (writer *FileWriter) workerProc() {
	if writer.FlushTime == 0 {
		writer.FlushTime = 0xFFFFFFFF
	}
	if writer.MonitorTimer == 0 {
		writer.MonitorTimer = 0xFFFFFFFF
	}

	flush_td := time.Duration(writer.FlushTime) * time.Second
	tick_td := time.Duration(writer.MonitorTimer) * time.Second
	flush_timer := time.NewTimer(flush_td)
	tick_tm := time.NewTimer(tick_td)

	for true {
		select {
			case <- writer.quitCtrl:
				flush_timer.Stop()
				tick_tm.Stop()
				writer.quitEvent <- 0
				return
			case <- flush_timer.C:
				flush_timer.Reset(flush_td)
				writer.fileLock.Lock()
				if writer.file != nil {
					writer.file.Sync()
				}
				writer.fileLock.Unlock()
			case t := <- tick_tm.C:
				tick_tm.Reset(tick_td)
				writer.tickProc(&t)
		}
	}
}

func (writer *FileWriter) tickProc(tm *time.Time) {
	writer.fileLock.Lock()
	defer writer.fileLock.Unlock()
	writer.closeFile()

	files, paths, err := writer.findFiles()
	if err != nil {
		return
	}

	files = writer.deleteTimeoutFiles(files, tm)
	files = writer.deleteOverLimitFile(files)
	writer.deleteEmptyDirs(paths)
}

func (writer *FileWriter) findFiles() (map[time.Time] *fileInfo, []string, error) {
	files := make(map[time.Time] *fileInfo)
	paths := make([]string, 0)

	err := filepath.Walk(writer.Path, func(path string, info os.FileInfo, err error) error {
		if err != nil || info == nil {
			return err
		}

		if info.IsDir() {
			paths = append(paths, path)
		} else {
			files[info.ModTime()] = &fileInfo{
				path: path,
				size: uint64(info.Size()),
				modTime: info.ModTime(),
			}
		}

		return nil
	})

	if err != nil {
		return nil, nil, err
	}

	return files, paths, nil
}

func (writer *FileWriter) deleteTimeoutFiles(files map[time.Time] *fileInfo, tm *time.Time) map[time.Time] *fileInfo {
	new_files := make(map[time.Time] *fileInfo)
	for k, v := range files {
		to := v.modTime.Add(time.Duration(writer.SavedTime) * time.Second)
		if to.Before(*tm) {
			os.Remove(v.path)
		} else {
			new_files[k] = v
		}
	}

	return new_files
}

func (writer *FileWriter) deleteOverLimitFile(files map[time.Time] *fileInfo) map[time.Time] *fileInfo {
	new_files := make(map[time.Time] *fileInfo)

	var size_cnt uint64 = 0
	for _, v := range files {
		size_cnt += v.size
	}

	for k, v := range files {
		if size_cnt > writer.MaxSize {
			os.Remove(v.path)
		} else {
			new_files[k] = v
		}
	}

	return new_files
}

func (writer *FileWriter) deleteEmptyDirs(dirs []string) {
	var deleteDir = true
	for deleteDir {
		deleteDir = false
		for _, v := range dirs {
			if v != "" {
				fs, _ := ioutil.ReadDir(v)
				if len(fs) == 0 {
					os.Remove(v)
					v = ""
					deleteDir = true
				}
			}
		}
	}
}
