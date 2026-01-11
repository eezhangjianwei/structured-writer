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
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	sceneCfg, err := prompt.GetScene(req.SceneCode)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	p := prompt.Prompt{
		System: prompt.BaseSystem,
		Scene:  sceneCfg.Scene,
		Format: sceneCfg.Format,
	}

	userText, err := p.Render(req.Params)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := s.Generator.Generate(
		r.Context(),
		provider.GenerateRequest{
			System: prompt.BaseSystem,
			User:   userText,
		},
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"result": result,
	})
}
