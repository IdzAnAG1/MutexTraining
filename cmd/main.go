package main

import (
	"MutexTraining/internal/RWMutex"
	"MutexTraining/internal/mutex_tr"
	"fmt"
	"time"
)

func main() {
	res := RWMutex.Launch("<h1>", "<script>")
	fmt.Println(res)
}
func testLS_P() {
	start := time.Now()
	err := mutex_tr.Writer(mutex_tr.Reader(), "../internal/mutex_tr/outputs/")
	if err != nil {
		return
	}
	end := time.Now()
	elap := end.Sub(start)
	fmt.Println(elap)
}
