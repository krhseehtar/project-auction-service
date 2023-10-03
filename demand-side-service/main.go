package main

import (
	"auction-service/demand-side-service/database"
	"auction-service/demand-side-service/handlers"
	"auction-service/demand-side-service/repositories"
	"auction-service/demand-side-service/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// get db connection
	db, err := database.NewMySQLConnection()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	router := gin.Default()

	// Initialize repositories
	bidRepository := repositories.NewBidRepository(db)
	// Initialize services
	bidService := services.NewBidService(bidRepository)
	// Initialize handlers
	bidHandler := handlers.NewBidHandler(bidService)

	//group routes
	biddersRouter := router.Group("/demand_service/bidders")
	{
		biddersRouter.POST("", bidHandler.HandleRegisterBidder)
		biddersRouter.GET("", bidHandler.HandleGetAllBidders)
		biddersRouter.GET("/:id", bidHandler.HandleGetBidderByID)
		biddersRouter.POST("/:id/bids", bidHandler.HandlePlaceBid)
		biddersRouter.GET("/:id/bids", bidHandler.HandleGetAllBidsByBidderID)
		biddersRouter.GET("/:id/adspaces/:adspaceID/bids", bidHandler.HandleGetAllBidsByBidderIDAndAdSpaceID)
	}
	adspaceRouter := router.Group("/demand_service/adspaces")
	{
		adspaceRouter.GET("/:id/bids", bidHandler.HandleGetBidsByAdSpaceID)

	}
	// listen and serve on 0.0.0.0:8080
	router.Run(":8081")
}
