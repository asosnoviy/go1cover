package reportgen

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strconv"

	coveragdata "github.com/asosnoviy/go1cover/internal/coverageData"
)

type Generic struct {
	Reporter
}

func NewGeneric(path string) *Generic {

	return &Generic{Reporter{Path: path}}
}

func (r *Generic) Generate(coverage *coveragdata.Coverage) {
	fmt.Println("Generic generate", r.Path)

	files := []file{}
	for cfile, clines := range coverage.Data {

		lines := []lineToCover{}
		for line, covered := range clines {
			lines = append(lines, lineToCover{LineNumber: strconv.Itoa(line), Covered: covered})
		}
		files = append(files, file{File: cfile, LineToCover: lines})
	}

	v := &generCoverage{Version: "1", Files: files}

	xmlfile, _ := os.Create(r.Path)
	defer xmlfile.Close()

	xmlWriter := io.Writer(xmlfile)

	enc := xml.NewEncoder(xmlWriter)
	enc.Indent("", "  ")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}

}

// Exaple:
// <coverage version="1">
//   <file path="xources/hello/NoConditions.xoo">
//     <lineToCover lineNumber="6" covered="true"/>
//     <lineToCover lineNumber="7" covered="false"/>
//   </file>
//   <file path="xources/hello/WithConditions.xoo">
//     <lineToCover lineNumber="3" covered="true" branchesToCover="2" coveredBranches="1"/>
//   </file>
// </coverage>

type generCoverage struct {
	XMLName xml.Name `xml:"coverage"`
	Version string   `xml:"version,attr"`
	Files   []file
}

type file struct {
	XMLName     xml.Name `xml:"file"`
	File        string   `xml:"path,attr"`
	LineToCover []lineToCover
}

type lineToCover struct {
	XMLName    xml.Name `xml:"lineToCover"`
	LineNumber string   `xml:"lineNumber,attr"`
	Covered    bool     `xml:"covered,attr"`
}
