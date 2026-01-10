package types

type GenerateRequest struct {
	SceneCode string `json:"scene_code"`
	Topic     string `json:"topic"`
	Audience  string `json:"audience"`
}
