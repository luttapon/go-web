package main

import (
	"fmt"
	"net/http"
)

func hellogoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go!")
}
func hiGoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, Go!")
}

func main() {
	http.HandleFunc("/hellogo",hellogoHandler)
	http.HandleFunc("/higo" , hiGoHandler)
	fmt.Println("Server is running on port 8080")
	
	//start the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}