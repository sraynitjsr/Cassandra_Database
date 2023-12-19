package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sraynitjsr/gocql/gocql"
	"github.com/sraynitjsr/gorilla/mux"
)

var (
	cluster *gocql.ClusterConfig
	session *gocql.Session
)

type Book struct {
	ID     gocql.UUID `json:"id"`
	Title  string     `json:"title"`
	Author string     `json:"author"`
}

func init() {
	cluster = gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "library"
	cluster.Consistency = gocql.Quorum

	var err error
	session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal("Error creating Cassandra session:", err)
	}

	err = session.Query(`
		CREATE TABLE IF NOT EXISTS books (
			id UUID PRIMARY KEY,
			title TEXT,
			author TEXT
		)
	`).Exec()
	if err != nil {
		log.Fatal("Error creating 'books' table:", err)
	}
}

func main() {
	defer session.Close()
	router := mux.NewRouter()
	router.HandleFunc("/books", getBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", createBook).Methods("POST")
  	log.Fatal(http.ListenAndServe(":8080", router))
}

func parseJSON(r *http.Request, data interface{}) error {
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(data)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	var books []Book
	iter := session.Query("SELECT id, title, author FROM books").Iter()
	var id gocql.UUID
	var title, author string

	for iter.Scan(&id, &title, &author) {
		books = append(books, Book{ID: id, Title: title, Author: author})
	}

	if err := iter.Close(); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := gocql.ParseUUID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	var book Book
	err = session.Query("SELECT id, title, author FROM books WHERE id = ?", id).
		Scan(&book.ID, &book.Title, &book.Author)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Book not found")
		return
	}

	respondWithJSON(w, http.StatusOK, book)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	if err := parseJSON(r, &book); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	book.ID = gocql.TimeUUID()

	if err := session.Query(`
		INSERT INTO books (id, title, author) VALUES (?, ?, ?)`,
		book.ID, book.Title, book.Author).Exec(); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, book)
}
