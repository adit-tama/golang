package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
	fmt.Fprintf(w, "Heyyyyyyyyyy\n")
	var x [2]interface{}
	var z [2]interface{}
	x[0] = "a"
	x[1] = "b"
	z[0] = "aa"
	z[1] = "bf"
	result := Equal(x, z)
	fmt.Fprint(w, strconv.FormatBool(result))
}

func Equal(fromDb, fetched [2]interface{}) bool {

	if fromDb[0] == fetched[0] && fromDb[1] == fetched[1] {
		return true
	} else if fromDb[1] == fetched[0] && fromDb[0] == fetched[1] {
		return true
	} else {
		return false
	}
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
