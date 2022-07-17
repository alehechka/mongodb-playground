package rest

import (
	"github.com/alehechka/go-utils/ginshared"
	"github.com/alehechka/mongodb-playground/opentel"
	"github.com/alehechka/mongodb-playground/types"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func SetupRouter() *gin.Engine {
	engine := gin.Default()
	engine.Use(ginshared.CorsConfigMiddleware)
	engine.Use(otelgin.Middleware("mongodb-playground-api", otelgin.WithTracerProvider(opentel.TracerProvider)))

	RegisterHandlers(engine)

	return engine
}

func RegisterHandlers(engine *gin.Engine) {
	router := engine.Group("/api")
	router.GET("")

	{
		podcasts := router.Group("/podcasts")
		podcasts.GET("", getPodcasts)
		podcasts.GET("/:"+types.PodcastID, getPodcast)
		podcasts.POST("", createPodcast)
		podcasts.PUT("/:"+types.PodcastID, replacePodcast)
		podcasts.DELETE("/:"+types.PodcastID, deletePodcast)

		{
			episodes := podcasts.Group("/:" + types.PodcastID + "/episodes")
			episodes.GET("", getPodcastEpisodes)
			episodes.GET("/:"+types.EpisodeID, getPodcastEpisode)
			episodes.POST("", createPodcastEpisode)
			episodes.PUT("/:"+types.EpisodeID, replacePodcastEpisode)
			episodes.DELETE("/:"+types.EpisodeID, deletePodcastEpisode)
		}
	}
}
