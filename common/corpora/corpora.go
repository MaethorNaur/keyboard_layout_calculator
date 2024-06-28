package corpora

const CacheFile = "corpora.index.yaml"

type Corpora struct {
	Force       bool
	Parallelism int
	CacheFile   string
	index       *CorporaIndex
}
