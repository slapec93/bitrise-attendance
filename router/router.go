package router

import (
	"github.com/bitrise-team/bitrise-api/service/v0"
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

	r.Handle("/", middlewareProvider.CommonMiddleware().Then(
		httpresponse.InternalErrHandlerFuncAdapter(v0.BranchListHandler))).Methods("GET", "OPTIONS")

	return r
}
