package app

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/asosnoviy/go1cover/internal/config"
	coverserver "github.com/asosnoviy/go1cover/internal/coverServer"
	coveragedata "github.com/asosnoviy/go1cover/internal/coverageData"
	"github.com/asosnoviy/go1cover/pkg/bslparser"
	"github.com/asosnoviy/go1cover/pkg/fdbc"
	"github.com/asosnoviy/go1cover/pkg/metareader"
	"github.com/asosnoviy/go1cover/pkg/reportgen"
)

type context struct {
	config      *config.Config
	coverserver *coverserver.CoverageServer
	reader      *metareader.Metareader
	parser      *bslparser.Bslparser
	fdebug      *fdbc.Fdbc
	reportgen   *reportgen.Reportgen
	RunChan     chan bool
	stopChan    chan bool
	termChan    chan os.Signal
}

func Init(config *config.Config) *context {

	runChan := make(chan bool, 1)
	stopChan := make(chan bool)
	termchan := make(chan os.Signal, 5)
	signal.Notify(termchan, os.Interrupt, syscall.SIGTERM)

	server := coverserver.New(":"+config.Port, runChan, stopChan)
	go server.Run()

	reader := metareader.New(config.SrcPath)
	parser := bslparser.New()
	fdebug := fdbc.New(config.DebugerUrl)

	reporters := reportgen.NewReporters(config.Reporters)
	reportgen := reportgen.New(reporters...)

	context := &context{
		config:      config,
		coverserver: server,
		reader:      reader,
		parser:      parser,
		fdebug:      fdebug,
		reportgen:   reportgen,
		RunChan:     runChan,
		stopChan:    stopChan,
		termChan:    termchan,
	}
	return context
}

func Run(context *context) {

	for {
		select {
		case <-context.RunChan:
			{
				go start(context)
			}
		case <-context.stopChan:
			{
				Stop(context)
				fmt.Println("Goob bye")
				return
			}
		case <-context.termChan:
			{
				Stop(context)
				os.Exit(10)
			}

		}
	}

}

func start(context *context) {

	println("---Start---")

	context.reader.ParseMeta()

	files := context.reader.Reader.Files()
	go context.parser.Parse(files)

	context.fdebug.Init()
	context.fdebug.Attach()

}

func Stop(context *context) {

	println("---Stop---")

	context.fdebug.Stop()
	context.fdebug.Deattach()

	coverage := coveragedata.Convert(context.parser.LinesToCover, context.reader.CoverData().Data, context.fdebug.Storage, context.config.RootPath)
	covered := coveragedata.Covered(context.fdebug, context.parser)
	fmt.Println("Line coverage:", strconv.Itoa(covered)+"%")
	context.reportgen.Report(coverage)

}
