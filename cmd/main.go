package main

import (
	"log"

	atest "github.com/Elhemist/avito-test"
	"github.com/Elhemist/avito-test/internal/handler"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(atest.Server)
	if err := srv.Start("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Running error: %s", err.Error())
	}
}
