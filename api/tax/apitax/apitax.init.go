package apitax

import (
	"context"

	"github.com/tax/internal/usecase"
	"github.com/tax/lib/router"
)

var instance *Module

// Module for APITax
type Module struct {
	// put all usecase used
	billsUseCase usecase.Bills
}

// New module
func New(billsUC usecase.Bills) *Module {
	return &Module{
		billsUseCase: billsUC,
	}
}

// Define ALL Related Module for Api Test here

// API Module
func (m *Module) GetBills(version int) func(ctx context.Context, param *router.HandlerParam) (apiresult router.HandlerResult) {
	return m.getBills
}

func (m *Module) CreateBills(version int) func(ctx context.Context, param *router.HandlerParam) (apiresult router.HandlerResult) {
	return m.createBills
}
