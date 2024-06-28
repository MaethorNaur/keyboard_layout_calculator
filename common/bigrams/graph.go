package bigrams

import (
	"fmt"
	"os"

	"github.com/awalterschulze/gographviz"
	// "github.com/hmdsefi/gograph"
	// "github.com/hmdsefi/gograph/traverse"
	// "golang.org/x/exp/utf8string"
)

func (n *NGram) Dot(filename string) error {
	graph := gographviz.NewGraph()
	for _, e := range n.Monograms {
		if err := graph.AddNode("", e.Value, nil); err != nil {
			return err
		}
	}

	var err error
	var f *os.File

	if f, err = os.Create(filename); err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(graph.String()); err != nil {
		return err
	}
	return nil
}
func (n *NGram) ExtractLetters() (res []string, err error) {
	res = make([]string, 0)
	for _, e := range n.MonogramsInBigrams {
		res = append(res, e.Value)
	}
	return
	// graph := gograph.New[string](gograph.Weighted())

	// scores := make(map[string]float64)

	// layer := 0
	// for i, e := range n.Monograms {
	// 	position := i % 8
	// 	if position == 0 {
	// 		layer++
	// 	}
	// 	if layer > 6 {
	// 		break
	// 	}
	// 	scores[e.Value] = float64(layer * (position + 1))
	// }
	// for _, e := range n.Monograms {
	// 	graph.AddVertexByLabel(e.Value)
	// }

	// max := float64(n.Bigrams[0].Frequency)
	// for _, e := range n.Bigrams {
	// 	u := utf8string.NewString(e.Value)
	// 	s, d := u.Slice(0, 1), u.Slice(1, 2)

	// 	vs, vd := graph.GetVertexByID(s), graph.GetVertexByID(d)
	// 	weight := 1. - float64(e.Frequency)/max
	// 	_, _ = graph.AddEdge(vs, vd, gograph.WithEdgeWeight(weight))
	// }

	// res = make([]string, 0)
	// set := make(map[string]struct{})
	// for _, e := range n.Monograms {
	// 	if _, ok := set[e.Value]; ok {
	// 		continue
	// 	}
	// 	var it traverse.Iterator[string]
	// 	it, err = traverse.NewClosestFirstIterator[string](graph, e.Value)
	// 	if err != nil {
	// 		return
	// 	}

	// 	_ = it.Iterate(func(v *gograph.Vertex[string]) error {
	// 		if _, ok := set[v.Label()]; !ok {
	// 			res = append(res, v.Label())
	// 			set[v.Label()] = struct{}{}
	// 		}
	// 		return nil
	// 	})
	// }
	// // for _, e := range n.Monograms {
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
}

func escape(s string) string {
	esc := ""
	if s == `"` {
		esc = `\`
	}
	return fmt.Sprintf(`"%s%s"`, esc, s)
}
