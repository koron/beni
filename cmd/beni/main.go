package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/koron/beni"
	"github.com/koron/beni/formatter"
	"github.com/koron/beni/lexer"
	"github.com/koron/beni/theme"
)

// CLOptions has command line options
type CLOptions struct {
	Help      bool
	Lexer     string
	Formatter string
	Theme     string
}

var options CLOptions

func usage() {
	fmt.Println(`beni - Code highlighter

Usage: beni [OPTIONS] [FILES...]`)
	flag.PrintDefaults()
	fmt.Println(`
  Formatters:`)
	for _, v := range formatter.All {
		fmt.Println("    " + v.Info().Name)
	}
	fmt.Println(`
  Languages:`)
	for _, v := range lexer.All {
		fmt.Println("    " + v.Info().Name)
	}
	fmt.Println(`
  Themes:`)
	for _, v := range theme.All {
		fmt.Println("    " + v.GetName())
	}
	os.Exit(0)
}

func convert(filename string, w io.Writer, lexer, theme, format string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	if lexer == "" {
		lexer = filename
	}
	return beni.Highlight(f, w, lexer, theme, format)
}

func run(filenames []string, o CLOptions) error {
	for _, name := range filenames {
		if err := convert(name, getStdoutWriter(), o.Lexer, o.Theme, o.Formatter); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	flag.BoolVar(&options.Help, "h", false, "show help message")
	flag.StringVar(&options.Lexer, "l", "", "force lexer")
	flag.StringVar(&options.Formatter, "f", "Terminal256", "choose formatter")
	flag.StringVar(&options.Theme, "t", "base16", "choose formatter")
	flag.Parse()

	if options.Help || flag.NArg() == 0 {
		usage()
		return
	}

	if err := run(flag.Args(), options); err != nil {
		log.Fatal(err)
	}
}
