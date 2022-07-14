package rest

import (
	"net/http"

	"github.com/alehechka/mongodb-playground/database"
	"github.com/alehechka/mongodb-playground/ginshared"
	"github.com/alehechka/mongodb-playground/types"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getPodcasts(c *gin.Context) {
	findPodcast := types.Podcast{
		Title:  c.Query("filter[title]"),
		Author: c.Query("filter[author]"),
	}
	findPodcast.Tags.ParseTags(c.Query("filter[tags]"))

	podcasts, err := database.FindPodcasts(findPodcast)
	if ginshared.ShouldAbortWithError(c)(http.StatusInternalServerError, err) {
		return
	}

	c.JSON(http.StatusOK, types.PodcastsResponse{Podcasts: podcasts})
}

func getPodcast(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		return
	}

	podcast, err := database.FindPodcast(types.Podcast{ID: id})

	c.JSON(http.StatusOK, types.PodcastResponse{Podcast: podcast})
}

func createPodcast(c *gin.Context) {
	var podcast types.Podcast
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, c.ShouldBind(&podcast)) {
		return
	}

	id, err := database.InsertPodcast(podcast)
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		return
	}

	created, err := database.FindPodcast(types.Podcast{ID: id})
	if ginshared.ShouldAbortWithError(c)(http.StatusInternalServerError, err) {
		return
	}

	c.JSON(http.StatusCreated, types.PodcastResponse{Podcast: created})
}

func replacePodcast(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		return
	}

	var podcast types.Podcast
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, c.ShouldBind(&podcast)) {
		return
	}

	err = database.ReplacePodcast(id, podcast)
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		return
	}

	getPodcast(c)
}

func deletePodcast(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		return
	}

	err = database.DeletePodcast(id)
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		return
	}

	c.Status(http.StatusNoContent)
}
