package main

import (
	"fmt"
	"net/http"
)

const PortNumber = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	//		n, err := fmt.Fprintf(w, "hello world")
	//		if err != nil {
	//			fmt.Println(err)
	//		}
	//		fmt.Println(fmt.Sprintf("Number ofbytes written %d", n))
	//	})
	fmt.Println("Application starting on Port")
	_ = http.ListenAndServe(PortNumber, nil)
}
