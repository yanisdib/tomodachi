package platform

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CreatePlatformInput struct {
	Name            int    `json:"name" binding:"required"`
	Server          string `json:"server" binding:"required"`
	NitroLevel      uint8  `json:"nitro_level" binding:"required"`
	VideoResolution string `json:"video_resolution" binding:"required"`
	NetworkQuality  uint8  `json:"network_quality" binding:"required"`
	AccessURL       string `json:"access_url"`
}

// Create creates a new platform
// @Summary Create a new platform
// @Description Create a new platform
// @Tags platform
// @Accept json
// @Produce json
// @Param input body CreateInput true "Platform input"
// @Success 201 {object} Platform
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /platforms [post]
func Create(dbPool *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c, 10*time.Second)
		var input CreatePlatformInput
		defer cancel()

		if err := c.ShouldBindBodyWithJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"error":   err.Error(),
				"message": "Failed to parse request body. Ensure that the request body is a valid JSON object.",
			})

			return
		}

		repository := NewPgRepository(dbPool)
		newPlatform, err := repository.Store(ctx, &input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"error":   err.Error(),
				"message": ErrCreateFailed.Error(),
			})

			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusCreated,
			"message": "Platform created successfully",
			"data":    newPlatform,
		})
	}
}
