package content_parser

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

type TextStruct struct {
	Tag, Text string
	Children  []TextStruct
}

type BoxText []TextStruct

func NewParser() *BoxText {
	return &BoxText{}
}

func (b *BoxText) CreateBoxText(r io.Reader) error {
	doc, err := html.Parse(r)
	if err != nil {
		return err
	}

	b.Parse(doc)

	return nil
}

func (b *BoxText) Parse(n *html.Node) {

	f := func(n *html.Node) {
		fmt.Println("#", n.Data)
		*b = append(*b, ParseNode(n))

		if n.FirstChild != nil {
			b.Parse(n.FirstChild)
		}

		if n.NextSibling != nil {
			b.Parse(n.NextSibling)
		}
	}
	f(n)
}

func ParseNode(n *html.Node) TextStruct {
	if n.Type == html.ElementNode {
		fmt.Println("Tag: ", n.Data)
	} else if n.Type == html.TextNode {
		fmt.Println("Text: ", n.Data)
	}
	return TextStruct{}
}

func (b *BoxText) String() string {
	return fmt.Sprintf("%v", *b)
}
