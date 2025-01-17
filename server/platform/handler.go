package platform

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateInput struct {
	Name            int    `json:"id"`
	Server          string `json:"server,omitempty"`
	NitroLevel      string `json:"nitroLevel,omitempty"`
	VideoResolution string `json:"videoResolution,omitempty"`
	NetworkQuality  uint8  `json:"networkQuality,omitempty"`
	AccessURL       string `json:"accessURL,omitempty"`
}

func Create() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var input *CreateInput
		defer cancel()

		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  http.StatusBadRequest,
				"error":   err.Error(),
				"message": "Failed to parse JSON",
			})

			return
		}
	}
}
