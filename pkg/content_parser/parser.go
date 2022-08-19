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
		return true
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
	node, err := getBody(n)
	if err != nil {
		return
	}

	var saveContent func(n *html.Node)
	saveContent = func(n *html.Node) {
		if n.Type == html.ElementNode && tagChecker(n.Data) {
			*b = append(*b, *ParseNode(n))
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			saveContent(c)
		}
	}
	saveContent(node)
}

func getBody(n *html.Node) (*html.Node, error) {
	var body *html.Node

	var searchForBody func(n *html.Node)
	searchForBody = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "body" {
			body = n
			return
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			searchForBody(c)
		}
	}
	searchForBody(n)

	if body != nil {
		return body, nil
	}
	return nil, fmt.Errorf("cannot find body")
}

func ParseNode(n *html.Node) *TextStruct {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			if len(strings.TrimSpace(c.Data)) > 1 {
				fmt.Println(strings.TrimSpace(c.Data))
			}	
		}
	}

	return &TextStruct{}
}

func ParseTextStruct(n *html.Node, tag string) *TextStruct {

	return &TextStruct{}
}

func (b *BoxText) String() string {
	return fmt.Sprintf("%v", *b)
}
