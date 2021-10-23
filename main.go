package main

import (
    "fmt"
    "log"
    "net/http"
    //"encoding/json"
    "github.com/gorilla/mux"
    //"io/ioutil"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)

    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/find", returnData).Methods("POST")
    http.Handle("/", myRouter)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func returnData(w http.ResponseWriter, r *http.Request){
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    search := r.FormValue("search")
    fmt.Fprintf(w, findMusic(search))
}

func main() {
    fmt.Println("Rest API v2.0 - Mux Routers")

    handleRequests()
}