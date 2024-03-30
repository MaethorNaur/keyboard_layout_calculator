package bigrams

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/awalterschulze/gographviz"
	"github.com/hmdsefi/gograph"
	"github.com/hmdsefi/gograph/traverse"
	"golang.org/x/exp/utf8string"
)

func (n *NGram) ExtractLetters() (res []string, err error) {
	graph := gograph.New[string](gograph.Weighted())
	scores := make(map[string]float64)
	for _, e := range n.Monograms {
		graph.AddVertexByLabel(e.Value)
		scores[e.Value] = 1
	}

	max := float64(n.Bigrams[0].Frequency)
	for _, e := range n.Bigrams {
		u := utf8string.NewString(e.Value)
		s, d := u.Slice(0, 1), u.Slice(1, 2)

		vs, vd := graph.GetVertexByID(s), graph.GetVertexByID(d)
		weight := 1. - float64(e.Frequency)/max
		_, _ = graph.AddEdge(vs, vd, gograph.WithEdgeWeight(weight))
	}

	res = make([]string, 0)
	set := make(map[string]struct{})
	for _, e := range n.Monograms {
		if _, ok := set[e.Value]; ok {
			continue
		}
		it, _ := traverse.NewClosestFirstIterator[string](graph, e.Value)
		_ = it.Iterate(func(v *gograph.Vertex[string]) error {
			if _, ok := set[v.Label()]; !ok {
				res = append(res, v.Label())
				set[v.Label()] = struct{}{}
			}
			return nil
		})
	}
	// for _, e := range n.Monograms {
	// 	distances := path.Dijkstra[string](graph, e.Value)
	// 	max := 0.

	// 	for _, v := range distances {
	// 		if v == math.MaxFloat64 {
	// 			continue
	// 		}
	// 		max = math.Max(max, v)
	// 	}

	// 	for k, v := range distances {
	// 		v = v / max
	// 		d := math.Abs(scores[k] - v)
	// 	}
	// 	//TODO e := n.Monograms[len(n.Monograms)-1-i]
	// 	if _, ok := set[e.Value]; ok {
	// 		continue
	// 	}

	// 	if v := graph.GetVertexByID(e.Value); v == nil {
	// 		continue
	// 	}

	// 	it, _ := traverse.NewClosestFirstIterator[string](graph, e.Value)
	// 	_ = it.Iterate(func(v *gograph.Vertex[string]) error {
	// 		if _, ok := set[v.Label()]; ok {
	// 			return nil
	// 		}
	// 		res = append(res, v.Label())
	// 		set[v.Label()] = struct{}{}
	// 		return nil
	// 	})
	// }
	return
}

func (n *NGram) Graphivz(output, name string) {
	filename := filepath.Join(output, fmt.Sprintf("%s.dot", name))
	var f *os.File
	var err error
	if f, err = os.Create(filename); err != nil {
		return
	}
	defer f.Close()
	graph := gographviz.NewGraph()
	_ = graph.SetName("G")
	_ = graph.AddAttr("G", "root", escape(n.Monograms[0].Value))
	_ = graph.AddAttr("G", "overlap", "scalexy")
	for _, e := range n.Monograms {
		_ = graph.AddNode("G", escape(e.Value), map[string]string{"shape": "box"})
	}
	for _, e := range n.Bigrams {
		u := utf8string.NewString(e.Value)
		s, d := u.Slice(0, 1), u.Slice(1, 2)
		_ = graph.AddEdge(escape(s), escape(d), false, map[string]string{"weight": fmt.Sprintf("%d", e.Frequency)})
	}
	_, _ = f.WriteString(graph.String())
}

func escape(s string) string {
	esc := ""
	if s == `"` {
		esc = `\`
	}
	return fmt.Sprintf(`"%s%s"`, esc, s)
}
