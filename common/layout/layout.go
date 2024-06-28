package layout

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/8VIM/keyboard_layout_calculator/common/bigrams"
	"gopkg.in/yaml.v3"
)

const CharactersPerLayer = 32

type Layout struct {
	Version string `yaml:"version"`
	Info    Info   `yaml:"info"`
	Layers  Layers `yaml:"layers"`
}

type Info struct {
	Name string `yaml:"name"`
}

type Layers struct {
	Default    Layer            `yaml:"default"`
	ExtraLayer map[string]Layer `yaml:"extra_layers,omitempty"`
}

type Layer struct {
	Sectors map[string]Sector `yaml:"sectors"`
}

type Sector struct {
	Parts map[string][]Part `yaml:"parts,omitempty"`
}

type Part struct {
	LowerCase string `yaml:"lower_case"`
	UpperCase string `yaml:"upper_case,omitempty"`
}

type otherCharacter struct {
	layer, pos   int
	sector, part string
}

func New(name string) *Layout {
	l := &Layout{Version: "2", Info: Info{Name: name}}
	l.Layers.ExtraLayer = make(map[string]Layer)
	return l
}

func (l *Layout) AddFromNGram(n *bigrams.NGram) (err error) {
	re := regexp.MustCompile(`\p{L}`)
	layers := make([]Layer, 0)
	sector := "right"
	layer := 0
	pos := 0
	part := "bottom"
	var other *otherCharacter
	var entries []string
	entries, err = n.ExtractLetters()
	if err != nil {
		return
	}
	for _, entry := range entries {
		isChar := re.MatchString(entry)
		if other != nil && !isChar {
			layers[other.layer].Sectors[other.sector].Parts[other.part][other.pos].UpperCase = entry
			other = nil
			continue
		}
		if layer >= 6 {
			break
		}
		if len(layers) < layer+1 {
			layers = append(layers, Layer{Sectors: make(map[string]Sector)})
		}

		if _, ok := layers[layer].Sectors[sector]; !ok {
			layers[layer].Sectors[sector] = Sector{Parts: make(map[string][]Part)}
		}
		if _, ok := layers[layer].Sectors[sector].Parts[part]; !ok {
			layers[layer].Sectors[sector].Parts[part] = make([]Part, 0)
		}
		if isChar || other == nil {
			layers[layer].Sectors[sector].Parts[part] = append(layers[layer].Sectors[sector].Parts[part], Part{LowerCase: entry})
			if !isChar {
				other = &otherCharacter{layer: layer, pos: pos, sector: sector, part: part}
			}
		}
		sector, part = nextSectorPart(sector, part)
		if sector == "right" && part == "bottom" {
			pos++
			if pos >= 4 {
				layer++
				pos = 0
			}
		}
	}
	for i, layer := range layers {
		if i == 0 {
			l.Layers.Default = layer
		} else {
			l.Layers.ExtraLayer[idToString(i)] = layer
		}
	}
	return
}

func (l *Layout) Save(output string) (err error) {
	var f *os.File
	var d []byte
	if _, errStat := os.Stat(output); os.IsNotExist(errStat) {
		err = os.MkdirAll(output, os.ModeDir)
		if err != nil {
			return
		}
	}

	name := filepath.Join(output, fmt.Sprintf("%s.yaml", l.Info.Name))
	if f, err = os.Create(name); err != nil {
		return
	}
	defer f.Close()

	if d, err = yaml.Marshal(l); err != nil {
		return
	}
	_, err = f.Write(d)
	return
}

func nextSectorPart(sector, part string) (nextSector, nextPart string) {
	nextSector = sector
	nextPart = part
	switch sector {
	case "right":
		switch part {
		case "top":
			nextPart = "bottom"
		default:
			nextSector = "bottom"
			nextPart = "right"
		}
	case "bottom":
		switch part {
		case "left":
			nextSector = "left"
			nextPart = "bottom"
		default:
			nextPart = "left"
		}
	case "left":
		switch part {
		case "bottom":
			nextPart = "top"
		default:
			nextSector = "top"
			nextPart = "left"
		}
	case "top":
		switch part {
		case "left":
			nextPart = "right"
		default:
			nextSector = "right"
			nextPart = "top"
		}
	}
	return nextSector, nextPart
}
