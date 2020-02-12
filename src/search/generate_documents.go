package search

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
)

const OutputFolder = "../out"

type EnglishDictionary struct {
	words []string
}

func CreateRandomTextFiles(numDocs, numWordsPerDoc int) {
	englishDictionary := loadEnglishDictionary()

	fmt.Println("Cleaning output folder")

	err := os.RemoveAll(OutputFolder)

	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(OutputFolder, 0777)

	if err != nil {
		panic(err)
	}

	fmt.Println("Randomly generating files")

	for i := 0; i < numDocs; i++ {
		documentName := OutputFolder + "/" + strconv.Itoa(i) + ".txt"

		if i%10 == 0 {
			fmt.Printf("Finished %d files\n", i)
		}

		createdFile, err := os.Create(documentName)
		defer createdFile.Close()

		if err != nil {
			panic(err)
		}

		for j := 0; j < numWordsPerDoc; j++ {
			randomWordIndex := rand.Intn(len(englishDictionary.words))

			_, err := createdFile.WriteString(englishDictionary.words[randomWordIndex] + "\n")

			if err != nil {
				panic(err)
			}
		}
	}

	fmt.Println("Finished randomly generating files")
}

func loadEnglishDictionary() *EnglishDictionary {
	fmt.Println("Loading English dictionary into memory")

	words := make([]string, 10000)

	englishDictionary := &EnglishDictionary{words: words}

	englishWordFile, err := os.Open("../res/words.txt")

	if err != nil {
		panic(err)
	}

	defer englishWordFile.Close()

	scanner := bufio.NewScanner(englishWordFile)

	for scanner.Scan() {
		line := scanner.Text()

		englishDictionary.words = append(englishDictionary.words, line)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println("Finished loading English dictionary into memory")

	return englishDictionary
}
