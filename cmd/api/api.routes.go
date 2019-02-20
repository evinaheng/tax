package main

import (
	"net/http"
	"time"

	"github.com/tax/api"
	apitax "github.com/tax/api/tax/apitax"
	"github.com/tax/cmd/internal"
	"github.com/tax/lib/panics"
)

// Initialize all routes for API
func initRoutes(config internal.Config, panicsWrapper panics.Panics, ucase *internal.Usecase) http.Handler {

	// Get environment
	// var isDevelopment = config.Server.Env == internal.Development

	// Middlewares
	// var mw router.Middleware

	// Init ROUTER
	apiRouter := api.NewRouter(api.RouterConfig{
		Timeout: time.Duration(config.Environment.GlobalTimeout) * time.Second,
	})

	// -- Set Panics handler TODO
	apiRouter.PanicsHandler(panicsWrapper.RouterHandler)

	// -----------------------------------------------------------
	// START - ADD ROUTE
	// WARNING! Duplicate routes will cause a panic!
	// -----------------------------------------------------------

	// -- Initialize API Module
	apiTax := apitax.New(ucase.Bills)

	// -- Routes WITHOUT authentication
	apiRouter.AddRoute(http.MethodGet, "tax/getBills", apiTax.GetBills)
	apiRouter.AddRoute(http.MethodGet, "tax/createBills", apiTax.CreateBills)

	// -- Routes WITH authentication

	// -----------------------------------------------------------
	// END - ADD ROUTE
	// -----------------------------------------------------------

	return apiRouter.GetHandler()

}
