package main

import (
	"auction-service/supply-side-service/database"
	"auction-service/supply-side-service/handlers"
	"auction-service/supply-side-service/repositories"
	"auction-service/supply-side-service/services"
	"log"

	"github.com/gin-gonic/gin"
	// Import other necessary packages
)

func main() {
	db, err := database.NewMySQLConnection()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	router := gin.Default()

	adspaceRepository := repositories.NewAdSpaceRepository(db)
	adspaceService := services.NewAdSpaceService(adspaceRepository)
	adspaceHandler := handlers.NewAdSpaceHandler(adspaceService)

	router.POST("supply-service/adspaces", adspaceHandler.CreateAdSpace)
	router.GET("supply-service/adspaces", adspaceHandler.GetAllAdSpaces)
	router.GET("supply-service/adspaces/:id", adspaceHandler.GetAdSpaceByID)
	router.GET("supply-service/adspaces/:id/winner", adspaceHandler.GetWinner)
	// Add other routes

	router.Run(":8080")

}
