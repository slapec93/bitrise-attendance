package main

import (
	"log"
	"net/http"

	"github.com/pkg/errors"
	"github.com/slapec93/bitrise-attendance/configs"
	"github.com/slapec93/bitrise-attendance/router"
)

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
