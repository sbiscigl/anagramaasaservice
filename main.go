package main

import (
	"fmt"
	"net/http"

	"github.com/sbiscigl/anagramaasaservice/dictionary"
	"github.com/sbiscigl/anagramaasaservice/word"
)

func main() {
	fmt.Println("Starting Anagram as service")

	dictionaryData := dictionary.NewData()
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
