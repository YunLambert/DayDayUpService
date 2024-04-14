package main

import (
	"daydayup/cron"
	"daydayup/db"
	"daydayup/router"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	database := db.Init()
	defer database.Close()

	cron.InitAutoResetStreak(&database)
	r := router.Init(database)

	gin.SetMode(gin.DebugMode)
	err := http.ListenAndServe(":19042", r)
	if err != nil {
		log.Fatal(err.Error())
	}
}
