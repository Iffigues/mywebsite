package main

import  (
	"net/http"
	"time"
	"log"
	"fmt"
	"io/ioutil"
)

func Handler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

func main() {

	r := http.NewServeMux()
	webapp := http.FileServer(http.Dir("/build"))
	r.Handle("/", Handler(webapp))
	files, err := ioutil.ReadDir("/build")
	if err != nil {
	    fmt.Println(err)
	}
	for _, f := range files {
            fmt.Println(f.Name())
	}
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}
