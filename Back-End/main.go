package main

import (
	"TRS/app/midwares"
	"TRS/config/database"
	"TRS/config/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	r := gin.Default()
	r.NoMethod(midwares.HandleNotFound)
	r.NoRoute(midwares.HandleNotFound)
	router.Init(r)
	err := r.Run()
	if err != nil {
		log.Fatal("Server start error:", err)
	}
}
