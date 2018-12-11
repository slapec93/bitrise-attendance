package service

import (
	"net/http"

	"github.com/justinas/alice"
	"github.com/rs/cors"
	"github.com/slapec93/bitrise-api-utils/middleware"
	"github.com/slapec93/bitrise-attendance/configs"
	"github.com/slapec93/bitrise-attendance/sheets"
)

// MiddlewareProvider ...
type MiddlewareProvider struct {
	Config       configs.Model
	SheetsClient sheets.Client
}

func createSetConfigMiddleware(config configs.Model) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := ContextWithConfig(r.Context(), config)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func createSetSheetsClientMiddleware(sheetsClient sheets.Client) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := ContextWithSheetsClient(r.Context(), sheetsClient)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

// CommonMiddleware ...
func (m MiddlewareProvider) CommonMiddleware() alice.Chain {
	commonMiddleware := alice.New(
		cors.AllowAll().Handler,
	)

	if m.Config.EnvMode == configs.ServerEnvModeProd {
		commonMiddleware = commonMiddleware.Append(
			middleware.CreateRedirectToHTTPSMiddleware(),
		)
	}

	return commonMiddleware.Append(
		middleware.CreateOptionsRequestTerminatorMiddleware(),
		createSetConfigMiddleware(m.Config),
		createSetSheetsClientMiddleware(m.SheetsClient),
	)
}
