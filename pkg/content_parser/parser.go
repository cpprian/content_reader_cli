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

type Parser interface {
	ParseContent(r io.Reader) error
}

func NewParser() *BoxText {
	return &BoxText{}
}

func (b *BoxText) ParseContent(r io.Reader) error {
	doc, err := html.Parse(r)
	if err != nil {
		return err
	}

	fmt.Println(doc)

	return nil
}
