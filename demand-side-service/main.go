package main

import (
	"auction-service/demand-side-service/database"
	"auction-service/demand-side-service/handlers"
	"auction-service/demand-side-service/repositories"
	"auction-service/demand-side-service/services"
	"log"

	"github.com/gin-gonic/gin"
	//"net/http"
	// Import other necessary packages
)

func main() {
	db, err := database.NewMySQLConnection()
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	defer db.Close()

	router := gin.Default()

	bidRepository := repositories.NewBidRepository(db)
	bidService := services.NewBidService(bidRepository)
	bidHandler := handlers.NewBidHandler(bidService)

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
	router.Run(":8081")
}
