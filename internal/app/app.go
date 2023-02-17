package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/asosnoviy/go1cover/internal/config"
	coverserver "github.com/asosnoviy/go1cover/internal/coverServer"
	coveragedata "github.com/asosnoviy/go1cover/internal/coverageData"
	"github.com/asosnoviy/go1cover/pkg/bslparser"
	"github.com/asosnoviy/go1cover/pkg/fdbc"
	"github.com/asosnoviy/go1cover/pkg/metareader"
	"github.com/asosnoviy/go1cover/pkg/reportgen"
)

func Run() {

	config := &config.Config{
		SrcPath:    "testresourse/cf/designer",
		DebugerUrl: "http://localhost:1760",
		Reporters:  []string{"lcov", "generic"},
		RootPath:   "",
		Port:       ":4040",
	}

	runChan := make(chan bool)
	stopChan := make(chan bool)
	termchan := make(chan os.Signal, 5)
	signal.Notify(termchan, os.Interrupt, syscall.SIGTERM)

	server := coverserver.New(config.Port, runChan, stopChan)
	go server.Run()

	reader := metareader.New(config.SrcPath)
	parser := bslparser.New()
	fdebug := fdbc.New(config.DebugerUrl)

	reporters := reportgen.NewReporters(config.Reporters)
	reportgen := reportgen.New(reporters...)

	for {
		select {
		case <-runChan:
			{
				go Start(reader, parser, fdebug, config)
			}
		case <-stopChan:
			{
				go Stop(reader, parser, fdebug, reportgen, config)
				fmt.Println("Goob bye")
			}
		case <-termchan:
			{
				os.Exit(10)
			}

		}
	}

}

func Start(reader *metareader.Metareader, parser *bslparser.Bslparser, fdebug *fdbc.Fdbc, config *config.Config) {

	println("---Start---")

	reader.Parse()

	files := reader.Files()
	go parser.Parse(files)

	fdebug.Init()
	fdebug.Attach()

}

func Stop(reader *metareader.Metareader, parser *bslparser.Bslparser, fdebug *fdbc.Fdbc, reportgen *reportgen.Reportgen, config *config.Config) {

	println("---Stop---")

	fdebug.Deattach()
	fdebug.Stop()
	coverage := coveragedata.Convert(parser, reader, fdebug, config.RootPath)
	reportgen.Report(coverage)

}
