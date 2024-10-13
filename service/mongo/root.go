package mongo

import (
	"eCommerce/repository"
)

type MService struct {
	repository *repository.Repository
}

func NewMService(repository *repository.Repository) *MService {
	r := &MService{
		repository: repository,
	}

	return r
}
