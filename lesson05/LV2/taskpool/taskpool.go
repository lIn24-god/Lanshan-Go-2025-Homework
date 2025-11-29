package taskpool

import (
	"fmt"
	"sync"
)

type WorkerPool struct {
	tasks chan func()
	wg    sync.WaitGroup
}

func NewWorkPoll(workerCount int) *WorkerPool {
	pool := &WorkerPool{
		tasks: make(chan func(), 100),
	}
	for i := 0; i < workerCount; i++ {
		pool.wg.Add(1)
		go pool.worker(i)
	}
	return pool
}

func (p *WorkerPool) worker(id int) {
	defer p.wg.Done()
	for task := range p.tasks {
		fmt.Println(id, "开始执行任务")
		task()
		fmt.Println(id, "任务执行完成")
	}
}

func (p *WorkerPool) Submit(task func()) {
	p.tasks <- task
}

func (p *WorkerPool) Wait() {
	close(p.tasks)
	p.wg.Wait()
}
