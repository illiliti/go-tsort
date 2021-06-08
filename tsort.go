package tsort

import (
	"errors"
	"fmt"
)

const (
	white = iota
	black
	grey
)

var ErrCircularLoop = errors.New("endless loop")

type Graph []*Vertex

type Vertex struct {
	Name string

	cl int
	tg *Graph
	ee map[string]*Edge
}

type Edge struct {
	Name string

	v *Vertex
}

func (g *Graph) AddVertex(n string) *Vertex {
	// TODO optimize
	for _, v := range *g {
		if v.Name == n {
			return v
		}
	}

	v := &Vertex{
		Name: n,
		tg:   g,
		ee:   make(map[string]*Edge),
	}

	*g = append(*g, v)
	return v
}

func (v *Vertex) AddEdge(n string) *Edge {
	if e, ok := v.ee[n]; ok {
		return e
	}

	e := &Edge{
		Name: n,
		v:    v.tg.AddVertex(n),
	}

	v.ee[n] = e
	return e
}

func (g *Graph) Sort() ([]string, error) {
	var oo []string

	for _, v := range *g {
		if err := traversal(v, &oo); err != nil {
			return nil, err
		}
	}

	return oo, nil
}

func traversal(v *Vertex, oo *[]string) error {
	if v.cl == black {
		return nil
	}

	v.cl = grey

	// pre-order
	*oo = append(*oo, v.Name)

	for _, e := range v.ee {
		switch e.v.cl {
		case white:
			if err := traversal(e.v, oo); err != nil {
				return err
			}
		case black:
			continue
		case grey:
			return fmt.Errorf("%w: [%s %s]", ErrCircularLoop, e.v.Name, v.Name)
		}
	}

	// post-order
	// *oo = append(*oo, v.Name)

	v.cl = black
	return nil
}
