package word

import (
	"net/http"
	"strings"

	"github.com/sbiscigl/anagramaasaservice/resterror"
)

/*Controller type for controller for dictionary requests*/
type Controller struct {
	Path string
	ws   *Service
}

/*NewController constructor for word controller*/
func NewController(ws *Service) *Controller {
	return &Controller{
		"/anagrams/",
		ws,
	}
}

/*ServeHTTP implimentation for http handler interface*/
func (wc *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		word := r.URL.Path
		if strings.HasSuffix(word, ".json") {
			word = strings.TrimSuffix(word, ".json")
			word = strings.TrimPrefix(word, "/anagrams/")
			w.WriteHeader(200)
			w.Write(wc.ws.GetAnagrams(word).ToJSON())
		} else {
			w.WriteHeader(400)
			w.Write(resterror.NewError("must be in foramt /words/{your word}.json", 400).ToJSON())
		}
	} else {
		w.WriteHeader(400)
		w.Write(resterror.NewError("Http method not accepted", 400).ToJSON())
	}
}
