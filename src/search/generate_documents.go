package search

import (
	"bufio"
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

	err := os.RemoveAll(OutputFolder)

	if err != nil {
		panic(err)
	}

	err = os.MkdirAll(OutputFolder, 0777)

	if err != nil {
		panic(err)
	}

	for i := 0; i < numDocs; i++ {
		documentName := OutputFolder + "/" + strconv.Itoa(i) + ".txt"

		createdFile, err := os.Create(documentName)
		defer createdFile.Close()

		if err != nil {
			panic(err)
		}

		for j := 0; j < numWordsPerDoc; j++ {
			randomWordIndex := rand.Intn(len(englishDictionary.words))

			_, err := createdFile.WriteString(englishDictionary.words[randomWordIndex] + " ")

			if err != nil {
				panic(err)
			}
		}
	}
}

func loadEnglishDictionary() *EnglishDictionary {
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

	return englishDictionary
}
