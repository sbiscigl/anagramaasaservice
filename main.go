package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/sbiscigl/anagramaasaservice/dictionary"
	"github.com/sbiscigl/anagramaasaservice/word"
)

func main() {
	fmt.Println("Starting Anagram as service")
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	dictionaryData := dictionary.NewDataFromFile(dir + "/resources/dictionary.txt")
	wordData := word.NewData()
	dictionaryService := dictionary.NewService(dictionaryData)
	wordService := word.NewService(wordData, dictionaryService)
	dictionaryController := dictionary.NewController(dictionaryService)
	dictionarySingleController := dictionary.NewSingleController(dictionaryService)
	wordController := word.NewController(wordService)

	http.Handle(dictionaryController.Path, dictionaryController)
	http.Handle(dictionarySingleController.Path, dictionarySingleController)
	http.Handle(wordController.Path, wordController)
	http.ListenAndServe(":8080", nil)
}
