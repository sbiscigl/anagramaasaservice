package dictionary

/*Service service level logic for dictionary*/
type Service struct {
	/*Nothing yet*/
	data *Data
}

/*NewService constructor for service*/
func NewService(data *Data) *Service {
	return &Service{
		data: data,
	}
}

/*Post posts words to the data store*/
func (s *Service) Post(words []string) {
	s.data.Post(words)
}

/*DeleteAll deletes all words from the data store*/
func (s *Service) DeleteAll() {
	s.data.DeleteAll()
}

/*DeleteOne deletes one word from the data store*/
func (s *Service) DeleteOne(word string) {
	s.data.DeleteOne(word)
}

/*CheckForWords checks for words in the dictionary*/
func (s *Service) CheckForWords(words []string) []string {
	return s.data.CheckForWords(words)
}
