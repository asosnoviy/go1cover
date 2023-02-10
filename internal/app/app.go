package app

import (
	"fmt"

	"github.com/asosnoviy/go1cover/pkg/fdbc"
	"github.com/asosnoviy/go1cover/pkg/metareader"
)

func Run() {

	reader := metareader.New("./testresourse/cf/designer")
	reader.Parse()

	fdbc := fdbc.New()
	fdbc.Init()
	fdbc.Attach()
	fdbc.Deattach()

	for md, v := range fdbc.Storage {
		module := metareader.Module{ModuleTUuid: md.ObjectID, ModuleType: md.PropertyID}
		fmt.Println(reader.CoverData().Data[module], v)
	}

}
