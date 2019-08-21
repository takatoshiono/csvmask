package mask

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
)

type (
	// Reader reads from a CSV data and mask it.
	Reader struct {
		r       *csv.Reader
		rules   Rules
		lineNum int

		// If SkipHeader is true, the first line is skipped masking as header.
		SkipHeader bool
	}
)

// NewReader returns a new reader that reads from r while masking data according to the rule.
func NewReader(r io.Reader, rules Rules) *Reader {
	return &Reader{
		r:     csv.NewReader(r),
		rules: rules,
	}
}

// Read reads one record from r.r while masking and return it as a CSV string.
// If there is no data left to be read, Read returns "", io.EOF.
func (r *Reader) Read() (string, error) {
	record, err := r.r.Read()
	if err != nil {
		return "", err
	}
	r.lineNum++

	if r.lineNum == 1 && r.SkipHeader {
		s, err := toCSV(record)
		if err != nil {
			return "", err
		}
		return s, nil
	}

	if len(record) != len(r.rules) {
		return "", fmt.Errorf("the number of rules is %d but must same as csv fields %d", len(r.rules), len(record))
	}

	for i, field := range record {
		record[i], err = r.rules[i].Convert(field)
		if err != nil {
			return "", fmt.Errorf("failed to convert field %d at line %d: %v", i, r.lineNum, err)
		}
	}

	s, err := toCSV(record)
	if err != nil {
		return "", err
	}

	return s, nil
}

func toCSV(record []string) (string, error) {
	var buf bytes.Buffer
	w := csv.NewWriter(&buf)
	if err := w.Write(record); err != nil {
		return "", fmt.Errorf("failed to write record as csv: %v", err)
	}
	w.Flush()
	if err := w.Error(); err != nil {
		return "", fmt.Errorf("failed to flush record as csv: %v", err)
	}
	return buf.String(), nil
}
