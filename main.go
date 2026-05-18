package main

import (
	"classification-engine/internal/handlers"
	"fmt"
	"net/http"
)

func main() {

	handler := handlers.NewClassificationHandler()

	http.HandleFunc("/classify", handler.Classify)

	fmt.Println("Server running on :8080")

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	}
}
