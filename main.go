package main

import (
	"fmt"
	"net/http"
	"pkg/handlers"
)

const PortNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//		n, err := fmt.Fprintf(w, "hello world")
	//		if err != nil {
	//test
	//			fmt.Println(err)
	//		}
	//		fmt.Println(fmt.Sprintf("Number ofbytes written %d", n))
	//	})
	fmt.Println("Application starting on Port", PortNumber)
	_ = http.ListenAndServe(PortNumber, nil)
}
