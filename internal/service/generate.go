package service

import (
	"context"
	"errors"

	"ai-product/internal/provider"
)

type GenerateService struct {
	Providers []provider.Provider
}

func (s *GenerateService) Generate(ctx context.Context, req provider.GenerateRequest) (string, error) {
	var lastErr error
	for _, p := range s.Providers {
		result, err := p.Generate(ctx, req)
		if err == nil {
			return result, nil
		}
		if errors.Is(err, provider.ErrRateLimited) || errors.Is(err, provider.ErrEmptyResponse) {
			lastErr = err
			continue
		}
		return "", err
	}

	if lastErr != nil {
		return "", lastErr
	}

	return "", errors.New("all providers failed")
}
