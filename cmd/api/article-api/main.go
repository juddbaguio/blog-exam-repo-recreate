package main

import (
	"log"
	"os"

	"github.com/juddbaguio/blog-exam-repo-recreate/pkg/api"
	"github.com/juddbaguio/blog-exam-repo-recreate/pkg/config"
	"github.com/juddbaguio/blog-exam-repo-recreate/pkg/logic"
	"github.com/juddbaguio/blog-exam-repo-recreate/pkg/store/mysql"
)

func main() {
	cfg, err := config.ReadConfig(".env")
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	db, err := mysql.NewStore(&cfg.DB)

	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	_logic := logic.NewLogicHandler(db)

	srv, err := api.InitServer(_logic)

	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	if err := srv.Start(); err != nil {
		log.Printf("server error: %v\n", err.Error())
		os.Exit(1)
	}
}
