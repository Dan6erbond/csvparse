package csvparse_test

import (
	"encoding/csv"
	"fmt"
	"strings"
	"testing"

	"github.com/Dan6erbond/csvparse"
	"github.com/stretchr/testify/assert"
)

func TestScanReaderWithHeaderRow_Scan(t *testing.T) {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	r := csv.NewReader(strings.NewReader(in))

	sr := csvparse.NewScanReader(r, csvparse.WithHeaderRow)

	var (
		fn string
		ln string
		un string
	)

	err := sr.Scan(&fn, &ln, &un)

	assert.NoError(t, err)

	assert.Equal(t, "Rob", fn)
	assert.Equal(t, "Pike", ln)
	assert.Equal(t, "rob", un)
}

func ExampleScanReader_Scan() {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	r := csv.NewReader(strings.NewReader(in))

	sr := csvparse.NewScanReader(r, csvparse.WithHeaderRow)

	var (
		fn string
		ln string
		un string
	)

	sr.Scan(&fn, &ln, &un)

	fmt.Println(fn, ln, un)

	// Output: Rob Pike rob
}

func TestScanReader_Scan_WithTooManyTargets(t *testing.T) {
	in := `first_name,last_name,username
"Rob","Pike",rob
Ken,Thompson,ken
"Robert","Griesemer","gri"
`
	r := csv.NewReader(strings.NewReader(in))

	sr := csvparse.NewScanReader(r, csvparse.WithHeaderRow)

	var (
		fn string
		ln string
		un string
		pw string
	)

	err := sr.Scan(&fn, &ln, &un, &pw)

	if assert.Error(t, err) {
		assert.Equal(t, csvparse.ErrRowCountIsSmallerThanTargetCount, err)
	}
}
