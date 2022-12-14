package golib

import (
	"sync"

	"github.com/hashicorp/go-multierror"
)

type Manager struct {
	worker  int
	wg      *sync.WaitGroup
	stopped bool

	tasks chan Task
}

type Worker struct {
	id      int
	manager *Manager
}

type Task interface {
	Run() error
	WhenError(error)
	StopWhenError() bool
}

func NewManager(worker int, bufSize ...int) *Manager {
	bs := 10000
	if len(bufSize) > 0 {
		bs = bufSize[0]
	}
	return &Manager{
		worker:  worker,
		wg:      &sync.WaitGroup{},
		stopped: false,

		tasks: make(chan Task, bs),
	}
}

func (w *Worker) listen(tasks chan Task, errCh chan error) {
	// 	log.Debug.Println(w.id, "listen")
	for task := range tasks {
		if !w.manager.stopped {
			err := task.Run()
			if err != nil {
				task.WhenError(err)
				errCh <- err
				if task.StopWhenError() {
					w.manager.stopped = true
				}
			}
		}
		w.manager.wg.Done()
	}
	// 	log.Debug.Println(w.id, "done")
}

func (w *Worker) startListen(tasks chan Task) {
	for task := range tasks {
		w.manager.wg.Add(1)
		err := task.Run()
		if err != nil {
			task.WhenError(err)
		}
		w.manager.wg.Done()
	}
}

func (a *Manager) Start(tasks chan Task) {
	go func() {
		for t := range tasks {
			a.tasks <- t
		}
	}()
	for i := 0; i < a.worker; i++ {
		w := &Worker{
			id:      i,
			manager: a,
		}
		go w.startListen(a.tasks)
	}
}

func (a *Manager) LenTasks() int {
	return len(a.tasks)
}

func (a *Manager) Stop() {
	close(a.tasks)
	a.wg.Wait()
}

func (a *Manager) Do(tasks []Task) error {
	ts := make(chan Task)
	err := make(chan error)
	done := make(chan struct{})
	var res error
	go func() {
	THIS:
		for {
			select {
			case <-done:
				break THIS
			case e := <-err:
				res = multierror.Append(res, e)
			}
		}
	}()
	for i := 0; i < a.worker; i++ {
		w := &Worker{
			id:      i,
			manager: a,
		}
		go w.listen(ts, err)
	}
	a.wg.Add(len(tasks))
	for _, task := range tasks {
		ts <- task
	}
	close(ts)
	a.wg.Wait()
	done <- struct{}{}
	close(err)
	return res
}
