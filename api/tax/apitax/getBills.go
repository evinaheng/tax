package apitax

import (
	"context"

	"github.com/tax/lib/router"
)

func (m *Module) getBills(ctx context.Context, param *router.HandlerParam) (apiResult router.HandlerResult) {

	res, grandTotal := m.billsUseCase.GetBillsList()

	response := router.Object{
		Attributes: res,
	}
	apiResult.JSON = &router.Data{
		Data: response,
		Info: grandTotal,
	}

	return apiResult
}
