package reportgen

import (
	"fmt"
	"log"
	"os"
	"strconv"

	coveragdata "github.com/asosnoviy/go1cover/internal/coverageData"
)

type Lcov struct {
	Reporter
}

func NewLcov(path string) *Lcov {

	return &Lcov{Reporter{Path: path}}
}

func (r *Lcov) Generate(coverage *coveragdata.Coverage) {
	fmt.Println("Lcov generate", r.Path)

	f, err := os.Create(r.Path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	for filename, lines := range coverage.Data {
		f.WriteString("TN:" + "\n")
		f.WriteString("SF:" + filename + "\n")

		for k, v := range lines {
			covered := "0"
			if v {
				covered = "1"
			}
			f.WriteString("DA:" + strconv.Itoa(k) + "," + covered + "\n")
		}

		countlines := strconv.Itoa(len(lines))
		f.WriteString("LH:0" + "\n")
		f.WriteString("LF:" + countlines + "\n")
		f.WriteString("end_of_record" + "\n")
	}
	f.WriteString("\n")
}
