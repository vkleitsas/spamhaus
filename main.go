package main

import (
	"assignment/app"
	"assignment/db"
	h "assignment/http"
	"assignment/utils"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {

	mongo, err := db.NewMongoStore()
	if err != nil {
		log.Fatal(fmt.Sprintf("%d:%s", http.StatusInternalServerError, err))
	}
	client := &http.Client{}
	repo := db.NewRepository(*mongo)

	utils := utils.NewUrlDownload(client)

	crudService := app.NewCrudService(repo, utils)

	saveNewWebsiteHandler := h.NewSaveNewWebsiteUrlHandler(*crudService)
	websiteRoutes := h.InitWebsiteRoutes(*saveNewWebsiteHandler)

	mainRouter := h.InitMainRouter(websiteRoutes)
	s := gocron.NewScheduler(time.UTC)
	var intervalSecs int
	interval := os.Getenv("INTERVAL")
	if interval == "" {
		intervalSecs = 60
		log.Printf("Defaulting to interval %ds", intervalSecs)
	}
	intervalSecs, err = strconv.Atoi(interval)
	if err != nil {
		intervalSecs = 60
		log.Printf("Failed to convert INTERVAL env variable to int, defaulting to interval %ds", intervalSecs)
	}
	s.Every(intervalSecs).Seconds().Do(func() {
		log.Print("Scheduled func start")
		res, _ := repo.GetMostSubmitted()
		utils.DownloadUrls(res)

	})

	s.StartAsync()
	mainMuxRouter := mainRouter.InitRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), mainMuxRouter))

}
