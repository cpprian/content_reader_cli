package testing

import (
	con "github.com/cpprian/content_reader_cli/pkg/content_parser"
)

var (
	want_test_1 = con.BoxText{
		{
			Tag:  "div",
			Text: "",
			Children: []con.TextStruct{
				{
					Tag:      "h1",
					Text:     "Hello world",
					Children: nil,
				},
				{
					Tag:      "p",
					Text:     "a brand new text, to learn some new things",
					Children: nil,
				},
			},
		},
		{
			Tag:  "div",
			Text: "",
			Children: []con.TextStruct{
				{
					Tag:      "p",
					Text:     "hello world new dir",
					Children: nil,
				},
			},
		},
		{
			Tag:  "p",
			Text: "Sick content",
			Children: nil,
		},
	}
)
