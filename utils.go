package main

import (
	"encoding/csv"
	"os"
)

/*func sortTsv(data [][]string) {
	sort.Slice(data[:], func(i, j int) bool {
		for x := range data[i] {
			if data[i][x] == data[j][x] {
				continue
			}
			return data[i][x] < data[j][x]
		}
		return false
	})
}*/

// TODO: just return the file and let caller handle
// this will allow the caller to create a new reader
// after seeking to the desired position
func OpenTsv(path string) (*csv.Reader, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(f)
	r.LazyQuotes = true
	r.FieldsPerRecord = -1
	r.Comma = '\t'

	return r, nil
}
