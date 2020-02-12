package search

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"fmt"
)

type SearchEngine struct {
	index *InverseIndex
}

const IndexFolder = "../out"

type InverseIndex struct {
	data map[string][]string
}

func InitSearchEngine() *SearchEngine {
	return &SearchEngine{index: createInverseIndex()}
}

func (searchEngine *SearchEngine) Query(query string) []string {
	result, ok := searchEngine.index.data[query]

	if ok {
		return result
	} else {
		return make([]string, 1)
	}
}

func (searchEngine *SearchEngine) Print() {
	for key, value := range searchEngine.index.data {
		fmt.Printf("%s=%v\n", key, value)
		break
	}
}

func createInverseIndex() *InverseIndex {
	inverseIndex := &InverseIndex{data: make(map[string][]string, 100)}

	fmt.Println("Creating inverse index")

	i := 0

	err := filepath.Walk(IndexFolder,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if i % 10 == 0 {
				fmt.Printf("Finished %d files\n", i)
			}

			if !info.IsDir() {
				currentFile, err := os.Open(path)

				if err != nil {
					panic(err)
				}

				defer currentFile.Close()

				scanner := bufio.NewScanner(currentFile)

				for scanner.Scan() {
					line := scanner.Text()

					currentDocuments, ok := inverseIndex.data[line]

					allDocuments := make([]string, 1)

					if ok {
						allDocuments = append(currentDocuments, path)
					}

					inverseIndex.data[line] = allDocuments
				}

				if scanner.Err() != nil {
					panic(scanner.Err())
				}
			}

			i++

			return nil
		})

	if err != nil {
		log.Println(err)
	}

	fmt.Println("Finished creating inverse index")

	return inverseIndex
}