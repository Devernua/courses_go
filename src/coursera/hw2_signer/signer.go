package main

import (
	"sync"
	"sort"
	"strconv"
)

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

// TODO: need wrapper on MD5 which will protect function with mutex
// TODO: need wrapper on crc32 which will get result in put order!

func SingleHash(in, out chan interface{}) {
	// TODO: multithreading
	for data := range in {
		first := DataSignerCrc32(strconv.Itoa(data.(int)))
		second := DataSignerCrc32(DataSignerMd5(strconv.Itoa(data.(int))))
		out <- first + "~" + second
	}
}

func MultiHash(in, out chan interface{}) {
	// TODO: multithreading
	for data := range in {
		var result string
		for i := 0; i < 6; i++ {
			result += DataSignerCrc32(strconv.Itoa(i) + data.(string))
		}
		out <- result
	}
}

func CombineResults(in, out chan interface{}) {
	var collect []string
	for data := range in {
		collect = append(collect, data.(string))
	}
	sort.Strings(collect)

	var result string
	for idx, col := range collect {
		result += col
		if idx + 1 != len(collect) {
			result += "_"
		}
	}
	out <- result
}