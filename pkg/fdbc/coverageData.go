package fdbc

type LIneCoverage struct {
	LineNo  string
	Covered bool
}

type ModuleData struct {
	ObjectID   string
	PropertyID string
}

type CoverageData struct {
	ModuleData
	LIneCoverage
}
