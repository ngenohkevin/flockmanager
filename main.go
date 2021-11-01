package main

import (
	"fmt"
	"github.com/joho/godotenv"
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

	err := godotenv.Load()
	if err != nil {
		return 
	}
	port := os.Getenv("PORT")

	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%s","0.0.0.0",port),
		Handler:      application.SetRouters(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())



}