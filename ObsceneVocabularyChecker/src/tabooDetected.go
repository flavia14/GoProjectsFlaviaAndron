package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	tabooFileName := "../tabooWords.txt"
// 	wordsFileName := "../words.txt"

// 	tabooWords, err := readTabooWordsFromFile(tabooFileName)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	err = processWordsFile(wordsFileName, tabooWords)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}

// 	fmt.Println("Bye!")
// }

// func processWordsFile(fileName string, tabooWords []string) error {
// 	file, err := os.Open(fileName)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		word := scanner.Text()
//         wordLower := strings.ToLower(word)
    
// 		replacedWord := replaceTabooWords(wordLower, tabooWords)
// 		fmt.Println(replacedWord)
// 	}

// 	if err := scanner.Err(); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func replaceTabooWords(word string, tabooWords []string) string {

//     isTaboo := 0
// 	for _, taboo := range tabooWords {
// 		if strings.Contains(word, taboo) {
// 			replacement := strings.Repeat("*", len(taboo))
// 			word = strings.ReplaceAll(word, taboo, replacement)
//             isTaboo = 1
// 		}
// 	}

//     if isTaboo == 0 {
//         fmt.Println("Your text is cute!")
//     }
// 	return word
// }
