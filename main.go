package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/pkg/errors"
	"github.com/slapec93/bitrise-attendance/configs"
	"github.com/slapec93/bitrise-attendance/router"
	"google.golang.org/api/sheets/v4"
)

func writeNameToSpreadSheet(service *sheets.Service, spreadSheetID, name, sheetName, cell string) error {
	value := sheets.ValueRange{
		Range:  fmt.Sprintf("%s!%s:%s", sheetName, cell, cell),
		Values: [][]interface{}{[]interface{}{"valamike"}},
	}
	_, err := service.Spreadsheets.Values.Update(
		spreadSheetID,
		fmt.Sprintf("%s!%s:%s", sheetName, cell, cell),
		&value,
	).ValueInputOption("RAW").Do()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// func main() {
// 	// If modifying these scopes, delete your previously saved token.json.
// 	client := configs.CreateNewConfig()
// 	// client := getClient(config)

// 	srv, err := sheets.New(client.Client)
// 	if err != nil {
// 		log.Fatalf("Unable to retrieve Sheets client: %v", err)
// 	}

// 	// Prints the names and majors of students in a sample spreadsheet:
// 	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
// 	spreadsheetID := "1J0tqdA4frZ0khRUB05fPRC0vBsi_W-8kIgf6iYkoK0c"
// 	// copyRange := "Sheet2!A1:F44"

// 	values, err := srv.Spreadsheets.Values.Get(spreadsheetID, "data!names").Do()
// 	for _, name := range values.Values {
// 		fmt.Printf("%#v", name[0])
// 	}
// }

func mainImplementation() error {
	conf := configs.CreateNewConfig()

	// Routing
	http.Handle("/", router.New(conf))
	log.Println("Starting - using port:", conf.Port)
	if err := http.ListenAndServe(":"+conf.Port, nil); err != nil {
		return errors.Wrap(errors.WithStack(err), "Failed to ListenAndServe")
	}
	return nil
}

func main() {
	if err := mainImplementation(); err != nil {
		log.Fatalf(" [!] Exception: Failed to initialize Bitrise Attendance: %+v", err)
	}
}
