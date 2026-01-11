package provider

import "context"

type GenerateRequest struct {
	System string
	User   string
}

type Provider interface {
	Name() string
	Generate(ctx context.Context, req GenerateRequest) (string, error)
}
