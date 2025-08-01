package handlers

import (
	"flight-api-provider/internal/business"
)

type Handlers struct {
}

func NewHandler(business business.Business) Handlers {
	return Handlers{}
}
