package coveragdata

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/asosnoviy/go1cover/pkg/bslparser"
	"github.com/asosnoviy/go1cover/pkg/fdbc"
	"github.com/asosnoviy/go1cover/pkg/metareader"
)

type Coverage struct {
	Data map[string]map[int]bool
}

func Convert(parser *bslparser.Bslparser, reader *metareader.Metareader, fdbc *fdbc.Fdbc) *Coverage {

	data := map[string]map[int]bool{}
	cwd, _ := os.Getwd()

	for k, v := range parser.LinesToCover {

		linesToCover := map[int]bool{}
		for _, notCoveredLines := range v {
			linesToCover[notCoveredLines] = false
		}

		path, _ := filepath.Rel(cwd, k)
		sleshedpath := filepath.ToSlash(path)
		data[sleshedpath] = linesToCover
	}

	for md, v := range fdbc.Storage {
		module := metareader.Module{ModuleUuid: md.ObjectID, ModuleType: md.PropertyID}
		for _, lc := range v {
			lineNo, _ := strconv.Atoi(lc.LineNo)

			path, _ := filepath.Rel(cwd, reader.CoverData().Data[module])
			sleshedpath := filepath.ToSlash(path)
			data[sleshedpath][lineNo] = true
		}

	}

	return &Coverage{Data: data}
}
