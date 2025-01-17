package config

import (
	"github.com/gin-gonic/gin"
)

func initRouter() error {
	router := gin.Default()
	router.POST("/platforms", nil)
	return nil
}
