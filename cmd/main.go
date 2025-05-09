package main

import (
	"MutexTraining/internal/RWMutex"
	"MutexTraining/internal/mutex_tr"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	res := RWMutex.Launch("<h1>", "<script>")
	jd, err := json.MarshalIndent(res, "", " ")
	if err != nil {
		fmt.Println("Ошибка js")
	}
	err = os.WriteFile("../internal/RWMutex/Output.json", jd, 0644)
	if err != nil {
		fmt.Println("Ошибка записи в файл", err)
	}
	end := time.Now()
	elap := end.Sub(start)
	fmt.Println(elap)
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
