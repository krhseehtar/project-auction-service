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

	router.POST("/demand_service/bidders", bidHandler.HandleRegisterBidder)
	router.GET("/demand_service/bidders", bidHandler.HandleGetAllBidders)
	router.GET("/demand_service/bidders/:id", bidHandler.HandleGetBidderByID)
	router.POST("/demand_service/bidders/:id/bids", bidHandler.HandlePlaceBid)
	router.GET("/demand_service/adspaces/:id/bids", bidHandler.HandleGetBidsByAdSpaceID)
	router.GET("/demand_service/bidders/:id/bids", bidHandler.HandleGetAllBidsByBidderID)
	router.GET("/demand_service/bidders/:id/adspaces/:adspaceID/bids", bidHandler.HandleGetAllBidsByBidderIDAndAdSpaceID)
	router.Run(":8081")
}
