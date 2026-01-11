package main

import (
	"log"
	"net/http"
	"os"

	"ai-product/internal/provider"
	"ai-product/internal/service"
	"ai-product/internal/transport"
)

func main() {
	openai := &provider.OpenAIProvider{
		APIKey: os.Getenv("OPENAI_API_KEY"),
		Model:  "gpt-4o-mini",
	}

	gemini := &provider.GeminiProvider{
		APIKey: os.Getenv("GEMINI_API_KEY"),
		Model:  "gemini-2.0-flash",
	}

	genService := &service.GenerateService{
		Providers: []provider.Provider{
			openai,
			gemini,
		},
	}

	server := &transport.Server{
		Generator: genService,
	}

	http.HandleFunc("/generate", server.HandleGenerate)

	log.Println("Server running at :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
