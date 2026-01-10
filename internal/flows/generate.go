package flows

import (
	"errors"
	"fmt"

	"structured-content-genkit-go/internal/genkit"
	"structured-content-genkit-go/internal/prompts"
	"structured-content-genkit-go/internal/types"
)

func GenerateStructuredContent(req types.GenerateRequest) (string, error) {
	scene, ok := prompts.Scenes[req.SceneCode]
	if !ok {
		return "", errors.New("invalid scene_code")
	}

	prompt := fmt.Sprintf(`
%s
%s
%s
主题：%s
目标受众：%s
`, prompts.SystemPrompt, scene, prompts.OutputFormat, req.Topic, req.Audience)

	return genkit.Generate(prompt)
}
