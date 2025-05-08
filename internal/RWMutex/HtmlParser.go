package RWMutex

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func temp() {
	urls := []string{
		"https://go.dev/",
		"https://www.alldevstack.com/ru/golang-coding-conventions/naming-conventions.html",
		"https://www.kp.ru/expert/dom/luchshie-gazovye-plity-s-gazovoj-dukhovkoj/",
		"https://www.kp.ru/expert/elektronika/telefony/smartfony/",
	}

	resp, err := http.Get(urls[2])
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	html, errR := io.ReadAll(resp.Body)
	if errR != nil {
		panic(errR)
	}
	st := 0
	var arr []string
	doc := string(html)
	var line string
	for in, el := range doc {
		if string(el) == "<" {
			for i, e := range doc[in:] {
				if string(e) == ">" {
					line = doc[in : in+i+1]
					break
				}
			}
			if strings.HasPrefix(line, "<script") {
				st = in + len(line)
			}
			if line == "</script>" {
				arr = append(arr, doc[st:in])
			}
		}
	}
	fmt.Println(arr)
}
