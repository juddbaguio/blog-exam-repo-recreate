package main

import (
	"log"
	"os"

	"github.com/juddbaguio/blog-exam-repo-recreate/pkg/api"
	"github.com/juddbaguio/blog-exam-repo-recreate/pkg/infra"
)

func main() {
	db, err := infra.InitDB()

	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	srv, err := api.InitServer(db)

	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	if err := srv.Start(); err != nil {
		log.Printf("server error: %v\n", err.Error())
		os.Exit(1)
	}
}
