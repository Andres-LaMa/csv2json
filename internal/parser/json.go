package parser

import (
	"encoding/json"
	"io"
	"log"
)

// ParseJSON читает JSON потоково (по одному объекту).
func ParseJSON(r io.Reader) (<-chan map[string]interface{}, error) {
	decoder := json.NewDecoder(r)
	rows := make(chan map[string]interface{})

	go func() {
		defer close(rows)

		for {
			var row map[string]interface{}
			err := decoder.Decode(&row)
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("[ERROR] JSON decode: %v", err)
				continue
			}
			rows <- row
		}
	}()

	return rows, nil
}
