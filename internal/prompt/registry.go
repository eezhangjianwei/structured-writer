package prompt

import (
	"errors"

	"ai-product/internal/prompt/formats"
	"ai-product/internal/prompt/scenes"
)

type SceneConfig struct {
	Scene  string
	Format string
}

var SceneRegistry = map[string]SceneConfig{
	"CONTENT_POST_GENERIC": {
		Scene:  scenes.ContentPost,
		Format: formats.Generic,
	},
	"CONTENT_POST_XHS": {
		Scene:  scenes.ContentPost,
		Format: formats.XHS,
	},
	"CONTENT_POST_ZHIHU": {
		Scene:  scenes.ContentPost,
		Format: formats.Zhihu,
	},
	"PRODUCT_DESC": {
		Scene:  scenes.ProductDescription,
		Format: formats.Generic,
	},
}

func GetScene(code string) (SceneConfig, error) {
	if code == "" {
		code = "CONTENT_POST_GENERIC"
	}
	cfg, ok := SceneRegistry[code]
	if !ok {
		return SceneConfig{}, errors.New("unknown scene_code")
	}
	return cfg, nil
}
