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

	adspaceRouter := router.Group("/supply-service/adspaces")
	{
		adspaceRouter.POST("", adspaceHandler.CreateAdSpace)
		adspaceRouter.GET("", adspaceHandler.GetAllAdSpaces)
		adspaceRouter.GET("/:id", adspaceHandler.GetAdSpaceByID)
		adspaceRouter.GET("/:id/winner", adspaceHandler.GetWinner)

	}

	router.Run(":8080")

}
