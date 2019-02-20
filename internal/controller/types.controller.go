package controller

import (
	"github.com/tax/internal/entity"
	"github.com/tax/internal/repository"
)

// Service module
type (
	systemSvc struct {
		systemRP repository.System
	}
	billsSvc struct {
		billsRP repository.Bills
	}
)

type (
	// A System service provides all functions related for System
	System interface {
		LogOpenFile()
	}
	Bills interface {
		GetBillsList() ([]entity.FormattedBills, entity.BillsInfo)
		CreateBills(entity.Bills) error
	}
)
