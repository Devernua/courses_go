package main

import "sync"

func startTask(task job, in, out chan interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(out)

	task(in, out)
}

// сюда писать код
func ExecutePipeline(tasks ...job) {
	wg := &sync.WaitGroup{}

	var curIn, curOut chan interface{}
	for _, task := range tasks {
		curIn = curOut
		curOut = make(chan interface{})

		wg.Add(1)
		go startTask(task, curIn, curOut, wg)
	}

	wg.Wait()
}

func SingleHash(in, out chan interface{}) {

}

func MultiHash(in, out chan interface{}) {

}

func CombineResults(in, out chan interface{}) {

}