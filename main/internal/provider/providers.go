package provider

import (
	"flight-api/bootstrap"
)

type Provider struct {
}

func NewProvider(cfg bootstrap.Providers) Provider {
	return Provider{}
}
