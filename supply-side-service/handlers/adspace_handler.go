package handlers

import (
	"auction-service/supply-side-service/models"
	"auction-service/supply-side-service/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AdSpaceHandler handles HTTP requests related to ad spaces.
type AdSpaceHandler struct {
	service services.AdSpaceService
}

// NewAdSpaceHandler creates a new AdSpaceHandler instance with the provided AdSpaceService.
func NewAdSpaceHandler(service services.AdSpaceService) AdSpaceHandler {
	return AdSpaceHandler{service: service}
}

// GetAllAdSpaces handles the HTTP request to retrieve all ad spaces.
func (h *AdSpaceHandler) GetAllAdSpaces(c *gin.Context) {
	adSpaces, err := h.service.GetAllAdSpaces()
	if err != nil {
		log.Println("internal server error. error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	if adSpaces == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no adspaces are available"})
	} else {
		c.JSON(http.StatusOK, adSpaces)
	}
}

// GetAdSpaceByID handles the HTTP request to retrieve an ad space by its ID.
func (h *AdSpaceHandler) GetAdSpaceByID(c *gin.Context) {
	adSpaceIDStr := c.Param("id")
	adSpaceID, err := strconv.Atoi(adSpaceIDStr)
	if err != nil {
		log.Println("invalid ad-space-id. Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ad-space-id"})
		return
	}
	adSpace, err := h.service.GetAdSpaceByID(adSpaceID)
	if err != nil {
		log.Println("error in getAdSpaceByID(). error:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "adSpace not found"})
		return
	}
	c.JSON(http.StatusOK, adSpace)
}

// CreateAdSpace handles the HTTP request to create a new ad space.
func (h *AdSpaceHandler) CreateAdSpace(c *gin.Context) {
	var adSpace models.AdSpace
	if err := c.ShouldBindJSON(&adSpace); err != nil {
		log.Println("invalid request payload. Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}
	if adSpaceID, err := h.service.CreateAdSpace(adSpace); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusCreated, gin.H{"message": "ad space created successfully", "adSpaceID": adSpaceID})
	}

}

// GetWinner handles the HTTP request to retrieve the winner of an ad space auction.
func (h *AdSpaceHandler) GetWinner(c *gin.Context) {
	adSpaceIDStr := c.Param("id")
	adSpaceID, err := strconv.Atoi(adSpaceIDStr)
	if err != nil {
		log.Println("invalid ad-space-id. error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ad-space-id"})
		return
	}
	winnerID, err := h.service.GetWinner(adSpaceID)
	if err != nil {
		log.Println("error in GetWinner(). error:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"winnerID": winnerID})
}
