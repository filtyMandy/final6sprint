package main

import (
	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
	"log"
	"os"
)

func main() {

	logger := log.New(os.Stdout, "[morse]", log.LstdFlags)
	srv := server.New(logger)
	if err := srv.Start(); err != nil {
		logger.Fatal(err)
	}
}
