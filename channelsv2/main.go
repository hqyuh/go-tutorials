package main

import (
	"fmt"
	"sync"
)

func main()  {
	var wg = sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	go func(ch <- chan int) { // lấy từ chan đẩy dô
		fmt.Println("Chan 1")
		i := <- ch
		fmt.Println(i)
		wg.Done()
	}(ch)
	go func(ch chan <- int) { // lấy value đẩy dô chan
		fmt.Println("Chan 2")
		ch <- 42
		wg.Done()
	}(ch)
	wg.Wait()
}

