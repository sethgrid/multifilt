package multifilt

import (
	"bufio"
	"bytes"
	"io"
)

// Filter reads through the src line by line, and passes through lines that don't match any line in the filter
func Filter(src io.Reader, filter io.Reader, output io.Writer, requireFullMatch bool) error {
	var err error
	sIn := bufio.NewScanner(src)
	sFlt := bufio.NewScanner(filter)
	var filters [][]byte
	for sFlt.Scan() {
		var alloc []byte
		alloc = append(alloc, sFlt.Bytes()...)
		filters = append(filters, alloc)
	}

	for sIn.Scan() {
		var lineShouldFilter bool

		for _, b := range filters {
			if isMatch(sIn.Bytes(), b, requireFullMatch) {
				lineShouldFilter = true
				break
			}
		}

		if lineShouldFilter {
			continue
		}

		_, err = output.Write(sIn.Bytes())
		if err != nil {
			return err
		}
		_, err = output.Write([]byte("\n"))
		if err != nil {
			return err
		}
	}
	return nil
}

func isMatch(a []byte, b []byte, requireFullMatch bool) bool {
	if !requireFullMatch && bytes.Contains(a, b) {
		return true
	}
	return bytes.Compare(a, b) == 0
}
