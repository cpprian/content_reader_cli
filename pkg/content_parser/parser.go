package content_parser

import (
	"fmt"
	"io"
	"strings"

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

func tagChecker(tag string) bool {
	switch tag {
	case "div", "p", "h1", "a":
		return true;
	}
	return false
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
		data := ParseNode(n)
		if data != nil {
			*b = append(*b, *data)
		}

		if n.FirstChild != nil {
			b.Parse(n.FirstChild)
		}

		if n.NextSibling != nil {
			b.Parse(n.NextSibling)
		}
	}
	f(n)
}

func ParseNode(n *html.Node) *TextStruct {
	if n.Type == html.ElementNode && tagChecker(n.Data) {
		if n.FirstChild != nil {
			return ParseTextStruct(n.FirstChild, n.Data)
		}
	}
	return nil
}

func ParseTextStruct(n *html.Node, tag string) *TextStruct {
	if n.Type == html.TextNode {
		data := strings.TrimSpace(n.Data)
		if len(data) < 1 {
			return nil
		}
		return &TextStruct{Tag: tag, Text: n.Data}
	}
	return nil
}

func (b *BoxText) String() string {
	return fmt.Sprintf("%v", *b)
}
