package router

import (
	"github.com/bitrise-io/go-utils/log"
	"github.com/gorilla/mux"
	"github.com/slapec93/bitrise-attendance/configs"
	"github.com/slapec93/bitrise-attendance/service"
	"github.com/slapec93/bitrise-attendance/sheets"

	"github.com/slapec93/bitrise-api-utils/httpresponse"
	gsheets "google.golang.org/api/sheets/v4"
)

// New ...
func New(config configs.Model) *mux.Router {
	// StrictSlash: allow "trim slash"; /x/ REDIRECTS to /x
	r := mux.NewRouter().StrictSlash(true)

	srv, err := gsheets.New(config.Client)
	if err != nil {
		log.Errorf("Failed to create Google Sheets Service object")
	}

	client := sheets.NewClient(srv)
	middlewareProvider := service.MiddlewareProvider{
		Config:       config,
		SheetsClient: client,
	}

	r.Handle("/open-new-month", middlewareProvider.CommonMiddleware().Then(
		httpresponse.InternalErrHandlerFuncAdapter(service.OpenNewMonth))).Methods("POST", "OPTIONS")

	return r
}
