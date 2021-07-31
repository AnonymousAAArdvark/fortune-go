package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var files []string

func randomInt(min, max int) int {
	return min + rand.Intn(max - min)
}

func main() {
	var files []string

	root, _ := filepath.Abs("../fortune/fortunes")

	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if f.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())
	i := randomInt(0, len(files))
	randomFile := files[i]
	file, err := os.Open(randomFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	quotes := string(b)

	quotesSlice := strings.Split(quotes, "%")
	j := randomInt(0, len(quotesSlice) - 1)

	fmt.Println(quotesSlice[j][1:len(quotesSlice[j])-1])
}