package mysql

import "eCommerce/repository"

type MySQLService struct {
	repository *repository.Repository
}

func NewMySQLService(repository *repository.Repository) *MySQLService {
	r := &MySQLService{
		repository: repository,
	}

	return r
}
