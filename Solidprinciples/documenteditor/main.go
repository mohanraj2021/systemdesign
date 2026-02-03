package main

type DocumentElement interface {
	render() string
}

type textElement struct {
	text string
}

func (t *textElement) render() string {
	return t.text
}

type imageElement struct {
	url string
}

func (i *imageElement) render() string {
	return "<img src='" + i.url + "' />"
}

type Document struct {
	elements []DocumentElement
}

type DocumentEditor struct {
	Doc Document
}

func (d *DocumentEditor) AddText(text string) {
	d.Doc.elements = append(d.Doc.elements, &textElement{text: text})
}

func (d *DocumentEditor) AddImage(url string) {
	d.Doc.elements = append(d.Doc.elements, &imageElement{url: url})
}

func (d *DocumentEditor) RenderDocument() string {
	result := ""
	for _, element := range d.Doc.elements {
		result += element.render() + "\n"
	}
	return result
}

func main() {
	editor := DocumentEditor{}
	editor.AddText("Hello")
	editor.AddImage("http://example.com/image.jpg")
	editor.AddText("World")
	println(editor.RenderDocument())
}
