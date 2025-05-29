package parser

import (
	"strings"
	"testing"
)

func TestParseCSV(t *testing.T) {
	input := `name,age
Alice,25`

	rows, err := ParseCSV(strings.NewReader(input))
	if err != nil {
		t.Fatal(err)
	}

	row := <-rows
	if row["name"] != "Alice" || row["age"] != 25 {
		t.Errorf("Parsing failed: %v", row)
	}
}
