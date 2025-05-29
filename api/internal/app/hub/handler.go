package hub

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yanisdib/tomodachi/domain/hub"
)

type HubHandler struct {
	service *HubService
}

func NewHubHandler(service *HubService) *HubHandler {
	return &HubHandler{service: service}
}

func (h *HubHandler) CreateHubHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 10*time.Second)
		var input hub.CreateHubRequest
		defer cancel()

		if err := c.ShouldBindBodyWithJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"error":   err.Error(),
				"message": "Failed to parse request body. Ensure that the request body is a valid JSON object.",
			})

			return
		}

		newHub, err := h.service.CreateHub(&input, ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"error":   err.Error(),
				"message": hub.ErrHubCreationFailed.Error(),
			})

			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusCreated,
			"message": "Hub created successfully.",
			"item":    newHub,
		})
	}
}

func (h *HubHandler) GetHubsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 10*time.Second)
		defer cancel()

		hubs, err := h.service.GetHubs(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"error":   err.Error(),
				"message": hub.ErrHubsFetchFailed.Error(),
			})
			return
		}

		if len(hubs) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"status":  http.StatusOK,
				"message": hub.ErrNoHubsFound.Error(),
				"items":   []hub.Hub{},
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  http.StatusOK,
			"message": "Hubs fetched successfully.",
			"items":   hubs,
		})
	}
}
