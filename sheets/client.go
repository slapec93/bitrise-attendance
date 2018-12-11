package sheets

import (
	"github.com/pkg/errors"
	sheets "google.golang.org/api/sheets/v4"
)

// Client ...
type Client struct {
	service *sheets.Service
}

// NewClient ...
func NewClient(service *sheets.Service) Client {
	return Client{
		service: service,
	}
}

// CreateSpreadsheet ...
func (c Client) CreateSpreadsheet(title string) (*sheets.Spreadsheet, error) {
	spreadSheet, err := c.service.Spreadsheets.Create(&sheets.Spreadsheet{
		Properties: &sheets.SpreadsheetProperties{
			Title: title,
		},
	}).Do()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return spreadSheet, nil
}

// RenameSheet ...
func (c Client) RenameSheet(spreadSheetID string, sheetID int64, newTitle string) error {
	updateRequest := sheets.BatchUpdateSpreadsheetRequest{
		Requests: []*sheets.Request{
			&sheets.Request{
				UpdateSheetProperties: &sheets.UpdateSheetPropertiesRequest{
					Fields: "title",
					Properties: &sheets.SheetProperties{
						SheetId: sheetID,
						Title:   newTitle,
					},
				},
			},
		},
	}
	_, err := c.service.Spreadsheets.BatchUpdate(spreadSheetID, &updateRequest).Do()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// DuplicateSheet ...
func (c Client) DuplicateSheet(spreadSheetID string, sheetID int64) error {
	copyRequest := sheets.CopySheetToAnotherSpreadsheetRequest{
		DestinationSpreadsheetId: spreadSheetID,
	}
	_, err := c.service.Spreadsheets.Sheets.CopyTo(spreadSheetID, sheetID, &copyRequest).Do()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// GetSpreadSheet ...
func (c Client) GetSpreadSheet(spreadSheetID string) (*sheets.Spreadsheet, error) {
	return c.service.Spreadsheets.Get(spreadSheetID).Do()
}
