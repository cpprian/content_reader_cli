package testing

import (
	con "github.com/cpprian/content_reader_cli/pkg/content_parser"
)

var (
	want_test_1 = con.BoxText{
		{
			Tag:  "h1",
			Text: "Hello world",
		},
		{
			Tag:  "p",
			Text: "a brand new text, to learn some new things",
		},
		{
			Tag:  "p",
			Text: "hello world new dir",
		},
		{
			Tag:  "p",
			Text: "Sick content",
		},
		{
			Tag:  "a",
			Text: "it should also appears",
		},
	}
)
