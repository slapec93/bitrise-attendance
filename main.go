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
