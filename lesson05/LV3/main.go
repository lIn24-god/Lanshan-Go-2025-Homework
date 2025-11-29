package main

import (
	"fmt"
	"os"
	"sync"
)

type WorkPoll struct {
	tasks chan func()
	wg    sync.WaitGroup
}

func New(workNumber int) *WorkPoll {
	pool := &WorkPoll{
		tasks: make(chan func()),
		wg:    sync.WaitGroup{},
	}
	for i := 0; i < workNumber; i++ {
		pool.wg.Add(1)
		pool.worker(i)
	}
	return pool
}

func (wp *WorkPoll) worker(id int) {
	defer wp.wg.Done()
	for task := range wp.tasks {
		task()
	}
}

func (wp *WorkPoll) Submit(task func()) {
	wp.tasks <- task
}

func (wp *WorkPoll) Wait() {
	close(wp.tasks)
	wp.wg.Wait()
}
