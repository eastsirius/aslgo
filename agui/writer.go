// 写入器
package agui

import (
	"io"
	"fmt"
)

type Writer interface {
	WriteString(s string)
	Write(a ...interface{})
	WriteIndent()
	WriteReturn()
	WriteLine(a ...interface{})
	IncIndent()
	DecIndent()
}

type TextWriter struct {
	Writer io.Writer
	Indent int
	IndentString string
	ReturnString string
}

func NewTextWriter(writer io.Writer) Writer {
	return &TextWriter{
		Writer: writer,
		Indent: 0,
		IndentString: "\t",
		ReturnString: "\r\n",
	}
}

func (writer *TextWriter) WriteString(s string) {
	writer.Writer.Write([]byte(s))
}

func (writer *TextWriter) Write(a ...interface{}) {
	str := fmt.Sprint(a...)
	writer.WriteString(str)
}

func (writer *TextWriter) WriteIndent() {
	if writer.Indent == 0 {
		return
	}

	str := ""
	for i := 0; i < writer.Indent; i++ {
		str += writer.IndentString
	}

	writer.WriteString(str)
}

func (writer *TextWriter) WriteReturn() {
	writer.WriteString(writer.ReturnString)
}

func (writer *TextWriter) WriteLine(a ...interface{}) {
	writer.WriteIndent()
	writer.Write(a...)
	writer.WriteReturn()
}

func (writer *TextWriter) IncIndent() {
	writer.Indent++
}

func (writer *TextWriter) DecIndent() {
	if writer.Indent > 0 {
		writer.Indent--
	}
}
