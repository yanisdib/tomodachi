package hub

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func RegisterRoutes(router *gin.Engine, dbPool *pgxpool.Pool) {

	repository := NewHubRepository(dbPool)
	service := NewHubService(repository)
	handler := NewHubHandler(service)

	router.GET("api/v1/hubs", handler.GetHubsHandler())
	router.GET("api/v1/hub", nil)
	router.GET("api/v1/hub/:id", nil)
	router.POST("api/v1/hub", handler.CreateHubHandler())
	router.PUT("api/v1/hub/:id", nil)
	router.DELETE("api/v1/hub/:id", nil)
}
