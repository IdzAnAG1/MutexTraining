package RWMutex

import (
	"io"
	"net/http"
	"strings"
	"sync"
)

func Launch(tags ...string) map[string]map[string][]string {
	result := make(map[string]map[string][]string)
	urls := []string{
		"https://go.dev/",
		"https://www.alldevstack.com/ru/golang-coding-conventions/naming-conventions.html",
		"https://www.kp.ru/expert/dom/luchshie-gazovye-plity-s-gazovoj-dukhovkoj/",
		"https://www.kp.ru/expert/elektronika/telefony/smartfony/",
	}
	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)
	for _, el := range urls {

		go func(url string) {
			wg.Add(1)
			defer wg.Done()
			var page Page
			html, err := page.LoadPageHTML(url)
			if err != nil {
				return
			}
			page.SetHTML(html)
			for _, tag := range tags {
				go func(t, url string) {
					wg.Add(1)
					defer wg.Done()

					tagsWithRes := make(map[string][]string)
					arr := page.Parser(t)
					tagsWithRes[t] = arr

					mu.Lock()
					result[url] = tagsWithRes
					mu.Unlock()

				}(tag, url)
			}
		}(el)
	}
	wg.Wait()
	return result
}

type Page struct {
	html string
	mu   sync.RWMutex
}

func (p *Page) Parser(tag string) []string {
	p.mu.RLock()
	defer p.mu.RUnlock()
	tagName := strings.Trim(tag, "<>")
	clTag := "</" + tagName + ">"
	st := 0
	var arr []string
	var line string
	for in, el := range p.html {
		if string(el) == "<" {
			for i, e := range p.html[in:] {
				if string(e) == ">" {
					line = p.html[in : in+i+1]
					break
				}
			}
			if strings.HasPrefix(line, tag[0:len(tag)-2]) {
				st = in + len(line)
			}
			if line == clTag {
				arr = append(arr, p.html[st:in])
			}
		}
	}
	return arr
}

func (p *Page) LoadPageHTML(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	html, errR := io.ReadAll(resp.Body)
	if errR != nil {
		return "", errR
	}
	return string(html), nil
}

func (p *Page) SetHTML(html string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.html = html
}
