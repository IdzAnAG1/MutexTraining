package mutex_tr

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sync"
)

func Reader() map[string]int {
	letters := make(map[string]int)
	var (
		wg sync.WaitGroup
		mu sync.Mutex
	)
	text, R_Err := fromFileToBuffer("../internal/mutex_tr/inputs/text.txt")
	if R_Err != nil {
		fmt.Printf("Error at : %v", R_Err)
	}
	for i := 0x0041; i <= 0x044F; i++ {
		if (i > 0x007A && i < 0x0410) || (i > 0x005A && i < 0x0061) {
			continue
		}
		wg.Add(1)
		go func(index int, buf []rune) {
			defer wg.Done()
			c := letterCounter(buf, index)
			mu.Lock()
			letters[string(rune(index))] = c
			mu.Unlock()
		}(i, text)
	}
	wg.Wait()
	return letters
}

func fromFileToBuffer(path string) ([]rune, error) {
	var (
		buffer []rune
	)
	file, foe := os.Open(path) // foe - File Opening Error
	if foe != nil {
		fmt.Printf("File Opening Error %v\n", foe)
		return []rune{}, foe
	}
	defer file.Close()
	scanner := bufio.NewReader(file)
	for {
		r, _, pfe := scanner.ReadRune() // pfe - Parsing File Error
		if pfe == io.EOF {
			break
		}
		buffer = append(buffer, r)
	}
	return buffer, nil
}

func letterCounter(buffer []rune, letter int) int {
	var counter int
	for _, val := range buffer {
		if val == rune(letter) {
			counter++
		}
	}
	return counter
}

func Writer(data map[string]int, path string) error {
	jsonData, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		fmt.Printf("Ошибка при приведении формата в JSON")
		return err
	}
	wfe := os.WriteFile(path+"Output_FirstWriteInFile_Failed.json", jsonData, 0644) // wfe - Write in File Error
	if wfe != nil {
		fmt.Printf("Ошибка записи в файл : %v\n", err)
		return wfe
	}
	fmt.Println("Данные записаны в файл")
	return nil
}
