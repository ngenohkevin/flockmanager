package main

import (
	"github.com/ngenohkevin/flock_manager/app"
	"github.com/ngenohkevin/flock_manager/database"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {

	config := database.GetConfig()

	application := &app.App{}
	application.Initialize(config)

	log.Println("Starting server")


	srv := &http.Server{
		Addr:         os.Getenv("PORT"),
		Handler:      application.SetRouters(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())



}