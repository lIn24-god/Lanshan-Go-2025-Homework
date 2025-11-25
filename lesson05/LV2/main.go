//Lv2
//分几个简单的阶段完成这个作业：
//
//阶段1：使用指定数量的协程完成对某个变量的指定数量的自增操作，同时保证并发安全，在自增一定次数之后，需要得到期望的结果。
//阶段2：将一个自增操作封装成 task ，将自增任务通过一些策略分发给 goroutine 进行执行（课件里面有类似的例子）
//阶段3：把你现在实现的 task -> goroutine 的模式封装成一个工具包，使得任何外部的库都可以通过你的包提供的函数/方法进行 task 分发。（PS：考虑一下提供什么样的 API 让用户使用起来更方便）
//进阶要求（选做）：
//
//性能优化：进行一些性能优化，我们的协程池还可以如何改造？（一些提示：sync.Pool，多管道轮询分发 task，属于是课外扩展的内容，不懂可以问 AI）
//思考：做完性能优化之后，尝试思考或者百度这几个问题：
//为什么需要协程池？
//sync.Pool 能拿来干什么，为什么需要池？
//原生的 map （前面课有讲）不是并发安全的，了解 1.24 之前及之后的 sync.Map 的实现，他是怎么实现并发安全的？

package main

import (
	"fmt"
	"sync"
)

//自增操作封装

type Task struct {
	AddSelf func(n int) int
}

func main() {
	//实现并发安全自增操作
	lock := sync.Mutex{}
	wg := sync.WaitGroup{}
	ch := make(chan Task, 10)
	num := 0
	sum := 0
	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for t := range ch {
				lock.Lock()
				sum = t.AddSelf(num)
				lock.Unlock()
			}
		}()
	}
	////自增操作封装,channel
	for range 20 {
		x := Task{
			AddSelf: func(n int) int {
				for range 100000 {
					n++
				}
				return n
			},
		}
		ch <- x
	}
	close(ch)
	wg.Wait()
	fmt.Println(sum)
}
