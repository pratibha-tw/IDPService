package main

import (
	"idp/configs"
	"idp/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	var cfg = configs.Config{}
	configs.GetConfigs(&cfg)

	router.RegisterRoutes(engine, cfg)
	engine.Run(":3000")

}
