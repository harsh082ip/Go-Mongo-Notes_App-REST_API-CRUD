package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/harsh082ip/Go-Mongo-Notes_App-REST_API-CRUD/controllers"
)

const (
	WEBPORT = ":27017"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/getnotes", controllers.GetNotes).Methods("GET")
	r.HandleFunc("/createnote", controllers.CreateNote).Methods("POST")
	r.HandleFunc("/getnote/{id}", controllers.GetNoteById).Methods("GET")
	r.HandleFunc("/updatenote/{id}", controllers.UpdateNote).Methods("POST")
	r.HandleFunc("/deletenote/{id}", controllers.DeleteNote).Methods("GET")

	fmt.Println("Starting Server at port 8083...")
	log.Fatal(http.ListenAndServe(WEBPORT, r))
	// controllers.GetNoteById()
}
