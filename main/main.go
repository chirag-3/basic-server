package main

import (
	"fmt"
	"log"
	"net/http"
)

// func helloHandler1() string {
// 	return "Thi is hello page"
// }

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello my name is Chirag")
}

func detailsHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Fprintf(w, "%s stays at %s", name, address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	fmt.Println("This is the main function")
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/submitDetails", detailsHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
