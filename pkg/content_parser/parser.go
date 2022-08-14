package content_parser

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"
)

type BoxText struct {
	Box []TextStruct
}

type TextStruct struct {
	Tag, Text string
}

type ContentContainer []*BoxText

type Parser interface {
	ParseContent(r io.Reader) error
}

func NewParser() *ContentContainer {
	return &ContentContainer{}
}

func tagChecker(tag string) bool {
	switch tag {
	case "p", "a", "div", "ul", "ol", "li", "code", "h1", "h2", "h3", "h4",
		"h5", "h6", "b", "i", "strong", "span", "section", "header", "label":
		return true
	}
	return false
}

func (c *ContentContainer) ParseContent(r io.Reader) error {
	content, err := html.Parse(r)
	if err != nil {
		return err
	}

	var newBox *BoxText
	var parse func(n *html.Node)
	parse = func(n *html.Node) {
		if n.Type == html.ElementNode && tagChecker(n.Data) && n.Data != "a" {
			if newBox = getContent(n, n.Data); len(newBox.Box) > 0 {
				*c = append(*c, newBox)
			}
		}
		if n.Data == "nav" || n.Data == "footer" {
			return
		}
	
		if n.FirstChild != nil {
			parse(n.FirstChild)
		}

		if n.NextSibling != nil {
			parse(n.NextSibling)
		}
	}
	parse(content)

	for i, a := range *c {
		fmt.Println("#", i, " ", a.Box)
	}
	fmt.Println("total ", len(*c))

	return nil
}

func getContent(n *html.Node, tag string) *BoxText {
	newBoxContainer := &BoxText{}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode {
			if c.FirstChild != nil {
				if !tagChecker(c.Data) {
					break
				}
				tag = c.Data
				newBoxContainer = mergeContent(newBoxContainer, getContent(c, tag))
			}
		} else if c.Type == html.TextNode {
			if strings.TrimSpace(c.Data) == "" {
				continue
			}
			newBoxContainer.Box = append(
				newBoxContainer.Box,
				TextStruct{
					Tag:  tag,
					Text: strings.TrimSuffix(c.Data, "\r\t\n"),
				},
			)
		}
	}

	return newBoxContainer
}

func mergeContent(first, second *BoxText) *BoxText {
	if first == nil || second == nil {
		return nil
	}

	for _, b := range second.Box {
		first.Box = append(first.Box, TextStruct{
			b.Tag,
			b.Text,
		})
	}

	return first
}
