package genkit

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

const geminiEndpoint = "https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash:generateContent"

type GeminiRequest struct {
	Contents []struct {
		Parts []struct {
			Text string `json:"text"`
		} `json:"parts"`
	} `json:"contents"`
}

type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

func Generate(prompt string) (string, error) {
	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return "", errors.New("GEMINI_API_KEY not set")
	}

	payload := map[string]any{
		"contents": []any{
			map[string]any{
				"parts": []any{
					map[string]string{"text": prompt},
				},
			},
		},
	}

	b, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", geminiEndpoint+"?key="+apiKey, bytes.NewBuffer(b))
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var res GeminiResponse
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return "", err
	}

	if len(res.Candidates) == 0 {
		return "", errors.New("empty response")
	}

	return res.Candidates[0].Content.Parts[0].Text, nil
}
