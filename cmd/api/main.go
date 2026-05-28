package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/francisjdev/community-blog/internal/app"
	"github.com/francisjdev/community-blog/internal/config"
	"github.com/francisjdev/community-blog/internal/database"
	"github.com/francisjdev/community-blog/internal/service"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	cfg := config.LoadCfg()
	db, err := sql.Open("pgx", cfg.DBUrl)
	if err != nil {
		log.Fatalf("error %v\n", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("error %v\n", err)
	}
	defer db.Close()
	dbQueries := database.New(db)
	userServ := service.NewUserService(dbQueries)
	postServ := service.NewPostService(dbQueries)
	nService := service.Services{
		Users: userServ,
		Posts: postServ,
	}
	app := app.Application{
		Config:  cfg,
		Service: &nService,
	}
	router := LoadRouter(&app)
	log.Printf("server started on %s", app.Config.Addr)
	log.Fatal(http.ListenAndServe(cfg.Addr, router))

}
