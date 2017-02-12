package word

import (
	"fmt"
	"os"
	"testing"

	"github.com/sbiscigl/anagramaasaservice/dictionary"
)

func TestPermuations(t *testing.T) {
	data := NewData()
	path := os.Getenv("GOPATH")
	dcData := dictionary.NewDataFromFile(path + "/src/github.com/sbiscigl/anagramaasaservice/resources/dictionary.txt")
	dcService := dictionary.NewService(dcData)
	b := NewService(data, dcService)
	a := b.GetAnagrams("read")
	fmt.Println(a)
}
