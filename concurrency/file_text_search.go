package concurrency

import "fmt"

type fileSearchManager struct {
	path string
}

func NewFileSearchManager(filepath string) *fileSearchManager {
	return &fileSearchManager{path: filepath}
}

func (fs *fileSearchManager) search(target string) {
}

func FileTextSearch() {
	fmt.Println("Text search")
}
