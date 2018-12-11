package service

import (
	"net/http"
	"os"

	"github.com/pkg/errors"
	"github.com/slapec93/bitrise-api-utils/httpresponse"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

// AuthorizationURL ...
func AuthorizationURL(w http.ResponseWriter, r *http.Request) error {
	b := []byte(os.Getenv("CREDENTIALS"))
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		return errors.WithStack(err)
	}

	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return httpresponse.RespondWithSuccess(w, map[string]string{"message": authURL})
}
