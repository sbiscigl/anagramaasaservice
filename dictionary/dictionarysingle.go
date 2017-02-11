package dictionary

import (
	"net/http"
	"strings"

	"github.com/sbiscigl/anagramaasaservice/error"
)

/*SingleController type for controller for dictionary requests*/
type SingleController struct {
	Path string
	ser  *Service
}

/*NewSingleController constructor for dictionary controller*/
func NewSingleController(service *Service) *SingleController {
	return &SingleController{
		"/words/",
		service,
	}
}

func (dc *SingleController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		/*TODO: Get word from path and delete*/
		word := r.URL.Path
		if strings.HasSuffix(word, ".json") {
			word = strings.TrimSuffix(word, ".json")
			word = strings.TrimPrefix(word, "/words/")
			dc.ser.DeleteOne(word)
			w.WriteHeader(200)
		} else {
			w.WriteHeader(400)
			w.Write(error.NewError("must be in foramt /words/{your word}.json", 400).ToJSON())
		}
	} else {
		w.WriteHeader(400)
		w.Write(error.NewError("Http method not accepted", 400).ToJSON())
	}
}
