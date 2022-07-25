package main

import (
    "fmt"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "time"
)

func main() {

    r := mux.NewRouter()

    r.HandleFunc("/fortune", index).Methods("GET")
    buildHandler := http.FileServer(http.Dir("/build"))
    r.PathPrefix("/").Handler(buildHandler)
    srv := &http.Server{
        Handler:      r,
        Addr:         ":8080",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }

    fmt.Println("Server started on PORT 8080")
    log.Fatal(srv.ListenAndServe())

}

func index(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "build/index.html")
}
