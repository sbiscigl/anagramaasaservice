package dictionary

import (
	"bufio"
	"log"
	"os"
)

/*Dictionary type for a list of all of the words*/
/*in the english language                       */
type Dictionary struct {
	wordMap map[string]bool
}

/*NewDictionary constructor for dictionary*/
func NewDictionary() *Dictionary {
	return &Dictionary{
		make(map[string]bool, 0),
	}
}

/*NewDictionaryFromFile loads a dictionary from a textfile*/
func NewDictionaryFromFile(filePath string) *Dictionary {
	log.Println("starting to try to poulate map from file")
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	dictMap := make(map[string]bool, 0)
	for scanner.Scan() {
		dictMap[scanner.Text()] = true
	}
	log.Println("map populated")
	return &Dictionary{
		dictMap,
	}
}

/*Update the value map*/
func (d *Dictionary) Update(word string) {
	d.wordMap[word] = true
}

/*DeleteAll entries in the dictionary*/
func (d *Dictionary) DeleteAll() {
	d.wordMap = make(map[string]bool, 0)
}

/*DeleteOne deletes one entry in the map*/
func (d *Dictionary) DeleteOne(word string) {
	delete(d.wordMap, word)
}

/*CheckForWords checks for words in the dictionary*/
func (d *Dictionary) CheckForWords(words []string) []string {
	existingWords := make([]string, 0)
	for _, e := range words {
		if val, ok := d.wordMap[e]; ok == true && val == true {
			existingWords = append(existingWords, e)
		}
	}
	return existingWords
}
