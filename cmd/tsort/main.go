package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/illiliti/go-tsort"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "usage: %s [<file>]", os.Args[0])
	}

	flag.Parse()

	if n := flag.NArg(); n > 0 && n > 1 {
		flag.Usage()
		os.Exit(2)
	}

	if err := run(flag.Arg(0)); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", os.Args[0], err)
		os.Exit(1)
	}
}

func run(p string) error {
	in := os.Stdin

	if p != "" {
		var err error

		in, err = os.Open(p)

		if err != nil {
			return err
		}

		defer in.Close()
	}

	var (
		tg tsort.Graph
		vx *tsort.Vertex
	)

	sc := bufio.NewScanner(in)

	for sc.Scan() {
		for _, n := range strings.Fields(sc.Text()) {
			switch {
			case vx == nil:
				vx = tg.AddVertex(n)
			case vx.Name != n:
				vx.AddEdge(n)
				vx = nil
			default:
				vx = nil
			}
		}
	}

	if err := sc.Err(); err != nil {
		return err
	}

	if vx != nil {
		return errors.New("input contains odd number of elements")
	}

	nn, err := tg.Sort()

	if err != nil {
		return err
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for _, n := range nn {
		w.WriteString(n + "\n")
	}

	return nil
}
