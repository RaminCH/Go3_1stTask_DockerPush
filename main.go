package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Item struct {
	ID      string `json:"id"`
	Content string `json:"content"`
}

//Items ...
var Items []Item = []Item{
	Item{"1", "abcd"},
}

//ErrorMessage ...
type ErrorMessage struct {
	Message string `json:"Message"`
}

//GetItems ...
func GetItems(w http.ResponseWriter, r *http.Request) {
	find := false
		if !find {
			w.WriteHeader(http.StatusNotFound) 
			var erM = ErrorMessage{Message: "Error No one items in stock!"}
			json.NewEncoder(w).Encode(erM)
		} else {
			json.NewEncoder(w).Encode(Items)
		}
}

//GetItemID ...
func GetItemID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	find := false
	for _, item := range Items {
		if item.ID == id {
			find = true
			json.NewEncoder(w).Encode(item)
		}
	}
	if !find {
		w.WriteHeader(http.StatusNotFound) // Изменить статус код запроса на 404
		var erM = ErrorMessage{Message: "Error: Item with that id not found!"}
		json.NewEncoder(w).Encode(erM)
	}
}

//PostItem ...
func PostItem(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var item Item
	json.Unmarshal(reqBody, &item)
	w.WriteHeader(http.StatusCreated)
	Items = append(Items, item)
}

//PutItemID ...
func PutItemID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	find := false

	for index, item := range Items {
		if item.ID == id {
			find = true
			reqBody, _ := ioutil.ReadAll(r.Body)
			w.WriteHeader(http.StatusAccepted)    
			json.Unmarshal(reqBody, &Items[index])
		}
	}

	if !find {
		w.WriteHeader(http.StatusNotFound) // Изменить статус код запроса на 404
		var erM = ErrorMessage{Message: "Error: Item with that id not found!"}
		json.NewEncoder(w).Encode(erM)
	}

}

//DeleteItemID ...
func DeleteItemID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	find := false

	for index, item := range Items {
		if item.ID == id {
			find = true
			w.WriteHeader(http.StatusAccepted) 
			Items = append(Items[:index], Items[index+1:]...)
		}
	}
	if !find {
		w.WriteHeader(http.StatusNotFound) // Изменить статус код запроса на 404
		var erM = ErrorMessage{Message: "Error: Item with that id not found!"}
		json.NewEncoder(w).Encode(erM)
	}

}

func main() {
	
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/items", GetItems).Methods("GET")
	router.HandleFunc("/item/{id}", GetItemID).Methods("GET")

	router.HandleFunc("/item", PostItem).Methods("POST")

	router.HandleFunc("/item/{id}", PutItemID).Methods("PUT")

	router.HandleFunc("/item/{id}", DeleteItemID).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}