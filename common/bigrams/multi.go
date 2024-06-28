package bigrams

import (
	"sync"

	"github.com/8VIM/keyboard_layout_calculator/common/config"
)

type task struct {
	language  *config.ConfigLanguage
	callbacks *Callbacks
}

type Runner struct {
	parallelism int
	force       bool
	cacheDir    string
	tasks       map[string]task
	onTask      func() *Callbacks
	wg          sync.WaitGroup
	mu          sync.Mutex
}

const (
	Downloading byte = iota
	Extracting
)

type Callbacks struct {
	onStart   func(string)
	onFetch   func(byte, string, string, string)
	onCompute func(string)
	onError   func(string, error)
	onDone    func(string)
}

func NewCallbacks(onStart func(string),
	onFetch func(byte, string, string, string),
	onCompute func(string),
	onError func(string, error),
	onDone func(string)) *Callbacks {
	return &Callbacks{onStart: onStart,
		onFetch:   onFetch,
		onCompute: onCompute,
		onError:   onError,
		onDone:    onDone}
}

func New(parallelism int, force bool, cacheDir string, onTask func() *Callbacks) *Runner {
	return &Runner{parallelism: parallelism,
		force:    force,
		cacheDir: cacheDir,
		tasks:    make(map[string]task),
		onTask:   onTask,
	}
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
		task.callbacks = r.onTask()
		task.callbacks.onStart(task.language.Name())
		bigram, err := r.load(task)
		if err != nil {
			task.callbacks.onError(task.language.Name(), err)
			r.wg.Done()
			return
		}
		r.mu.Lock()
		bigrams[task.language.Name()] = bigram
		r.mu.Unlock()
		r.wg.Done()
		task.callbacks.onDone(task.language.Name())
	}
}
