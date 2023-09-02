package main

import (
	work "demo/working"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("system started")
	router := mux.NewRouter()
	router.HandleFunc("/login", work.Login).Methods("POST")
	router.HandleFunc("/register", work.Register).Methods("POST")
	router.HandleFunc("/admin", work.Admin).Methods("POST")
	router.HandleFunc("/addtip", work.AddTip).Methods("POST")
	err := http.ListenAndServe(":8081", router)
	if err != nil {
		log.Fatal(err)
	}

}
