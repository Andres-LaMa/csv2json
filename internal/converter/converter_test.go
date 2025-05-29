package converter

import (
	"bytes"
	"testing"
)

func TestCSVToJSON(t *testing.T) {
	input := `name,age
Alice,25
Bob,30`

	var buf bytes.Buffer
	err := CSVToJSON(bytes.NewReader([]byte(input)), &buf)
	if err != nil {
		t.Fatalf("CSVToJSON failed: %v", err)
	}

	expected := `[
  {
    "name": "Alice",
    "age": 25
  },
  {
    "name": "Bob",
    "age": 30
  }
]`
	if buf.String() != expected {
		t.Errorf("Expected:\n%s\nGot:\n%s", expected, buf.String())
	}
}
