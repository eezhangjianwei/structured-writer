package genkit

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	apiURL = "https://generativelanguage.googleapis.com/v1beta/models/gemini-2.0-flash:generateContent"
)

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
					map[string]string{
						"text": prompt,
					},
				},
			},
		},
	}

	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest(
		"POST",
		apiURL,
		bytes.NewBuffer(body),
	)

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-goog-api-key", apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf(
			"gemini api error: status=%d body=%s",
			resp.StatusCode,
			string(raw),
		)
	}

	var result GeminiResponse
	if err := json.Unmarshal(raw, &result); err != nil {
		return "", err
	}

	if len(result.Candidates) == 0 ||
		len(result.Candidates[0].Content.Parts) == 0 {
		return "", errors.New("empty candidates from gemini-2.0-flash")
	}

	return result.Candidates[0].Content.Parts[0].Text, nil
}
