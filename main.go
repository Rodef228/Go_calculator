package main

import (
	"calc_service/calc_service"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type Request struct {
	Expression string `json:"expression"`
}

type Response struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func main() {
	http.HandleFunc("/api/v1/calculate", handler)
	fmt.Println("Server running on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Expression == "" {
		w.WriteHeader(http.StatusUnprocessableEntity)
		// fmt.Println("penis")
		json.NewEncoder(w).Encode(Response{Error: "Expression is not valid"})
		return
	}

	// fmt.Println("penis")
	result, err := calc_service.Calc(req.Expression)
	if err != nil {
		if errors.Is(err, errors.New("invalid expression format")) {
			w.WriteHeader(http.StatusUnprocessableEntity)
			// fmt.Println("penis")
			json.NewEncoder(w).Encode(Response{Error: "Expression is not valid"})
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			// fmt.Println("penis")
			json.NewEncoder(w).Encode(Response{Error: "Internal server error"})
		}
		return
	}

	json.NewEncoder(w).Encode(Response{Result: fmt.Sprintf("%v", result)})
	// fmt.Println("penis")
}

// на сточках 40 45 50 54, закоментиованы отладочные пенисы )