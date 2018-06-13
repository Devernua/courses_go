package main

import "sync"

func startTask(task job, in, out chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	task(in, out)
	close(out)
}

// сюда писать код
func ExecutePipeline(tasks ...job) {
	wg := &sync.WaitGroup{}

	var curIn chan interface{}
	curOut := make(chan interface{})
	for idx, task := range tasks {
		wg.Add(1)
		go startTask(task, curIn, curOut, wg)

		if idx + 1 != len(tasks) {
			curIn = curOut
			curOut = make(chan interface{})
		}
	}

	wg.Wait()
}

func SingleHash(in, out chan interface{}) {

}

func MultiHash(in, out chan interface{}) {

}

func CombineResults(in, out chan interface{}) {

}