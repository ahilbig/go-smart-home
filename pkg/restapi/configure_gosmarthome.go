// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"go-smart-home/pkg/raspi"
	"net/http"
	"regexp"

	"go-smart-home/pkg/restapi/operations"
	"go-smart-home/pkg/restapi/operations/switch_operations"
)

//go:generate swagger generate server --target ..\..\..\go-smart-home --name Gosmarthome --spec ..\..\swagger.yml --model-package pkg\models --server-package pkg\restapi

func configureFlags(api *operations.GosmarthomeAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.GosmarthomeAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.XMLConsumer = runtime.XMLConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.XMLProducer = runtime.XMLProducer()

	api.TxtProducer = runtime.TextProducer()

	api.SwitchOperationsAddSwitchHandler = switch_operations.AddSwitchHandlerFunc(func(params switch_operations.AddSwitchParams) middleware.Responder {
		s, err := raspi.GetConfInstance().CreateSwitch(params.Body)

		if err != nil {
			return switch_operations.NewAddSwitchInternalServerError().WithPayload(err.Error())
		}
		return switch_operations.NewAddSwitchOK().WithPayload(s)
	})
	api.SwitchOperationsListSwitchesHandler = switch_operations.ListSwitchesHandlerFunc(func(params switch_operations.ListSwitchesParams) middleware.Responder {
		switches := raspi.GetConfInstance().GetSwitches()

		return switch_operations.NewListSwitchesOK().WithPayload(switches)
	})

	// Not implemented

	api.SwitchOperationsGetSwitchByIDHandler = switch_operations.GetSwitchByIDHandlerFunc(func(params switch_operations.GetSwitchByIDParams) middleware.Responder {
		return middleware.NotImplemented("operation switch_operations.GetSwitchByID has not yet been implemented")
	})
	api.SwitchOperationsGetSwitchStateHandler = switch_operations.GetSwitchStateHandlerFunc(func(params switch_operations.GetSwitchStateParams) middleware.Responder {
		return middleware.NotImplemented("operation switch_operations.GetSwitchState has not yet been implemented")
	})

	api.SwitchOperationsSetSwitchStateHandler = switch_operations.SetSwitchStateHandlerFunc(func(params switch_operations.SetSwitchStateParams) middleware.Responder {
		return middleware.NotImplemented("operation switch_operations.SetSwitchState has not yet been implemented")
	})
	api.SwitchOperationsTapSwitchHandler = switch_operations.TapSwitchHandlerFunc(func(params switch_operations.TapSwitchParams) middleware.Responder {
		return middleware.NotImplemented("operation switch_operations.TapSwitch has not yet been implemented")
	})
	api.SwitchOperationsToggleSwitchStateHandler = switch_operations.ToggleSwitchStateHandlerFunc(func(params switch_operations.ToggleSwitchStateParams) middleware.Responder {
		return middleware.NotImplemented("operation switch_operations.ToggleSwitchState has not yet been implemented")
	})
	api.SwitchOperationsUpdateSwitchHandler = switch_operations.UpdateSwitchHandlerFunc(func(params switch_operations.UpdateSwitchParams) middleware.Responder {
		return middleware.NotImplemented("operation switch_operations.UpdateSwitch has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if m, _ := regexp.MatchString("/api.*|/docs.*|/swagger.*", r.URL.Path); m {
			handler.ServeHTTP(w, r)
		} else {
			http.FileServer(http.Dir("./views")).ServeHTTP(w, r)
		}
	})
}
