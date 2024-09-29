package book

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type Handler interface {
	GetBook(w http.ResponseWriter, r *http.Request)
	CreateBook(w http.ResponseWriter, r *http.Request)
}

type HandleBook struct {
	useCase UseCase
}

func NewHandlerBook(uc UseCase) Handler {
	return &HandleBook{useCase: uc}
}

func (h HandleBook) GetBook(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	book, err := h.useCase.Get(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(book)
}

func (h HandleBook) CreateBook(w http.ResponseWriter, r *http.Request) {
	var bookRequest struct {
		Title       string `json:"title"`
		Genre       string `json:"genre"`
		Author      string `json:"author"`
		ReleaseDate string `json:"release_date"`
	}

	if err := json.NewDecoder(r.Body).Decode(&bookRequest); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	releaseDate, err := time.Parse("2006-01-02", bookRequest.ReleaseDate)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id, err := h.useCase.Create(bookRequest.Title, bookRequest.Genre, bookRequest.Author, releaseDate)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(strconv.Itoa(id)))

}
