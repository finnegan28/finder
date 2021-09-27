package main

import (
    "fmt"
    "log"
    "net/http"

    "finder/finder"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func returnData(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(findMusic)
}

func main() {
    http.HandleFunc("/find", returnData)
    handleRequests()
}