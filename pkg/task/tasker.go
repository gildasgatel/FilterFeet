package task

import (
	"filterfleet/pkg/filter"
	"filterfleet/pkg/utils"

	"io/fs"
	"log"
	"sync"
)

type Tasker interface {
	ProcessWg()
	ProcessChan()
	SetDst(dst string)
	SetSrc(src string)
}

type Task struct {
	filter     filter.Filter
	filesEntry []fs.DirEntry
	src, dst   string
}

func New(f filter.Filter, filesEntry []fs.DirEntry) Tasker {
	return &Task{filter: f, filesEntry: filesEntry}
}

func (t Task) ProcessWg() {
	var wg sync.WaitGroup
	for _, file := range t.filesEntry {
		if utils.IsImage(file.Name()) {
			wg.Add(1)
			name := file.Name()
			go func() {
				err := t.filter.Process(t.src, t.dst, name)
				if err != nil {
					log.Fatal(err)
				}
				wg.Done()
			}()

		}

	}
	wg.Wait()
}

func (t Task) ProcessChan() {
	done := make(chan bool)
	for _, file := range t.filesEntry {
		if utils.IsImage(file.Name()) {
			name := file.Name()
			go func() {
				err := t.filter.Process(t.src, t.dst, name)
				if err != nil {
					log.Fatal(err)
				}
				done <- true
			}()
		}
	}

	for range t.filesEntry {
		<-done
	}
}

func (t *Task) SetSrc(src string) {
	t.src = src
}

func (t *Task) SetDst(dst string) {
	t.dst = dst
}
