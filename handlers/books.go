package handlers

import (
	"Lib-CURD/models"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

var books []models.Book 


func GetBooks(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(books)
}


func CreateBook(w http.ResponseWriter, r *http.Request) {
    var book models.Book
    _ = json.NewDecoder(r.Body).Decode(&book)
    books = append(books, book)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(book)
}


func GetBookByID(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for _, item := range books {
        if item.BookId == params["id"] {
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    http.Error(w, "Book not found", http.StatusNotFound)
}


func UpdateBook(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var updatedBook models.Book
    _ = json.NewDecoder(r.Body).Decode(&updatedBook)
    for index, item := range books {
        if item.BookId == params["id"] {
            books[index] = updatedBook
            w.Header().Set("Content-Type", "application/json")
            json.NewEncoder(w).Encode(updatedBook)
            return
        }
    }
    http.Error(w, "Book not found", http.StatusNotFound)
}


func DeleteBook(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    for index, item := range books {
        if item.BookId == params["id"] {
            books = append(books[:index], books[index+1:]...)
            w.Header().Set("Content-Type", "application/json")
            w.WriteHeader(http.StatusOK)
            response := map[string]string{"message": "Book deleted successfully"}
            json.NewEncoder(w).Encode(response)
            return
        }
    }
    http.Error(w, "Book not found", http.StatusNotFound)
}



func SearchBooks(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	resultsChannel := make(chan []models.Book)
	chunkSize := len(books) 
	for i := 0; i < 4; i++ {
		go searchChunk(i, chunkSize, query, resultsChannel)
	}
	var finalResults []models.Book
	for i := 0; i < 4; i++ {
		chunkResults := <-resultsChannel
		finalResults = append(finalResults, chunkResults...)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(finalResults)
}
func searchChunk(chunkIndex int, chunkSize int, query string, ch chan []models.Book) {
	start := chunkIndex * chunkSize
	end := start + chunkSize
	if end > len(books) {
		end = len(books)
	}
	var chunkResults []models.Book
	for _, book := range books[start:end] {
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(book.Description), strings.ToLower(query)) {
			chunkResults = append(chunkResults, book)
		}
	}

	ch <- chunkResults
}
