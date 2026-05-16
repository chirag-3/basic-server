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

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	fmt.Println("This is the main function")
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
