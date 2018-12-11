package sheets

import (
	"fmt"

	"github.com/pkg/errors"
	gsheets "google.golang.org/api/sheets/v4"
)

// Client ...
type Client struct {
	service *gsheets.Service
}

// NewClient ...
func NewClient(service *gsheets.Service) Client {
	return Client{
		service: service,
	}
}

// CreateSpreadsheet ...
func (c Client) CreateSpreadsheet(title string) (*gsheets.Spreadsheet, error) {
	spreadSheet, err := c.service.Spreadsheets.Create(&gsheets.Spreadsheet{
		Properties: &gsheets.SpreadsheetProperties{
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
	updateRequest := gsheets.BatchUpdateSpreadsheetRequest{
		Requests: []*gsheets.Request{
			&gsheets.Request{
				UpdateSheetProperties: &gsheets.UpdateSheetPropertiesRequest{
					Fields: "title",
					Properties: &gsheets.SheetProperties{
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
func (c Client) DuplicateSheet(spreadSheetID string, sheetID int64) (*gsheets.SheetProperties, error) {
	copyRequest := gsheets.CopySheetToAnotherSpreadsheetRequest{
		DestinationSpreadsheetId: spreadSheetID,
	}
	properties, err := c.service.Spreadsheets.Sheets.CopyTo(spreadSheetID, sheetID, &copyRequest).Do()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return properties, nil
}

// GetSpreadSheet ...
func (c Client) GetSpreadSheet(spreadSheetID string) (*gsheets.Spreadsheet, error) {
	return c.service.Spreadsheets.Get(spreadSheetID).Do()
}

// WriteNameToSpreadSheet ...
func (c Client) WriteNameToSpreadSheet(spreadSheetID, name, sheetName, cell string) error {
	value := gsheets.ValueRange{
		Range:  fmt.Sprintf("%s!%s:%s", sheetName, cell, cell),
		Values: [][]interface{}{[]interface{}{name}},
	}
	_, err := c.service.Spreadsheets.Values.Update(
		spreadSheetID,
		fmt.Sprintf("%s!%s:%s", sheetName, cell, cell),
		&value,
	).ValueInputOption("RAW").Do()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// GetNameListFromSheet ...
func (c Client) GetNameListFromSheet(spreadSheetID string) ([]string, error) {
	values, err := c.service.Spreadsheets.Values.Get(spreadSheetID, "data!names").Do()
	if err != nil {
		return []string{}, errors.WithStack(err)
	}
	names := []string{}
	for _, name := range values.Values {
		names = append(names, fmt.Sprintf("%s", name[0]))
	}
	return names, nil
}
