package flags

type FilesRetrieverFlags struct {
	Parallelism int  `default:"5" short:"p" help:"Download parallelism"`
	Force       bool `short:"f" help:"Bypass caches and reedownload files"`
}
