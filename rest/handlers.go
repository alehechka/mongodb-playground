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

	{
		podcasts := router.Group("/podcasts")
		podcasts.GET("", getPodcasts)
		podcasts.GET("/:podcastID", getPodcast)
		podcasts.POST("", createPodcast)
		podcasts.PUT("/:podcastID", replacePodcast)
		podcasts.DELETE("/:podcastID", deletePodcast)

		{
			episodes := podcasts.Group("/:podcastID/episodes")
			episodes.GET("", getPodcastEpisodes)
			episodes.GET("/:episodeID", getPodcastEpisode)
			episodes.POST("", createPodcastEpisode)
			episodes.PUT("/:episodeID", replacePodcastEpisode)
			episodes.DELETE("/:episodeID", deletePodcastEpisode)
		}
	}
}
