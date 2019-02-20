package apitax

import (
	"context"
	"net/http"

	"github.com/tax/ctxs"
	"github.com/tax/lib/router"
)

func (m *Module) createBills(ctx context.Context, param *router.HandlerParam) (apiResult router.HandlerResult) {

	typeCode, _ := ctxs.TypeCodeKeyFromContext(ctx)
	price, _ := ctxs.PriceKeyFromContext(ctx)
	name, _ := ctxs.NameKeyFromContext(ctx)

	if typeCode == 0 {
		apiResult.SetError("Type code must be defined", "86", http.StatusBadRequest)
		return apiResult
	} else if price == 0 {
		apiResult.SetError("Price must be defined", "87", http.StatusBadRequest)
		return apiResult
	} else if name == "" {
		apiResult.SetError("Name must be inputed", "88", http.StatusBadRequest)
		return apiResult
	}

	err := m.billsUseCase.CreateBills(name, typeCode, price)
	if err == nil {
		// Process result
		response := responseCreate{
			Label: "Result",
			Value: "Success Inputed Bills : " + name,
		}
		res := router.Object{
			Attributes: response,
		}
		apiResult.JSON = &router.Data{
			Data: res,
		}

	} else {
		apiResult.SetError(err.Error(), "23", http.StatusBadRequest)
		return apiResult
	}
	return apiResult
}
