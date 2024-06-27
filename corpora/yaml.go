package corpora

import (
	"io"
	"os"

	"gopkg.in/yaml.v3"
)

func unMarshal(cacheFile string) (index *CorporaIndex, err error) {
	var f *os.File
	var d []byte
	if f, err = os.Open(cacheFile); err != nil {
		return
	}
	defer f.Close()
	if d, err = io.ReadAll(f); err != nil {
		return
	}
	err = yaml.Unmarshal(d, &index)
	return
}

func (c *CorporaIndex) marshal(cacheFile string) (err error) {
	var f *os.File
	var d []byte

	if f, err = os.Create(cacheFile); err != nil {
		return
	}
	defer f.Close()

	if d, err = yaml.Marshal(c); err != nil {
		return
	}
	_, err = f.Write(d)
	return
}
