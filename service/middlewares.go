package service

import (
	"github.com/justinas/alice"
	"github.com/rs/cors"
	"github.com/slapec93/bitrise-api-utils/middleware"
	"github.com/slapec93/bitrise-attendance/configs"
)

// MiddlewareProvider ...
type MiddlewareProvider struct {
	Config configs.Model
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
	)
}
