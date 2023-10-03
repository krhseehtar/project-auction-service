package main

import (
	"auction-service/supply-side-service/database"
	"auction-service/supply-side-service/handlers"
	"auction-service/supply-side-service/repositories"
	"auction-service/supply-side-service/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.NewMySQLConnection()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	router := gin.Default()

	// Initialize repositories
	adspaceRepository := repositories.NewAdSpaceRepository(db)
	// Initialize services
	adspaceService := services.NewAdSpaceService(adspaceRepository)
	// Initialize handlers
	adspaceHandler := handlers.NewAdSpaceHandler(adspaceService)

	//group routes
	adspaceRouter := router.Group("/supply-service/adspaces")
	{
		adspaceRouter.POST("", adspaceHandler.CreateAdSpace)
		adspaceRouter.GET("", adspaceHandler.GetAllAdSpaces)
		adspaceRouter.GET("/:id", adspaceHandler.GetAdSpaceByID)
		adspaceRouter.GET("/:id/winner", adspaceHandler.GetWinner)

	}

	// listen and serve on 0.0.0.0:8080
	router.Run(":8080")

}
