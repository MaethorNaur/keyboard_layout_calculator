package bigrams

import (
	"fmt"
	"sync"

	"github.com/8VIM/keyboard_layout_calculator/config"
	"github.com/pterm/pterm"
)

type task struct {
	language *config.ConfigLanguage
	spinner  *pterm.SpinnerPrinter
}

type Runner struct {
	parallelism int
	force       bool
	cacheDir    string
	tasks       map[string]task
	multi       *pterm.MultiPrinter
	wg          sync.WaitGroup
	mu          sync.Mutex
}

func New(multi *pterm.MultiPrinter, parallelism int, force bool, cacheDir string) *Runner {
	return &Runner{parallelism: parallelism,
		force:    force,
		cacheDir: cacheDir,
		tasks:    make(map[string]task),
		multi:    multi}
}

func (r *Runner) Add(l *config.ConfigLanguage) {
	if _, ok := r.tasks[l.Name()]; !ok {
		r.tasks[l.Name()] = task{language: l}
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
		spinner, _ := pterm.DefaultSpinner.WithWriter(r.multi.NewWriter()).Start(fmt.Sprintf("[%s corpuses] Loading", task.language.Name()))
		task.spinner = spinner
		bigram, err := r.load(task)
		if err != nil {
			spinner.Fail(fmt.Sprintf("[%s corpuses] Failed: %s", task.language.Name(), err))
			r.wg.Done()
			return
		}
		r.mu.Lock()
		bigrams[task.language.Name()] = bigram
		r.mu.Unlock()
		r.wg.Done()
		spinner.Success(fmt.Sprintf("[%s corpuses] Done", task.language.Name()))
	}
}
