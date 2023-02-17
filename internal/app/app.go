package app

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	coveragdata "github.com/asosnoviy/go1cover/internal/coverageData"
	"github.com/asosnoviy/go1cover/pkg/bslparser"
	"github.com/asosnoviy/go1cover/pkg/fdbc"
	"github.com/asosnoviy/go1cover/pkg/metareader"
	"github.com/asosnoviy/go1cover/pkg/reportgen"
)

func Run() {

	var wg sync.WaitGroup
	wg.Add(1)

	reader := metareader.New("./testresourse/cf/designer")
	reader.Parse()
	files := reader.Files()

	parser := bslparser.New()
	go parser.Parse(files)

	fdebug := fdbc.New()
	fdebug.Init()
	fdebug.Attach()

	lcov := reportgen.NewLcov("lcov.info")
	generic := reportgen.NewGeneric("generic.xml")
	reportgen := reportgen.New(generic, lcov)

	termchan := make(chan os.Signal)
	signal.Notify(termchan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-termchan
		wg.Done()

	}()

	wg.Wait()

	fdebug.Deattach()
	coverage := coveragdata.Convert(parser, reader, fdebug)
	reportgen.Report(coverage)
	fmt.Println("Goob bye")
}
