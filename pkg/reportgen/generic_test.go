package reportgen

import (
	"io/ioutil"
	"os"
	"testing"

	coveragdata "github.com/asosnoviy/go1cover/internal/coverageData"
	"github.com/stretchr/testify/assert"
)

func TestGeneric_Generate(t *testing.T) {

	reportpath := "./temp/test.xml"
	reporter := NewGeneric(reportpath)

	data := map[string]map[int]bool{
		"xources/hello/NoConditions.xoo": {
			6: true,
			7: false,
		},
		"xources/hello/WithConditions.xoo": {
			3: true,
		},
	}

	coverage := coveragdata.Coverage{Data: data}
	reporter.Generate(&coverage)
	defer os.Remove(reportpath)

	b, err := ioutil.ReadFile(reportpath)
	if err != nil {
		t.Fail()
	}
	s := string(b)

	assert.Equal(t, exampleGeneric(), s)

}

func TestLcov_Generate(t *testing.T) {

	reportpath := "./temp/lcov.info"
	reporter := NewLcov(reportpath)

	data := map[string]map[int]bool{
		"xources/hello/NoConditions.xoo": {
			6: true,
			7: false,
		},
		"xources/hello/WithConditions.xoo": {
			3: true,
		},
	}

	coverage := coveragdata.Coverage{Data: data}
	reporter.Generate(&coverage)
	defer os.Remove(reportpath)

	b, err := ioutil.ReadFile(reportpath)
	if err != nil {
		t.Fail()
	}
	s := string(b)

	assert.Equal(t, exampleLcov(), s)

}

func exampleGeneric() string {
	return `<coverage version="1">
  <file path="xources/hello/NoConditions.xoo">
    <lineToCover lineNumber="6" covered="true"></lineToCover>
    <lineToCover lineNumber="7" covered="false"></lineToCover>
  </file>
  <file path="xources/hello/WithConditions.xoo">
    <lineToCover lineNumber="3" covered="true"></lineToCover>
  </file>
</coverage>`
}

func exampleLcov() string {
	return `TN:
SF:xources/hello/NoConditions.xoo
DA:6,1
DA:7,0
LH:0
LF:2
end_of_record
TN:
SF:xources/hello/WithConditions.xoo
DA:3,1
LH:0
LF:1
end_of_record

`
}
