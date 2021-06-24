package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"

	"github.com/Nv7-Github/Bpp/parser"
	arg "github.com/alexflint/go-arg"
)

var p *arg.Parser

func handle(err error) {
	if err != nil {
		p.Fail(err.Error())
	}
}

// Build defines the "build" sub-command
type Build struct {
	Output string `help:"output file for executable" arg:"-o"`
	CC     string `help:"LLVM compiler (optional)"`
	Asm    string `help:"LLC compiler (optional)"`
	File   string `arg:"positional,-i,--input" help:"input B++ program"`
}

// Run defines the "run" sub-command
type Run struct {
	Args string `help:"arguments for program, comma-seperated"`
	File string `arg:"positional,-i,--input" help:"input B++ program"`
}

// Convert defines the "convert" sub-command
type Convert struct {
	Output string `help:"output B++ program" arg:"-o"`
	File   string `arg:"positional,-i,--input" help:"input Go program"`
}

// Args defines the program's arguments
type Args struct {
	Build   *Build   `arg:"subcommand:build" help:"compile a B++ program"`
	Run     *Run     `arg:"subcommand:run" help:"run a B++ program"`
	Convert *Convert `arg:"subcommand:convert" help:"convert a go program to a B++ program"`
	Time    bool     `help:"print timing for each stage" arg:"-t"`

	CPUProf string `help:"CPU pprof statistics output file"`
	Memprof string `help:"Heap pprof statistics output file"`
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var args Args
	p = arg.MustParse(&args)

	if args.CPUProf != "" {
		pprofFile, err := os.OpenFile(args.CPUProf, os.O_WRONLY|os.O_CREATE, os.ModePerm)
		handle(err)
		err = pprof.StartCPUProfile(pprofFile)
		handle(err)
		defer pprof.StopCPUProfile()
	}

	switch {
	case args.Build != nil:
		prog := ParseProg(args.Time, args.Build.File)
		CompileCmd(args, prog)
	case args.Run != nil:
		prog := ParseProg(args.Time, args.Run.File)
		RunCmd(args, prog)
	case args.Convert != nil:
		ConvertCmd(args)
	default:
		p.WriteUsage(os.Stdout)
	}

	if args.Memprof != "" {
		pprofFile, err := os.OpenFile(args.Memprof, os.O_WRONLY|os.O_CREATE, os.ModePerm)
		handle(err)
		err = pprof.WriteHeapProfile(pprofFile)
		handle(err)
	}
}

// ParseProg parses a Go program
func ParseProg(isTiming bool, filename string) *parser.Program {
	src, err := os.ReadFile(filename)
	handle(err)

	var start time.Time
	if isTiming {
		start = time.Now()
	}
	prog, err := parser.Parse(string(src))
	handle(err)
	if isTiming {
		fmt.Println("Parsed program in", time.Since(start))
	}

	return prog
}
