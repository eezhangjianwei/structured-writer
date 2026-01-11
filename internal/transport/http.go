package transport

import (
	"encoding/json"
	"net/http"

	"ai-product/internal/prompt"
	"ai-product/internal/provider"
	"ai-product/internal/service"
)

type Server struct {
	Generator *service.GenerateService
}

func (s *Server) HandleGenerate(w http.ResponseWriter, r *http.Request) {
	var req struct {
		SceneCode string            `json:"scene_code"`
		Params    map[string]string `json:"params"`
	}
	_ = json.NewDecoder(r.Body).Decode(&req)

	p := prompt.ContentPostV1

	user, err := p.Render(req.Params)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	result, err := s.Generator.Generate(
		r.Context(),
		provider.GenerateRequest{
			System: p.System,
			User:   user,
		},
	)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"result": result,
	})
}
