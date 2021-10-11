package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type ToDo struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	Status string `json:"status"`
}

var ToDos []ToDo

func getToDos(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getToDos")
	json.NewEncoder(w).Encode(ToDos)
}

func getToDoById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: getToDoById")
	vars := mux.Vars(r)
	id := vars["id"]

	for _, todo := range ToDos {
		if todo.Id == id {
			json.NewEncoder(w).Encode(todo)
		}
	}
}

func createNewToDo(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var todo ToDo
	json.Unmarshal(reqBody, &todo)

	ToDos = append(ToDos, todo)

	json.NewEncoder(w).Encode(todo)
}

func deleteToDo(w http.ResponseWriter, r *http.Request) {
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)
	// we will need to extract the `id` of the article we
	// wish to delete
	id := vars["id"]

	// we then need to loop through all our articles
	for index, todo := range ToDos {
		// if our id path parameter matches one of our
		// articles
		if todo.Id == id {
			// updates our Articles array to remove the
			// article
			ToDos = append(ToDos[:index], ToDos[index+1:]...)
		}
	}

}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/todos", getToDos)
	myRouter.HandleFunc("/todos/{id}", getToDoById)
	myRouter.HandleFunc("/todo", createNewToDo).Methods("POST")
	myRouter.HandleFunc("/todo/{id}", deleteToDo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	ToDos = []ToDo{
		{Id: "1", Title: "Todo 1", Desc: "Work Work Work", Status: "InProgress"},
		{Id: "2", Title: "Todo 2", Desc: "Work much harder", Status: "Backlog"},
	}

	handleRequests()
}
