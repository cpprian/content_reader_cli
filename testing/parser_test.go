package testing

import (
	"fmt"
	"os"
	"testing"

	con "github.com/cpprian/content_reader_cli/pkg/content_parser"
)

func TestParser(t *testing.T) {

	t.Run("a basic page with content in div", func(t *testing.T) {
		got := con.NewParser()

		f, _ := os.Open("./tested_files/test_1.html")

		got.CreateBoxText(f)
		fmt.Println(got)
		CompareBoxText(t, got, &want_test_1)
	})
}

func CompareBoxText(t *testing.T, got *con.BoxText, want *con.BoxText) {

	if len(*got) != len(*want) {
		t.Fatalf("got %d want %d", len(*got), len(*want))
	}

	for i := range *got {
		CompareTextStruct(t, &(*got)[i], &(*want)[i])
	}
}

func CompareTextStruct(t *testing.T, got *con.TextStruct, want *con.TextStruct) {

	if got.Tag != want.Tag {
		t.Fatalf("got %s want %s", got.Tag, want.Tag)
	}
	if got.Text != want.Text {
		t.Fatalf("got %s want %s", got.Text, want.Text)
	}
	if len(got.Children) != len(want.Children) {
		t.Fatalf("got %d want %d", len(got.Children), len(want.Children))
	}

	for i := range got.Children {
		CompareTextStruct(t, &got.Children[i], &want.Children[i])
	}
}
