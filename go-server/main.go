package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintln(w, "POST request successful")
	name := r.FormValue("name")
	planet := r.FormValue("planet")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Planet = %s\n", planet)
}

func cookieHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/cookie" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "The cookie jar is empty...")
}
func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/cookie", cookieHandler)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
