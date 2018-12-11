package router

import (
	"github.com/gorilla/mux"
	"github.com/slapec93/bitrise-attendance/configs"
	"github.com/slapec93/bitrise-attendance/service"

	"github.com/slapec93/bitrise-api-utils/httpresponse"
)

// New ...
func New(config configs.Model) *mux.Router {
	// StrictSlash: allow "trim slash"; /x/ REDIRECTS to /x
	r := mux.NewRouter().StrictSlash(true)

	middlewareProvider := service.MiddlewareProvider{
		Config: config,
	}

	r.Handle("/auth", middlewareProvider.CommonMiddleware().Then(
		httpresponse.InternalErrHandlerFuncAdapter(service.AuthorizationURL))).Methods("GET", "OPTIONS")

	r.Handle("/open-new-month", middlewareProvider.CommonMiddleware().Then(
		httpresponse.InternalErrHandlerFuncAdapter(service.OpenNewMonth))).Methods("POST", "OPTIONS")
	return r
}
