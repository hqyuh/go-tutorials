package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func sayHi() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock() // đọc xong rồi, mở khóa
	wg.Done()
}

func increment()  {
	counter++
	m.Unlock() // mở
	wg.Done()
}

func main()  {
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock() // khóa lại đọc, không cho thay đổi giá trị
		go sayHi()
		m.Lock() // gọi để ghi
		go increment()
	}
	wg.Wait()
}

