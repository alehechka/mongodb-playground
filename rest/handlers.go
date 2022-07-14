package rest

import (
	"github.com/alehechka/mongodb-playground/ginshared"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	engine := gin.Default()
	engine.Use(ginshared.CorsConfigMiddleware)

	RegisterHandlers(engine)

	return engine
}

func RegisterHandlers(engine *gin.Engine) {
	router := engine.Group("/api")

	router.GET("/podcasts", getPodcasts)
	router.GET("/podcasts/:id", getPodcast)
	router.POST("/podcasts/", createPodcast)
	router.PUT("/podcasts/:id", replacePodcast)
	router.DELETE("/podcasts/:id", deletePodcast)
}
