package internal

import (
	"github.com/tax/internal/controller"
	"github.com/tax/internal/repository"
	"github.com/tax/internal/usecase"
)

// GetUsecase for project
func GetUsecase(config *Config, ipAddress string) *Usecase {

	// DATABASE
	db := initDatabase(config.Database)

	// REDIS
	// rds := initRedis(config.Redis)

	// NSQ PRODUCER
	// nsqProducer := initNSQProducer(config.NSQd.Endpoint)

	// ELASTICSEARCH
	// elasticSearch := initElastic(config.Elastic)

	// REPOSITORY
	systemRepo := repository.NewSystem(ipAddress)
	billsRepo := repository.NewBills(db["db_tax"])

	// CONTROLLER
	systemService := controller.NewSystem(systemRepo)
	billsService := controller.NewBills(billsRepo)

	// USECASE

	systemUsecase := usecase.NewSystem(systemService)
	billsUseCase := usecase.NewBills(billsService)

	ucase := &Usecase{
		System: systemUsecase,
		Bills:  billsUseCase,
	}

	return ucase
}
