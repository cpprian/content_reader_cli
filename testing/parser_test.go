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

		got.ParseContent(f)
		CompareContainer(t, got, want_test_1)
	})
}

func CompareContainer(t testing.TB, got, want *con.ContentContainer) {
	if len(*got) != len(*want) {
		t.Errorf("got %d, want %d", len(*got), len(*want))
		return
	}
	for i := range *got {
		if len((*got)[i].Box) != len((*want)[i].Box) {
			t.Errorf("got %d, want %d", len((*got)[i].Box), len((*want)[i].Box))
			return
		}
		for j := range (*got)[i].Box {
			if (*got)[i].Box[j].Tag != (*want)[i].Box[j].Tag {
				t.Errorf("got %s, want %s", (*got)[i].Box[j].Tag, (*want)[i].Box[j].Tag)
			}
			if (*got)[i].Box[j].Text != (*want)[i].Box[j].Text {
				fmt.Println(len((*got)[i].Box[j].Text), len((*want)[i].Box[j].Text))
				t.Errorf("got %s, want %s", (*got)[i].Box[j].Text, (*want)[i].Box[j].Text)
			}
		}
	}
}
