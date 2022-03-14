package shortenedurl

import (
	"encoding/json"
	"net/http"
	"text/template"
	"url-shortener/usecases"
	shortenedurl "url-shortener/usecases/shortenedUrl"

	"github.com/gorilla/mux"
)

type IShortenedUrlController interface {
	Create(w http.ResponseWriter, r *http.Request)
	FindByShortenedUrl(w http.ResponseWriter, r *http.Request)
	Index(w http.ResponseWriter, r *http.Request)
}

type controllers struct {
	usecases *usecases.Container
}

func New(usecases *usecases.Container) IShortenedUrlController {
	return &controllers{usecases: usecases}
}

func (ctr *controllers) Create(w http.ResponseWriter, r *http.Request) {
	var data shortenedurl.ICreateShortenedUrlDTO
	json.NewDecoder(r.Body).Decode(&data)

	newUrl, err := ctr.usecases.ShortenedUrl.Create(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUrl)
}

func (ctr *controllers) FindByShortenedUrl(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	shortenedID := params["id"]

	shortenedUrlByID, err := ctr.usecases.ShortenedUrl.FindByShortenedID(shortenedID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.Redirect(w, r, "https://"+shortenedUrlByID.Url, http.StatusMovedPermanently)
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func (ctr *controllers) Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", nil)
}
