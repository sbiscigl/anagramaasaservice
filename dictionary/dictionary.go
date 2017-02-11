package dictionary

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
/*TODO: a really space intensive function to load the*/
/*entire english language frome a textfile					 */
func NewDictionaryFromFile(filePath string) *Dictionary {
	return &Dictionary{
		make(map[string]bool, 0),
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
