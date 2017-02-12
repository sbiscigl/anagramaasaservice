package word

import (
	"log"

	"github.com/sbiscigl/anagramaasaservice/dictionary"
)

/*Service layer for word endpoint*/
type Service struct {
	data      *Data
	dcService *dictionary.Service
}

/*NewService constructor for word service*/
func NewService(data *Data, dcService *dictionary.Service) *Service {
	return &Service{
		data,
		dcService,
	}
}

/*GetAnagrams gets all the anagrams for a word*/
func (ws *Service) GetAnagrams(word string) Word {
	words := getAllPerms(word)
	log.Println("made all permutations")
	log.Println(words)
	words = filter(words, func(arg2 string) bool {
		if arg2 == word {
			return false
		}
		return true
	})
	realWorlds := ws.dcService.CheckForWords(words)
	return NewWord(word, realWorlds)
}

func filter(s []string, fn func(string) bool) []string {
	var p []string // == nil
	for _, v := range s {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}

/*Get all possible perms*/
func getAllPerms(word string) []string {
	log.Println("getting all permutations")
	return Permutations(word)
}

/*Permutations recursive get all instances of string*/
func Permutations(input string) []string {
	log.Println("recursed")
	if len(input) == 1 {
		return []string{input}
	}

	runes := []rune(input)
	subPermutations := Permutations(string(runes[0 : len(input)-1]))

	result := []string{}
	for _, s := range subPermutations {
		result = append(result, merge([]rune(s), runes[len(input)-1])...)
	}

	return result
}

func merge(ins []rune, c rune) (result []string) {
	for i := 0; i <= len(ins); i++ {
		result = append(result, string(ins[:i])+string(c)+string(ins[i:]))
	}
	return
}
