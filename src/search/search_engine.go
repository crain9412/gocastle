package search

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"runtime"
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

func (searchEngine *SearchEngine) Query(query string) ([]string, bool) {
	result, ok := searchEngine.index.data[query]

	if ok {
		return result, true
	}

	return nil, false
}

func (searchEngine *SearchEngine) Print() {
	for key, value := range searchEngine.index.data {
		fmt.Printf("%s=%v\n", key, value)
		break
	}
}

func (searchEngine *SearchEngine) PrintTopTen() {
	fmt.Println("Printing top ten most frequent terms")

	currentFreqs := make(map[string]int, 10)

	i := 0

	for key, value := range searchEngine.index.data {
		if i < 10 {
			currentFreqs[key] = len(value)
		} else {
			min, evict := getMinThreshold(currentFreqs)
			if len(value) > min {
				delete(currentFreqs, evict)
				currentFreqs[key] = len(value)
			}
		}
		i++
	}

	for key := range currentFreqs {
		fmt.Printf("%s=%v\n", key, searchEngine.index.data[key])
	}
}

func getMinThreshold(currentFreqs map[string]int) (int, string) {
	min := math.MaxInt64
	evict := "Placeholder"

	for key, value := range currentFreqs {
		if evict == "Placeholder" {
			/* if everything ties evict a random one */
			evict = key
		}

		if value < min {
			min = value
			evict = key
		}
	}

	return min, evict
}

func createInverseIndex() *InverseIndex {
	inverseIndex := &InverseIndex{data: make(map[string][]string, 100)}

	fmt.Println("Creating inverse index")

	PrintMemUsage()

	i := 0

	err := filepath.Walk(IndexFolder,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if i%10 == 0 {
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

					if line == " " || line == "" {
						continue
					}

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

	PrintMemUsage()

	return inverseIndex
}

// PrintMemUsage outputs the current, total and OS memory being used. As well as the number
// of garage collection cycles completed.
func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
