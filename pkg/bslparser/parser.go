package bslparser

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"sync"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
	"github.com/asosnoviy/go1cover/pkg/bslparser/parser"
)

type Bslparser struct {
	wg sync.WaitGroup
	sync.Mutex
	LinesToCover map[string][]int
}

func New() *Bslparser {
	return &Bslparser{LinesToCover: make(map[string][]int)}
}

func (p *Bslparser) Files(folder string) ([]string, error) {
	var files []string
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {

		if filepath.Ext(path) == ".bsl" {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

func (p *Bslparser) Parse(files []string) error {

	for _, file := range files {
		p.wg.Add(1)
		go p.parsefile(file)
	}

	p.wg.Wait()
	return nil
}

func (p *Bslparser) Parsefile(path string) []int {

	p.wg.Add(1)
	p.parsefile(path)

	v, ok := p.LinesToCover[path]
	if ok {
		return v
	}

	return []int{}
}

func (p *Bslparser) LinesCount() int {
	linescount := 0
	for _, v := range p.LinesToCover {
		linescount += len(v)
	}
	return linescount
}

func (p *Bslparser) parsefile(path string) {
	_, err := os.Stat(path)
	if err != nil {
		fmt.Println("filr not found:", path)
		return
	}
	is, _ := antlr.NewFileStream(path)
	lexer := parser.NewBSLLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := parser.NewBSLParser(stream)
	parser.RemoveErrorListeners()
	parser.Interpreter.SetPredictionMode(antlr.PredictionModeSLL)

	ast := parser.File()
	lines := getLines(ast)
	keys := make([]int, 0)
	for k := range lines {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	p.Mutex.Lock()
	p.LinesToCover[path] = keys
	p.Mutex.Unlock()
	p.wg.Done()
}
