package reportgen

import (
	"log"
	"strings"

	coveragdata "github.com/asosnoviy/go1cover/internal/coverageData"
)

type Reportgen struct {
	reporters []ReporterInt
	coverage  coveragdata.Coverage
}

type ReporterInt interface {
	Generate(*coveragdata.Coverage)
}

type Reporter struct {
	Path string
}

func New(reporter ...ReporterInt) *Reportgen {

	return &Reportgen{reporters: reporter}
}

func NewReporters(reportersName []string) []ReporterInt {

	reportest := []ReporterInt{}
	for _, reporterName := range reportersName {

		switch strings.TrimSpace(reporterName) {
		case "lcov":
			{
				reportest = append(reportest, NewLcov("lcov.info"))
			}
		case "generic":
			{
				reportest = append(reportest, NewGeneric("generic.xml"))
			}
		default:
			log.Fatalf("Unknown reporter %s", reporterName)

		}

	}

	return reportest
}

func (g *Reportgen) Report(coverage *coveragdata.Coverage) {

	for _, repot := range g.reporters {

		repot.Generate(coverage)
	}

}
