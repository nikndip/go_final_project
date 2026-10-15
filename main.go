package main

import (
	"log"

	"github.com/nikndip/go_final_project/pkg/db"
	"github.com/nikndip/go_final_project/pkg/server"
)

func main() {
	if err := db.Init("scheduler.db"); err != nil {
		log.Fatal(err)
	}

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
