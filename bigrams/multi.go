package bigrams

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/8VIM/keyboard_layout_calculator/config"
	"github.com/pterm/pterm"
)

type task struct {
	name     string
	language config.Language
	spinner  *pterm.SpinnerPrinter
}

type Runner struct {
	parallelism int
	force       bool
	tasks       map[string]task
	multi       *pterm.MultiPrinter
	wg          sync.WaitGroup
	mu          sync.Mutex
	cacheDir    string
}

func New(multi *pterm.MultiPrinter, parallelism int, force bool) (r *Runner, err error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	dir := filepath.Join(homeDir, ".config", "keyboard_layout_calculator", "cache")

	if err = os.MkdirAll(dir, os.ModeDir); err != nil && !os.IsExist(err) {
		return nil, err
	}
	return &Runner{parallelism: parallelism, force: force, tasks: make(map[string]task), multi: multi, cacheDir: dir}, nil
}

func (r *Runner) Add(name string, language config.Language) {
	name = strings.ToLower(name)
	if _, ok := r.tasks[name]; !ok {
		r.tasks[name] = task{name: name, language: language}
	}
}

func (r *Runner) Load() (bigrams map[string]*NGram) {
	total := len(r.tasks)
	r.wg.Add(total)

	tasks := make(chan task, total)
	bigrams = make(map[string]*NGram)

	for w := 1; w <= r.parallelism; w++ {
		go r.worker(tasks, bigrams)
	}

	for _, task := range r.tasks {
		tasks <- task
	}
	close(tasks)
	r.wg.Wait()
	return
}

func (r *Runner) worker(tasks <-chan task, bigrams map[string]*NGram) {
	for task := range tasks {
		spinner, _ := pterm.DefaultSpinner.WithWriter(r.multi.NewWriter()).Start(fmt.Sprintf("[%s corpuses] Loading", task.name))
		task.spinner = spinner
		bigram, err := r.load(task)
		if err != nil {
			spinner.Fail(fmt.Sprintf("[%s corpuses] Failed: %s", task.name, err))
			r.wg.Done()
			return
		}
		r.mu.Lock()
		bigrams[task.name] = bigram
		r.mu.Unlock()
		r.wg.Done()
		spinner.Success(fmt.Sprintf("[%s corpuses] Done", task.name))
	}
}
