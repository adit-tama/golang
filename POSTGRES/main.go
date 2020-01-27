package main

import (
	"fmt"
	"net/http"

	postgres "POSTGRES/utility/postgres"

	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
	fmt.Fprintf(w, "Heyyyyyyyyyy\n")
	postgres.GetData()
}

func main() {

	//handle routing
	router := mux.NewRouter()
	router.HandleFunc("/", hello).Methods("GET")
	http.Handle("/", router)

	//init server
	fmt.Println("Server is Running at Port 3200")
	err := http.ListenAndServe(":3200", router)
	if err != nil {
		fmt.Print("Running Server Error: ", err)
	}
}
