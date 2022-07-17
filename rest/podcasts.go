package rest

import (
	"net/http"

	"github.com/alehechka/go-utils/ginshared"
	"github.com/alehechka/mongodb-playground/database"
	"github.com/alehechka/mongodb-playground/opentel"
	"github.com/alehechka/mongodb-playground/types"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.opentelemetry.io/otel/attribute"
)

func getPodcasts(c *gin.Context) {
	ctx, span := opentel.GinTracer.Start(c.Request.Context(), "getPodcasts")
	defer span.End()

	filterPodcast := types.Podcast{
		Title:  c.Query("filter[title]"),
		Author: c.Query("filter[author]"),
	}
	filterPodcast.Tags.ParseTags(c.Query("filter[tags]"))

	span.SetAttributes(filterPodcast.Attributes()...)

	podcasts, err := database.FindPodcasts(ctx, filterPodcast)
	if ginshared.ShouldAbortWithError(c)(http.StatusInternalServerError, err) {
		span.RecordError(err)
		return
	}

	c.JSON(http.StatusOK, types.PodcastsResponse{Podcasts: podcasts})
}

func getPodcast(c *gin.Context) {
	ctx, span := opentel.GinTracer.Start(c.Request.Context(), "getPodcast")
	defer span.End()

	podcastID, err := primitive.ObjectIDFromHex(c.Param(types.PodcastID))
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		span.RecordError(err)
		return
	}
	span.SetAttributes(attribute.String(types.PodcastID, podcastID.Hex()))

	var podcast types.Podcast
	includeEpisodes := ginshared.GetIncludedParams(c).IsIncluded("episodes")
	span.SetAttributes(attribute.Bool("includeEpisodes", includeEpisodes))
	if includeEpisodes {
		podcast, err = database.GetPodcastWithEpisodes(ctx, podcastID)
	} else {
		podcast, err = database.FindPodcast(ctx, types.Podcast{ID: podcastID})
	}

	if ginshared.ShouldAbortWithError(c)(http.StatusNotFound, err) {
		span.RecordError(err)
		return
	}

	c.JSON(http.StatusOK, types.PodcastResponse{Podcast: podcast})
}

func createPodcast(c *gin.Context) {
	ctx, span := opentel.GinTracer.Start(c.Request.Context(), "createPodcast")
	defer span.End()

	var podcast types.Podcast
	err := c.ShouldBind(&podcast)
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		span.RecordError(err)
		return
	}
	span.SetAttributes(podcast.Attributes()...)

	podcastID, err := database.InsertPodcast(ctx, podcast)
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		span.RecordError(err)
		return
	}

	c.Params = append(c.Params, gin.Param{Key: types.PodcastID, Value: podcastID.Hex()})
	getPodcast(c)
}

func replacePodcast(c *gin.Context) {
	ctx, span := opentel.GinTracer.Start(c.Request.Context(), "replacePodcast")
	defer span.End()

	podcastID, err := primitive.ObjectIDFromHex(c.Param(types.PodcastID))
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		span.RecordError(err)
		return
	}
	span.SetAttributes(attribute.String(types.PodcastID, podcastID.Hex()))

	var podcast types.Podcast
	err = c.ShouldBind(&podcast)
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		span.RecordError(err)
		return
	}
	span.SetAttributes(podcast.Attributes()...)

	err = database.ReplacePodcast(ctx, podcastID, podcast)
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		span.RecordError(err)
		return
	}

	getPodcast(c)
}

func deletePodcast(c *gin.Context) {
	ctx, span := opentel.GinTracer.Start(c.Request.Context(), "deletePodcast")
	defer span.End()

	podcastID, err := primitive.ObjectIDFromHex(c.Param(types.PodcastID))
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		span.RecordError(err)
		return
	}
	span.SetAttributes(attribute.String(types.PodcastID, podcastID.Hex()))

	err = database.DeletePodcast(ctx, podcastID)
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		span.RecordError(err)
		return
	}

	c.Status(http.StatusNoContent)
}
