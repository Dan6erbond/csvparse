package csvparse

import (
	"encoding/csv"
)

type ScanReader struct {
	reader           *csv.Reader
	headerRow        bool
	skippedHeaderRow bool
}

type ScanReaderOption = func(sr *ScanReader)

var (
	WithHeaderRow ScanReaderOption = func(sr *ScanReader) {
		sr.headerRow = true
	}
)

func (sr *ScanReader) Scan(tgts ...*string) error {
	if sr.headerRow && !sr.skippedHeaderRow {
		_, err := sr.reader.Read()

		if err != nil {
			return err
		}

		sr.skippedHeaderRow = true
	}

	row, err := sr.reader.Read()

	if err != nil {
		return err
	}

	if len(row) < len(tgts) {
		return ErrRowCountIsSmallerThanTargetCount
	}

	for i := range tgts {
		*tgts[i] = row[i]
	}

	return nil
}

func NewScanReader(reader *csv.Reader, opts ...ScanReaderOption) *ScanReader {
	sr := &ScanReader{reader: reader}

	for _, opt := range opts {
		opt(sr)
	}

	return sr
}
