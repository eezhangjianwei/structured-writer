package provider

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GeminiProvider struct {
	APIKey string
	Model  string
}

func (g *GeminiProvider) Name() string {
	return "gemini"
}

func (g *GeminiProvider) Generate(ctx context.Context, req GenerateRequest) (string, error) {
	payload := map[string]any{
		"contents": []any{
			map[string]any{
				"parts": []any{
					map[string]string{"text": req.System},
					map[string]string{"text": req.User},
				},
			},
		},
	}

	body, _ := json.Marshal(payload)
	url := fmt.Sprintf("https://generativelanguage.googleapis.com/v1beta/models/%s:generateContent", g.Model)

	httpReq, _ := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("X-goog-api-key", g.APIKey)

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	raw, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == 429 {
			return "", ErrRateLimited
		}
		return "", fmt.Errorf("gemini error: status=%d body=%s", resp.StatusCode, string(raw))
	}

	var parsed struct {
		Candidates []struct {
			Content struct {
				Parts []struct {
					Text string `json:"text"`
				} `json:"parts"`
			} `json:"content"`
		} `json:"candidates"`
	}

	if err := json.Unmarshal(raw, &parsed); err != nil {
		return "", err
	}

	if len(parsed.Candidates) == 0 || len(parsed.Candidates[0].Content.Parts) == 0 {
		return "", ErrEmptyResponse
	}

	return parsed.Candidates[0].Content.Parts[0].Text, nil
}
