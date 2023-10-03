package handlers

import (
	"auction-service/demand-side-service/models"
	"auction-service/demand-side-service/services"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// BidHandler handles HTTP requests related to bids and bidders.
type BidHandler struct {
	service services.BidService
}

// NewBidHandler creates a new BidHandler instance with the provided BidService.
func NewBidHandler(service services.BidService) BidHandler {
	return BidHandler{service: service}
}

// HandleRegisterBidder handles the registration of new bidders.
// It expects a JSON payload containing bidder information.
// If successful, it returns the created bidder's ID.
func (h *BidHandler) HandleRegisterBidder(c *gin.Context) {
	var bidder models.Bidder
	if err := c.ShouldBind(&bidder); err != nil {
		log.Println("invalid request payload. error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}
	if bidderID, err := h.service.CreateBidder(bidder); err != nil {
		log.Printf("failed to create bidder: %s. error: %v", bidder.Name, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{
			"message": "bidder registered successfully", "BidderID": bidderID})
	}
}

// HandleGetAllBidders retrieves and returns all registered bidders.
func (h *BidHandler) HandleGetAllBidders(c *gin.Context) {
	bidders, err := h.service.GetAllBidders()
	if err != nil {
		log.Println("invalid request payload. error:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Bidder not found"})
		return
	}
	c.JSON(http.StatusOK, bidders)

}

// HandleGetBidderByID retrieves a bidder by their ID and returns it.
func (h *BidHandler) HandleGetBidderByID(c *gin.Context) {
	bidderIDStr := c.Param("id")
	bidderID, err := strconv.Atoi(bidderIDStr)
	if err != nil {
		log.Println("invalid bidderID. error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid bidderID"})
		return
	}
	bidder, err := h.service.GetBidderById(bidderID)
	if err != nil {
		log.Println("error in handleGetBidderByID. error:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "bidder not found"})
		return
	}
	c.JSON(http.StatusOK, bidder)

}

// HandleGetBidsByAdSpaceID retrieves all bids for a specific ad space using its ID.
func (h *BidHandler) HandleGetBidsByAdSpaceID(c *gin.Context) {
	adSpaceIDStr := c.Param("id")
	adSpaceID, err := strconv.Atoi(adSpaceIDStr)
	if err != nil {
		log.Println("invalid ad_space_id. error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ad_space_id"})
		return
	}
	bids, err := h.service.GetBidsByAdSpaceID(adSpaceID)
	fmt.Println(err)
	if err != nil {
		log.Println("error in handleGetBidsByAdSpaceID. error:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "bids not found"})
		return
	}
	if bids == nil {
		log.Println("error in HandleGetBidsByAdSpaceID. error:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "no bids are placed for this ad-space"})
		return
	}
	c.JSON(http.StatusOK, bids)
}

// HandlePlaceBid handles placing a new bid for a specific ad space.
// It expects a JSON payload containing bid information.
// If successful, it returns the created bid's ID.
func (h *BidHandler) HandlePlaceBid(c *gin.Context) {
	var bidID int64
	bidderIDStr := c.Param("id")
	bidderID, err := strconv.Atoi(bidderIDStr)
	if err != nil {
		log.Println("invalid ad-space-id. Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ad-space-id"})
		return
	}
	var bid models.Bid
	bid.BidderID = bidderID
	if err := c.ShouldBindJSON(&bid); err != nil {
		log.Println("invalid request payload. error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	if bidID, err = h.service.PlaceBid(bid); err != nil {
		log.Println("internalServerError. error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "bid placed successfully", "bidID": bidID})
}

// HandleGetAllBidsByBidderID retrieves all bids placed by a specific bidder using their ID.
func (h *BidHandler) HandleGetAllBidsByBidderID(c *gin.Context) {
	adSpaceIDStr := c.Param("id")
	adSpaceID, err := strconv.Atoi(adSpaceIDStr)
	if err != nil {
		log.Println("invalid ad-space-id. error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ad-space-id"})
		return
	}
	bids, err := h.service.GetAllBidsByBidderID(adSpaceID)
	if err != nil {
		log.Println("error in handleGetAllBidsByBidderID. error:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "bids not found"})
		return
	}
	if bids == nil {
		log.Println("error in handleGetAllBidsByBidderID. error:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "no bids are placed by the bidder"})
		return
	}
	c.JSON(http.StatusOK, bids)
}

// HandleGetAllBidsByBidderIDAndAdSpaceID retrieves all bids placed by a specific bidder
// for a specific ad space using their IDs.
func (h *BidHandler) HandleGetAllBidsByBidderIDAndAdSpaceID(c *gin.Context) {
	adSpaceIDStr := c.Param("adspaceID")
	adSpaceID, err := strconv.Atoi(adSpaceIDStr)
	if err != nil {
		log.Println("invalid ad-space-id. error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ad-space-id"})
		return
	}
	bidderIDStr := c.Param("id")
	bidderID, err := strconv.Atoi(bidderIDStr)
	if err != nil {
		log.Println("invalid bidder-id. Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid bidder-id"})
		return
	}
	bids, err := h.service.GetAllBidsByBidderIDAndAdSpaceID(bidderID, adSpaceID)
	if err != nil {
		log.Println("error in getAllBidsByBidderIDAndAdSpaceID. error:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "bids not found"})
		return
	}
	if bids == nil {
		log.Println("error in getAllBidsByBidderIDAndAdSpaceID. error:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "no bids are placed by the bidder for this ad-space"})
		return
	}
	c.JSON(http.StatusOK, bids)
}
