package handlers

import (
	"flight-api/internal/business"
)

type Handlers struct {
}

func NewHandler(business business.Business) Handlers {
	return Handlers{}
}
