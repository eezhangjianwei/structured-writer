package provider

import "errors"

var (
	ErrRateLimited   = errors.New("rate_limited")
	ErrEmptyResponse = errors.New("empty response")
)
