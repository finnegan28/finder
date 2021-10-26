package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    
    "github.com/gorilla/mux"
)

type searchText struct {
	Search string `json:"search"`
}

func returnData(w http.ResponseWriter, request *http.Request) {
    var search searchText

    decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&search)

	if err != nil {
		panic(err)
	}

    fmt.Fprintf(w, findMusic(search.Search))
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)

    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/find", returnData).Methods("POST")
    http.Handle("/", myRouter)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
    fmt.Println("Finder is now running")

    handleRequests()
}
