package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//Define a todo structure
type ToDo struct {
	Title string `json:"title"`
	Description string `json:"description"`
}

//Create todos variable that stores an array of todos 
var todos []ToDo

func main() {
	http.HandleFunc("/", ToDoListHandler)
	fmt.Println("Backend running on 8080 port")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func ToDoListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")


	//Preflight Request, checks if the request from the frontend is a valid option ie a GET, POST etc 
	if r.Method == http.MethodOptions{
		w.WriteHeader(http.StatusOK)
		return 
	}

	if r.Method == http.MethodGet{
		w.Header().Set("Content-Type" , "application/json")
		json.NewEncoder(w).Encode(todos)
	}else if r.Method == http.MethodPost{
		//Create the newTodo variable 
		var newTodo ToDo

		//Decode also us to do newTodo.title and newTodo.desc 
		err := json.NewDecoder(r.Body).Decode(&newTodo)
		if err != nil {
			http.Error(w, "Error when decoding request body data", http.StatusBadRequest)
			return
		}

		//Add the newTodo into the array with the other todos 
		todos = append(todos, newTodo)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(todos)
	}else{
		http.Error(w, "Invalid Method Option", http.StatusBadRequest)
		return
	}
}
