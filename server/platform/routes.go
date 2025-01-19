package platform

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewPlatformRoutes(router *gin.Engine, dbPool *pgxpool.Pool) {
	router.POST("v1/platform", Create(dbPool))
}
