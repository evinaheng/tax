package usecase

import (
	"github.com/tax/internal/controller"
	"github.com/tax/internal/entity"
)

// NewBills
func NewBills(billsSvc controller.Bills) Bills {
	ucase := &billsUsecase{
		bills: billsSvc,
	}
	return ucase
}

func (s *billsUsecase) GetBillsList() ([]entity.FormattedBills, entity.BillsInfo) {
	return s.bills.GetBillsList()
}

func (s *billsUsecase) CreateBills(name string, types int, price int) error {
	newObj := entity.Bills{
		Name:    name,
		TaxCode: types,
		Price:   float64(price),
	}
	return s.bills.CreateBills(newObj)
}
