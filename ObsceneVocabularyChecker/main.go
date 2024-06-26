package main

import (
	"fmt"
	"gitlab.pitechplus.com/dragos.contiu/goprojects/-/tree/main/FlaviaAndron/ObsceneVocabularyChecker/processPack"
	"gitlab.pitechplus.com/dragos.contiu/goprojects/-/tree/main/FlaviaAndron/ObsceneVocabularyChecker/readPack"
)

func main() {
	tabooFileName := "tabooWords.txt"
	wordsFileName := "words.txt"

	tabooWords, err := readPack.ReadTabooWordsFromFile(tabooFileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = processPack.ProcessWordsFile(wordsFileName, tabooWords)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Bye!")
}