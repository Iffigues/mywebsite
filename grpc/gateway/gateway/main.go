package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "time"
)

func main() {

    r := mux.NewRouter()
    srv := &http.Server{
        Handler:      r,
        Addr:         ":9090",
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }
    srv.ListenAndServe()
}
