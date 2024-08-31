package main

import (
	"fmt"
	"html"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		body, err := io.ReadAll(r.Body)
		if err != nil {
			fmt.Println("Error when parsing body: ", err.Error())
			w.Write([]byte(err.Error()))
			return
		}

		fmt.Printf("Incoming request...\nMethod: %s\nBody: %s\n", r.Method, string(body))
		fmt.Fprintf(w, "Hello, %q\n lengthBody: %d\n", html.EscapeString(r.URL.Path), len(body))
	})

	http.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hi")
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
