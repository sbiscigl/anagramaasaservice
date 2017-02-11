package dictionary

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/sbiscigl/anagramaasaservice/error"
)

/*Controller type for controller for dictionary requests*/
type Controller struct {
	Path string
	ser  *Service
}

/*NewController constructor for dictionary controller*/
func NewController(service *Service) *Controller {
	return &Controller{
		"/words.json",
		service,
	}
}

/*BodyJSON why cant this be in a function*/
// type BodyJSON struct {
// 	Words []string `json:"words"`
// }

func deserializePostWordData(body io.ReadCloser) ([]string, int) {
	defer body.Close()
	var words struct {
		Words []string `json:"words"`
	}
	v, err := ioutil.ReadAll(body)
	json.Unmarshal(v, &words)
	e := 0
	if err != nil {
		fmt.Println(err.Error())
		e = 1
	}
	return words.Words, e
}

/*ServeHTTP implimentation for http handler interface*/
func (dc *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		dc.ser.DeleteAll()
		w.WriteHeader(204)
	} else if r.Method == "POST" {
		words, err := deserializePostWordData(r.Body)
		fmt.Println(words)
		if err == 0 {
			dc.ser.Post(words)
			w.WriteHeader(201)
		} else {
			w.WriteHeader(400)
			w.Write(error.NewError("Json deserialization wasnt fun", 400).ToJSON())
		}
	} else {
		w.WriteHeader(400)
		w.Write(error.NewError("Http method not accepted", 400).ToJSON())
	}
}
