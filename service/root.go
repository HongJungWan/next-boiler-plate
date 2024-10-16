package service

import (
	"eCommerce/config"
	"eCommerce/repository"
	"eCommerce/service/mongo"
	"eCommerce/service/mysql"
)

type Service struct {
	config     *config.Config
	repository *repository.Repository

	MService     *mongo.MService
	MySQLService *mysql.MySQLService
}

func NewService(config *config.Config, repository *repository.Repository) (*Service, error) {
	r := &Service{
		config:       config,
		repository:   repository,
		MService:     mongo.NewMService(repository),
		MySQLService: mysql.NewMySQLService(repository),
	}

	return r, nil
}
