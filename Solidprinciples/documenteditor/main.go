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

func NewTextElement(text string) DocumentElement {
	return &textElement{text: text}
}

type imageElement struct {
	url string
}

func (i *imageElement) render() string {
	return "<img src='" + i.url + "' />"
}

func NewImageElement(url string) DocumentElement {
	return &imageElement{url: url}
}

type Document struct {
	elements []DocumentElement
}

type DocumentEditor struct {
	Doc Document
}

func (d *DocumentEditor) AddElement(element DocumentElement) {
	d.Doc.elements = append(d.Doc.elements, element)
}

func (d *DocumentEditor) RenderDocument() string {
	result := ""
	for _, element := range d.Doc.elements {
		result += element.render() + "\n"
	}
	return result
}

func main() {
	editor := &DocumentEditor{}
	editor.AddElement(NewTextElement("Hello"))
	editor.AddElement(NewImageElement("http://example.com/image.jpg"))
	editor.AddElement(NewTextElement("World"))
	println(editor.RenderDocument())
}
