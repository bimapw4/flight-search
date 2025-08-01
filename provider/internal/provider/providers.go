package provider

import (
	"flight-api-provider/bootstrap"
)

type Provider struct {
}

func NewProvider(cfg bootstrap.Providers) Provider {
	return Provider{}
}
