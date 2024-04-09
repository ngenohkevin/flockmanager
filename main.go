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

	fmt.Println("Starting server....")

	err := godotenv.Load()
	if err != nil {
		return
	}
	port := os.Getenv("PORT")

	//ch := goHandlers.CORS(goHandlers.AllowedOrigins([]string{"*"}))
	//ch := goHandlers.CORS(goHandlers.AllowedMethods([]string{"GET, POST, HEAD, PUT, OPTIONS, DELETE"}))

	//run server
	srv := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", "0.0.0.0", port),
		Handler:     application.SetRouters(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("....server started")
	log.Fatal(srv.ListenAndServe())

}
