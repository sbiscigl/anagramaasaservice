package dictionary

/*Data a data access object for the english dictionary*/
type Data struct {
	dictionary *Dictionary
}

/*NewData constructor for Data type*/
func NewData() *Data {
	return &Data{
		NewDictionary(),
	}
}

/*NewDataFromFile constructorfor data layer and*/
/*populates from a text file*/
func NewDataFromFile(filepath string) *Data {
	return &Data{
		NewDictionaryFromFile(filepath),
	}
}

/*Post word to map*/
func (d *Data) Post(words []string) {
	for _, e := range words {
		d.dictionary.Update(e)
	}
}

/*DeleteAll entries in the dictionary*/
func (d *Data) DeleteAll() {
	d.dictionary.DeleteAll()
}

/*DeleteOne entries in the dictionary*/
func (d *Data) DeleteOne(word string) {
	d.dictionary.DeleteOne(word)
}

/*CheckForWords checks for words in the dictionary*/
func (d *Data) CheckForWords(words []string) []string {
	return d.dictionary.CheckForWords(words)
}
