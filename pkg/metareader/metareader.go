package metareader

import (
	"errors"
	"os"
	"path"
)

type Metareader struct {
	ReaderCommon
	Reader Reader
}

type Reader interface {
	Parse()
	Files() []string
	CoverData() *coverData
}

func New(confpath string) *Metareader {

	var reader Reader

	if _, err := os.Stat(path.Join(confpath, "Configuration.xml")); errors.Is(err, os.ErrNotExist) {
		reader = NewEdtReader(confpath)
	} else {
		reader = NewDesignerReader(confpath)
	}

	return &Metareader{Reader: reader}
}

func (m *Metareader) ParseMeta() {
	m.Reader.Parse()
}

func (m *Metareader) CoverData() *coverData {

	return m.Reader.CoverData()
}
