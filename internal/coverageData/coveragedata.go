package coveragedata

import (
	"log"
	"path/filepath"
	"strconv"

	"github.com/asosnoviy/go1cover/pkg/bslparser"
	"github.com/asosnoviy/go1cover/pkg/fdbc"
	"github.com/asosnoviy/go1cover/pkg/metareader"
)

type Coverage struct {
	Data map[string]map[int]bool
}

func Convert(LinesToCover map[string][]int, metaData map[metareader.Module]string, Storage map[fdbc.ModuleData][]fdbc.LIneCoverage, basePath string) *Coverage {

	data := map[string]map[int]bool{}

	cwd, err := filepath.Abs(basePath)
	if err != nil {
		log.Panic(err)
	}

	for k, v := range LinesToCover {

		linesToCover := map[int]bool{}
		for _, notCoveredLines := range v {
			linesToCover[notCoveredLines] = false
		}

		path, _ := filepath.Rel(cwd, k)
		sleshedpath := filepath.ToSlash(path)
		data[sleshedpath] = linesToCover
	}

	for md, v := range Storage {
		module := metareader.Module{ModuleUuid: md.ObjectID, ModuleType: md.PropertyID}
		for _, lc := range v {
			lineNo, _ := strconv.Atoi(lc.LineNo)

			path, _ := filepath.Rel(cwd, metaData[module])
			sleshedpath := filepath.ToSlash(path)
			data[sleshedpath][lineNo] = true
		}

	}

	return &Coverage{Data: data}
}

func Covered(fdebug *fdbc.Fdbc, parser *bslparser.Bslparser) int {

	parsetLines := parser.LinesCount()
	if parsetLines == 0 {
		return 0
	}

	covered := fdebug.LinesCount() * 100 / parsetLines
	return covered
}
