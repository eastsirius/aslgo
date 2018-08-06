// 页面类型
package agui

import (
	"io"
	"bufio"
)

type Page interface {
	Control
	WritePage(writer io.Writer) error
}

type Html5Page struct {
	Title string
}

func NewHtml5Page(title string) Page {
	return &Html5Page{
		Title: title,
	}
}

func (page *Html5Page) WritePage(writer io.Writer) error {
	bufio_writer := bufio.NewWriter(writer)
	pageWriter := NewTextWriter(bufio_writer)
	page.Write(pageWriter)
	bufio_writer.Flush()

	return nil
}

func (page *Html5Page) Write(writer Writer) {
	writer.WriteLine("<!DOCTYPE html>")
	writer.WriteLine("<html>")

	writer.IncIndent()
	page.writeHead(writer)
	page.writeBody(writer)
	writer.DecIndent()

	writer.WriteLine("</html>")
}

func (page *Html5Page) writeHead(writer Writer) {
	writer.WriteLine("<head>")

	writer.IncIndent()
	writer.WriteLine("<meta charset=\"utf-8\">")
	writer.WriteLine("<title>" + page.Title + "</title>")
	writer.DecIndent()

	writer.WriteLine("</head>")
}

func (page *Html5Page) writeBody(writer Writer) {
	writer.WriteLine("<body>")

	writer.IncIndent()
	writer.DecIndent()

	writer.WriteLine("</body>")
}
