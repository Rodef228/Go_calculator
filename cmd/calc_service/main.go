package main

import (
	"http_calculator/internal/handlers"
	"fmt"
	"net/http"
)



func main() {
	http.HandleFunc("/api/v1/calculate", handlers.Handler)
	fmt.Println("Server running on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

