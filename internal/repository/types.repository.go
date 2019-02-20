package repository

import (
	"github.com/tax/internal/entity"
	"github.com/tax/lib/storage/database"
)

// Repository Module
type (
	systemRepo struct {
		ipAddress string
	}
	billsRepo struct {
		db database.Database
	}
)

type (
	// A System repository provides all queries related for System
	System interface {
		LogOpenFile()
	}
	Bills interface {
		CreateBills(param entity.Bills) error
		GetBills() []entity.Bills
	}
)
