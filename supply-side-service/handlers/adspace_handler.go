package handlers

import (
	"auction-service/supply-side-service/models"
	"auction-service/supply-side-service/services"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdSpaceHandler struct {
	service services.AdSpaceService
}

func NewAdSpaceHandler(service services.AdSpaceService) AdSpaceHandler {
	return AdSpaceHandler{service: service}
}

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
