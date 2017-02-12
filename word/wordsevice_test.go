package word

import (
	"fmt"
	"testing"

	"github.com/sbiscigl/anagramaasaservice/dictionary"
)

func TestPermuations(t *testing.T) {
	data := NewData()
	dcData := dictionary.NewDataFromFile("/Users/samuelbisciglia/go/src/github.com/sbiscigl/anagramaasaservice/resources/dictionary.txt")
	dcService := dictionary.NewService(dcData)
	b := NewService(data, dcService)
	a := b.GetAnagrams("read")
	fmt.Println(a)
}
