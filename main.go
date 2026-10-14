package main

import (
	"log"

	"github.com/nikndip/go_final_project/pkg/server"
)

func main() {
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
