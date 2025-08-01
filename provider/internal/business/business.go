package business

import (
	"flight-api-provider/internal/repositories"
)

type Business struct {
}

func NewBusiness(repo *repositories.Repository) Business {
	return Business{}
}
