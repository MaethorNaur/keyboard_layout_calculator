package cli

import "github.com/8VIM/keyboard_layout_calculator/config"

type Config struct {
	Verbose     int
	Force       bool
	Parallelism int
	Output      string
	Config      config.Config
}
