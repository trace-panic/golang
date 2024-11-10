package controllers

import (
	"03/pkg/models"
	"03/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()

	res, err := json.Marshal(books)
	if err != nil {
		http.Error(w, "Error marshalling books", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]

	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	book, _ := models.GetBookById(Id)
	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Error marshalling book", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)

	book := CreateBook.CreateBook()
	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Error marshalling created book", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["id"]

	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	book := models.DeleteBook(Id)
	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Error marshalling deleted book", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	UpdateBook := &models.Book{}
	utils.ParseBody(r, UpdateBook)

	vars := mux.Vars(r)
	bookId := vars["id"]

	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	book, db := models.GetBookById(Id)
	if UpdateBook.Name != "" {
		book.Name = UpdateBook.Name
	}
	if UpdateBook.Author != "" {
		book.Author = UpdateBook.Author
	}
	if UpdateBook.Publication != "" {
		book.Publication = UpdateBook.Publication
	}
	db.Save(&book)
	res, err := json.Marshal(book)
	if err != nil {
		http.Error(w, "Error marshalling updated book", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
