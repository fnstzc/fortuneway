package main

import (
	"fmt"
	_ "fortuneway/routers"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func main() {

}

func testPressure() {

	fmt.Println("begin")

	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func(i int) {
			if err := recover(); err != nil {
				fmt.Println("goroutine error")
			}
			fmt.Println("thread ",i," prepared ")
			wg.Wait()
			sendRequest()
			fmt.Println("thread ",i," end ")
		}(i)
		wg.Done()
	}
	fmt.Println("wait countdown")

	time.Sleep(10 * time.Second)
}

func sendRequest() {
	resp, err := http.Get("127.0.0.1:8080/hello")
	if  err != nil {
		fmt.Println("query failed")
	}
	if resp == nil {
		fmt.Println("resp is nil")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("transfer body failed")
	} else {
		fmt.Println(string(body))
	}
}


