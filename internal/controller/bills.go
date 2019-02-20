package controller

import (
	"errors"

	"github.com/tax/internal/constant"
	"github.com/tax/internal/entity"
	"github.com/tax/internal/repository"
)

// NewBills New Bills service
func NewBills(billsRP repository.Bills) Bills {
	svc := &billsSvc{
		billsRP: billsRP,
	}
	return svc
}

func (s *billsSvc) GetBillsList() ([]entity.FormattedBills, entity.BillsInfo) {
	res := []entity.FormattedBills{}
	grandTotal := entity.BillsInfo{}

	billsListDb := s.billsRP.GetBills()

	for b := range billsListDb {
		// default value if data is not exists
		typeFoodItem := "Unknown"
		isRefundAble := false
		defaultPrice := billsListDb[b].Price

		taxCode := billsListDb[b].TaxCode
		if _, ok := constant.TaxMap[taxCode]; ok {
			// exists in map
			typeFoodItem = constant.TaxMap[taxCode]
			isRefundAble = constant.ItemRefund[taxCode]
		}

		newBills := entity.FormattedBills{
			Name:       billsListDb[b].Name,
			Type:       typeFoodItem,
			Price:      billsListDb[b].Price,
			TaxCode:    billsListDb[b].TaxCode,
			Refundable: isRefundAble,
		}
		// do calculation for tax bills here

		if taxCode == constant.FOOD {
			percenTage := (float64(10) / float64(100))
			taxAmount := defaultPrice * float64(percenTage)
			newBills.Tax = taxAmount
			newBills.Amount = taxAmount + defaultPrice
		} else if taxCode == constant.TOBACCO {
			percenTage := float64(2) / float64(100)
			taxAmount := defaultPrice * percenTage
			newBills.Tax = taxAmount + 10
			newBills.Amount = taxAmount + defaultPrice + 10

		} else if taxCode == constant.ENTERTAINMENT {
			// tax free for under 100
			if (defaultPrice > 0 && defaultPrice < 100) || defaultPrice == 0 {
				newBills.Tax = 0
				newBills.Amount = defaultPrice
			} else if defaultPrice >= 100 {
				percenTage := (float64(1) / float64(100))
				taxAmount := (defaultPrice - 100) * percenTage
				newBills.Tax = taxAmount
				newBills.Amount = taxAmount + defaultPrice
			}
		} else {
			newBills.Tax = 0
			newBills.Amount = defaultPrice
			// do nothing here put normal price and zero tax
		}

		res = append(res, newBills)

		grandTotal.GrandTotal += newBills.Amount
		grandTotal.SubTotal += newBills.Price
		grandTotal.TaxSubTotal += newBills.Tax
	}

	return res, grandTotal
}

func (s *billsSvc) CreateBills(param entity.Bills) error {

	if param.TaxCode != constant.FOOD && param.TaxCode != constant.ENTERTAINMENT && param.TaxCode != constant.TOBACCO {
		return errors.New("category item is not exists")
	}
	return s.billsRP.CreateBills(param)
}
