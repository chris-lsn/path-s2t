package main

import (
	"os"
	"path/filepath"
	"sync"

	"github.com/longbridgeapp/opencc"
)

func convertPath(converter *opencc.OpenCC, p string, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()
	items, err := os.ReadDir(p)

	if (err != nil) {
		panic(err)
	}
	
	for _, item := range items {
		wg.Add(1)
		go func(item os.DirEntry) {
			defer wg.Done()
			newName, _ := converter.Convert(item.Name())
			oldPath := filepath.Join(p, item.Name())
			newPath := filepath.Join(p, newName)
			err := os.Rename(oldPath, newPath)
			if (err != nil) {
				println("Fail to rename file: " + item.Name())
			} else {
				println(newPath)
			}

			if (item.IsDir()) {
				go convertPath(converter, newPath, wg)
			}
		}(item)	
	}
}

func main() {
	path := os.Args[1]

	if path == "" {
		panic("Path is required")
	}
	converter, err := opencc.New("s2t")
	if err != nil {
		panic(err)
	}

	wg := &sync.WaitGroup{}
	convertPath(converter, path, wg)
	wg.Wait()
}