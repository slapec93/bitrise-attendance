package service

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
	"github.com/slapec93/bitrise-api-utils/httpresponse"
	"github.com/slapec93/bitrise-attendance/utils"
)

// OpenNewMonthParams ...
type OpenNewMonthParams struct {
	SpreadsheetID string `json:"spreadsheet_id"`
	SheetName     string `json:"sheet_name"`
	Month         string `json:"month"`
}

// OpenNewMonth ...
func OpenNewMonth(w http.ResponseWriter, r *http.Request) error {
	client, err := GetSheetsClientFromContext(r.Context())
	if err != nil {
		return errors.WithStack(err)
	}

	var params OpenNewMonthParams
	defer utils.RequestBodyCloseWithErrorLog(r)
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		return httpresponse.RespondWithBadRequestError(w, "Invalid request body, JSON decode failed")
	}

	names, err := client.GetNameListFromSheet(params.SpreadsheetID)
	if err != nil {
		return errors.WithStack(err)
	}
	spreadSheet, err := client.GetSpreadSheet(params.SpreadsheetID)
	if err != nil {
		return errors.WithStack(err)
	}

	sheetID := (int64)(-1)
	for _, sheet := range spreadSheet.Sheets {
		if sheet.Properties.Title == params.SheetName {
			sheetID = sheet.Properties.SheetId
			break
		}
	}
	if sheetID < 0 {
		return httpresponse.RespondWithBadRequestError(w, "No sheet name was provided")
	}

	for _, name := range names {
		sheetProps, err := client.DuplicateSheet(params.SpreadsheetID, sheetID)
		if err != nil {
			return errors.WithStack(err)
		}

		err = client.RenameSheet(params.SpreadsheetID, sheetProps.SheetId, name)
		if err != nil {
			return errors.WithStack(err)
		}

		err = client.WriteNameAndDateToSpreadSheet(params.SpreadsheetID, name, map[string]string{
			"C9":  name,
			"C10": params.Month,
		})
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return httpresponse.RespondWithSuccess(w, names)
}
