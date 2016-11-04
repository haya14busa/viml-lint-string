package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"

	vimlparser "github.com/haya14busa/go-vimlparser"
	"github.com/haya14busa/go-vimlparser/ast"
	"github.com/haya14busa/go-vimlparser/token"
)

func main() {
	flag.Parse()
	if err := run(os.Stdout, flag.Args()); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(w io.Writer, files []string) error {
	for _, file := range files {
		if err := lintFile(w, file); err != nil {
			return err
		}
	}
	return nil
}

// LISENSE:
// The MIT License (MIT)
// Copyright (c) 2014 Kuniwak
// ref: https://github.com/Kuniwak/vint/blob/51ce7d3b96a79ce62296e5b86407816ce005bdf6/vint/linting/policy/prohibit_unnecessary_double_quote.py#L10
var (
	specials = `\\(` + strings.Join([]string{
		`(?P<octal>[0-7]{1,3})`,
		`(?P<hexadecimal>[xX][0-9a-fA-F]{1,2})`,
		`(?P<numeric_character_reference>[uU][0-9a-fA-F]{4})`,
		`(?P<backspace>b)`,
		`(?P<escape>e)`,
		`(?P<form_feed>f)`,
		`(?P<new_line>n)`,
		`(?P<carriage_return>r)`,
		`(?P<tab>t)`,
		`(?P<backslash>\\)`,
		`(?P<double_quote>\")`,
		`(?P<special_key><[^>]+>)`,
	}, `|`) + ")"
	special = regexp.MustCompile(`'|` + specials)
)

func lintFile(w io.Writer, fname string) error {
	file, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer file.Close()

	var fbuf bytes.Buffer
	r := io.TeeReader(file, &fbuf)

	opt := &vimlparser.ParseOption{}
	f, err := vimlparser.ParseFile(r, file.Name(), opt)
	if err != nil {
		return err
	}

	// Inspect the AST and print all identifiers and literals.
	ast.Inspect(f, func(node ast.Node) bool {
		switch node := node.(type) {
		case *ast.BasicLit:
			if node.Kind == token.STRING {
				if node.Value[0] == '"' && !special.MatchString(node.Value) {
					fmt.Fprintf(w, "%s:%d:%d: %s", file.Name(), node.Pos().Line, node.Pos().Column, "Prefer single quoted strings\n")
				}
			}
		}
		return true
	})

	return nil
}
