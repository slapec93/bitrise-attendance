package service

import (
	"context"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/slapec93/bitrise-api-utils/httpresponse"
	"github.com/slapec93/bitrise-attendance/sheets"
	"golang.org/x/oauth2/google"
	gsheets "google.golang.org/api/sheets/v4"
)

func userAuthTokenFromHeader(h http.Header) (string, error) {
	authHeader := h.Get("Authorization")
	token := strings.TrimPrefix(authHeader, "token ")
	if token == "" {
		return "", errors.New("No Authorization header specified")
	}
	return token, nil
}

// MiddlewareUserAuthByTokenInHeaderHandler ...
func MiddlewareUserAuthByTokenInHeaderHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken, err := userAuthTokenFromHeader(r.Header)
		if err != nil {
			// no User Auth header specified, continue without auth (without adding to context)
			h.ServeHTTP(w, r)
			return
		}

		b := []byte(os.Getenv("CREDENTIALS"))
		config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
		if err != nil {
			httpresponse.RespondWithInternalServerError(w, err)
			log.Fatalf("Unable to parse client secret file to config: %v", err)
		}

		tok, err := config.Exchange(context.Background(), authToken)
		if err != nil {
			httpresponse.RespondWithInternalServerError(w, err)
			log.Fatalf("Unable to retrieve token from web: %v", err)
		}

		client := config.Client(context.Background(), tok)
		srv, err := gsheets.New(client)
		if err != nil {
			httpresponse.RespondWithInternalServerError(w, err)
			log.Fatalf("Failed to create Google Sheets Service object")
		}
		sheetsClient := sheets.NewClient(srv)

		ctx := ContextWithSheetsClient(r.Context(), sheetsClient)
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}
