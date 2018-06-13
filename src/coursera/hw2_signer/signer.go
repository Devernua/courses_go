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
//type MD5LockWrapper struct {
//	mtx *sync.Mutex
//}

// TODO: make one function which get separator and write result
func calcSingleHash(strData, md5Result string, prev, next chan struct {}, out chan interface{}) {
	defer func(){
		next <- struct{}{}
		close(next)
	}()

	lht := make(chan string)
	rht := make(chan string)

	go calcCRC32(strData, lht)
	go calcCRC32(md5Result, rht)

	<-prev // wait complete

	left := <-lht
	right := <- rht

	out <- left + "~" + right
}

func calcMultiHash(strs []string, prev, next chan struct {}, out chan interface{}) {
	defer func(){
		next <- struct{}{}
		close(next)
	}()

	channels := make([]chan string, len(strs))

	for idx, str := range strs {
		channels[idx] = make(chan string)
		go calcCRC32(str, channels[idx])
	}

	<- prev // wait complete

	result := ""
	for _, c := range channels {
		str := <-c
		result += str
	}

	out <- result
}

func calcCRC32(data string, out chan string) {
	defer close(out)
	out <- DataSignerCrc32(data)
	return
}

func SingleHash(in, out chan interface{}) {
	var prev chan struct{}
	next := make(chan struct{})

	needNext := true
	for data := range in {
		// todo: make like overhead lock in common
		strData := strconv.Itoa(data.(int))
		md5Result := DataSignerMd5(strData)

		prev = next
		next = make(chan struct{})

		go calcSingleHash(strData, md5Result, prev, next, out)

		if needNext {
			needNext = false
			prev <- struct{}{}
			close(prev)
		}
	}

	<- next // complete // TODO: may be not needed
	return
}

func MultiHash(in, out chan interface{}) {
	var prev chan struct{}
	next := make(chan struct{})

	needNext := true
	for data := range in {
		var result []string
		for i := 0; i < 6; i++ {
			result = append(result, strconv.Itoa(i) + data.(string))
		}

		prev = next
		next = make(chan struct{})

		go calcMultiHash(result, prev, next, out)

		if needNext {
			needNext = false
			prev <- struct{}{}
			close(prev)
		}
	}

	<- next // complete // TODO: may be not needed
	return
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
	return
}