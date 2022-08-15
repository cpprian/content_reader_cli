package testing

import (
	con "github.com/cpprian/content_reader_cli/pkg/content_parser"
)

var (
	want_test_1 = &con.ContentContainer{
		&con.BoxText{
			Box: []con.TextStruct{
				{
					Tag:  "h1",
					Text: "Hello world",
				},
				{
					Tag:  "p",
					Text: "a brand new text, to learn some new things",
				},
				{
					Tag:  "div",
					Text: "ok, why not try to learn some new things",
				},
				{
					Tag:  "a",
					Text: "I know some things!",
				},
				{
					Tag:  "div",
					Text: "hello world",
				},
				{
					Tag:  "i",
					Text: "this is an amazing",
				},
				{
					Tag:  "div",
					Text: " text ",
				},
				{
					Tag:  "b",
					Text: "to learn",
				},
				{
					Tag:  "div",
					Text: " some new things",
				},
				{
					Tag:  "div",
					Text: "ok, that was strange",
				},
			},
		},
		&con.BoxText{
			Box: []con.TextStruct{
				{
					Tag:  "div",
					Text: "some div",
				},
				{
					Tag:  "p",
					Text: "hello world new dir",
				},
			},
		},
		&con.BoxText{
			Box: []con.TextStruct{
				{
					Tag:  "p",
					Text: "Sick content",
				},
			},
		},
	}
)
