package main

// import (
// 	"testing"
// 	"os"
// 	"strings"
	
// )

// func TestReplaceTabooWords(t *testing.T) {
// 	tabooWords := []string{"bad", "ugly"}

// 	word1 := "This is a bad example"
// 	expected1 := "This is a *** example"
// 	result1 := replaceTabooWords(word1, tabooWords)
// 	if result1 != expected1 {
// 		t.Errorf("Expected '%s', but got '%s' for word: '%s'", expected1, result1, word1)
// 	}

// 	word2 := "This is a good example"
// 	expected2 := "This is a good example"
// 	result2 := replaceTabooWords(word2, tabooWords)
// 	if result2 != expected2 {
// 		t.Errorf("Expected '%s', but got '%s' for word: '%s'", expected2, result2, word2)
// 	}

// 	word3 := "The cat is ugly and the dog is also ugly"
// 	expected3 := "The cat is **** and the dog is also ****"
// 	result3 := replaceTabooWords(word3, tabooWords)
// 	if result3 != expected3 {
// 		t.Errorf("Expected '%s', but got '%s' for word: '%s'", expected3, result3, word3)
// 	}
// }

// func TestProcessWordsFile(t *testing.T) {
// 	tabooWords := []string{"bad", "ugly"}

// 	fileName := "test_words.txt"
// 	fileContent :="The cat is ugly and the dog is also ugly"
// 	err := createTestFile(fileName, fileContent)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	err = processWordsFile(fileName, tabooWords)
// 	if err != nil {
// 		t.Errorf("Error while processing words file: %v", err)
// 	}

// 	output, err := readTabooWordsFromFile("test_words.txt")
	
// 	if err != nil {
// 		t.Errorf("Error while riding words file: %v", err)
// 	}

	
// 	for i, word := range output {
// 	newWord := replaceTabooWords(word, tabooWords)
// 	output[i] = newWord
// 	}
	
// 	newOutput := strings.Join(output, " ")
// 	expectedOutput :="the cat is **** and the dog is also ****"
		

// 	if newOutput != expectedOutput {
// 		t.Errorf("Expected output:\n%s\n\nActual output:\n%s", expectedOutput, output)
// 	}
// }

// // Helper function to create a temporary file
// func createTestFile(fileName, content string) error {
// 	file, err := os.Create(fileName)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	_, err = file.WriteString(content)
// 	return err
// }

