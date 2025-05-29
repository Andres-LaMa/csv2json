package parser

import (
	"strings"
	"testing"
)

func TestParseJSON(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  map[string]interface{}
	}{
		{
			name:  "simple json",
			input: `{"name":"Alice","age":25}`,
			want:  map[string]interface{}{"name": "Alice", "age": 25.0},
		},
		// {
		// 	name:  "nested json",
		// 	input: `{"user":{"name":"Bob"}}`,
		// 	want:  map[string]interface{}{"user": map[string]interface{}{"name": "Bob"}},
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch, err := ParseJSON(strings.NewReader(tt.input))
			if err != nil {
				t.Fatal(err)
			}
			got := <-ch
			for k, v := range tt.want {
				if got[k] != v {
					t.Errorf("Field %s = %v, want %v", k, got[k], v)
				}
			}
		})
	}
}
