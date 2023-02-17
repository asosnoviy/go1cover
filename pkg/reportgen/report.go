package reportgen

import coveragdata "github.com/asosnoviy/go1cover/internal/coverageData"

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

func New(reporter ...ReporterInt) Reportgen {

	return Reportgen{reporters: reporter}
}

func (g Reportgen) Report(coverage *coveragdata.Coverage) {

	for _, repot := range g.reporters {

		repot.Generate(coverage)
	}

}
