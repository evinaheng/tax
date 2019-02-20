package usecase

import (
	"github.com/tax/internal/controller"
	"github.com/tax/internal/entity"
)

//  Usecase module
type (
	systemUsecase struct {
		// Controllers
		system controller.System
	}
	billsUsecase struct {
		bills controller.Bills
	}
)

// Usecase interfaces
type (
	// A System usecase provides functions for maintaining system information
	System interface {
		LogOpenFile()
	}
	Bills interface {
		CreateBills(name string, types int, price int) error
		GetBillsList() ([]entity.FormattedBills, entity.BillsInfo)
	}
)
