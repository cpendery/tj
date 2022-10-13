package csv

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
)

func Encode(blob []byte) ([]byte, error) {
	reader := csv.NewReader(bytes.NewReader(blob))
	rows := [][]string{}
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("unable to parse csv: %w", err)
		}
		rows = append(rows, line)
	}

	output, err := json.Marshal(rows)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal csv to json")
	}
	return output, nil
}
