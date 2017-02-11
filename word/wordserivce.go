package word

import "github.com/sbiscigl/anagramaasaservice/dictionary"

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
	return NewWord(word, ws.dcService.CheckForWords(getAllPerms(word)))
}

/*Get all possible perms*/
func getAllPerms(word string) []string {
	/*TODO: a really computationally expensive step to get*/
	/*all string permutations															*/
	return []string{"ok", "cool", "placeholder"}
}
