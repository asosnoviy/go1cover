package app

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/asosnoviy/go1cover/internal/config"
	coveragdata "github.com/asosnoviy/go1cover/internal/coverageData"
	"github.com/asosnoviy/go1cover/pkg/bslparser"
	"github.com/asosnoviy/go1cover/pkg/fdbc"
	"github.com/asosnoviy/go1cover/pkg/metareader"
	"github.com/asosnoviy/go1cover/pkg/reportgen"
)

func Run() {

	var wg sync.WaitGroup
	wg.Add(1)

	config := &config.Config{
		SrcPath:    "testresourse/cf/designer",
		DebugerUrl: "http://localhost:1760",
		Reporters:  []string{"lcov", "generic"},
		RootPath:   "",
	}

	reader := metareader.New(config.SrcPath)
	reader.Parse()
	files := reader.Files()

	parser := bslparser.New()
	go parser.Parse(files)

	fdebug := fdbc.New(config.DebugerUrl)
	fdebug.Init()
	fdebug.Attach()

	reporters := reportgen.NewReporters(config.Reporters)
	reportgen := reportgen.New(reporters...)

	termchan := make(chan os.Signal)
	signal.Notify(termchan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-termchan
		wg.Done()

	}()

	wg.Wait()

	fdebug.Deattach()
	coverage := coveragdata.Convert(parser, reader, fdebug, config.RootPath)
	reportgen.Report(coverage)

	fmt.Println("Coverage:", fmt.Sprint(coveragdata.Covered(fdebug, parser)), "%")
	fmt.Println("Goob bye")
}
