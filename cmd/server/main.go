package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/joho/godotenv"

	"structured-content-genkit-go/internal/flows"
	"structured-content-genkit-go/internal/types"
)

func main() {
	_ = godotenv.Load()

	http.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var req types.GenerateRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		result, err := flows.GenerateStructuredContent(req)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}

		_ = json.NewEncoder(w).Encode(map[string]string{
			"result": result,
			})
	})

	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
