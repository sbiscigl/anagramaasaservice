package word

import (
	"encoding/json"
	"fmt"
)

/*Word type for storing a word and its anagrams*/
type Word struct {
	word     string
	Anagrams []string `json:"anagrams"`
}

/*NewWord constructor for word*/
func NewWord(word string, anagrams []string) Word {
	return Word{
		word,
		anagrams,
	}
}

/*ToJSON returns a json format of a word*/
func (w Word) ToJSON() []byte {
	j, err := json.Marshal(w)
	if err != nil {
		fmt.Println("Json marshalling of a error failed")
	}
	return j
}
