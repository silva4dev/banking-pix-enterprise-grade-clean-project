package factory

import (
	"github.com/jinzhu/gorm"
	"github.com/silva4dev/banking-pix-enterprise-grade-clean-project/application/usecase"
	"github.com/silva4dev/banking-pix-enterprise-grade-clean-project/infrastructure/repository"
)

func TransactionUseCaseFactory(database *gorm.DB) usecase.TransactionUseCase {
	pixRepository := repository.PixKeyRepositoryDb{Db: database}
	transactionRepository := repository.TransactionRepositoryDb{Db: database}

	transactionUseCase := usecase.TransactionUseCase{
		TransactionRepository: &transactionRepository,
		PixRepository:         pixRepository,
	}

	return transactionUseCase
}
