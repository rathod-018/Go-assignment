package main

import (
	"encoding/json"
	"fmt"
	"goAssignment/processor"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Go server")
	})

	http.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {

		result := processor.ProcessData()

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})


	fmt.Println("Server running on port 8080")
	http.ListenAndServe(":8080", nil)
}
