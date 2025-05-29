package main

import (
	"log"
	"sync"

	"github.com/yanisdib/tomodachi/internal/app/hub"
	"github.com/yanisdib/tomodachi/internal/config"
)

func main() {
	var once sync.Once

	once.Do(func() {
		dbPool := config.OpenDBConnectionPool()
		defer dbPool.Close()

		router := config.ConfigRouter()
		hub.RegisterRoutes(router, dbPool)

		if err := router.Run(":5050"); err != nil {
			log.Fatal("an error occured while starting the server at localhost:5050")
		}
	})

}
