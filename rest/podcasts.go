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
	podcastID, err := primitive.ObjectIDFromHex(c.Param("podcastID"))
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		return
	}

	var podcast types.Podcast
	if ginshared.GetIncludedParams(c).IsIncluded("episodes") {
		podcast, err = database.GetPodcastWithEpisodes(podcastID)
	} else {
		podcast, err = database.FindPodcast(types.Podcast{ID: podcastID})
	}

	if ginshared.ShouldAbortWithError(c)(http.StatusNotFound, err) {
		return
	}

	c.JSON(http.StatusOK, types.PodcastResponse{Podcast: podcast})
}

func createPodcast(c *gin.Context) {
	var podcast types.Podcast
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, c.ShouldBind(&podcast)) {
		return
	}

	podcastID, err := database.InsertPodcast(podcast)
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		return
	}

	c.Params = append(c.Params, gin.Param{Key: "podcastID", Value: podcastID.Hex()})
	getPodcast(c)
}

func replacePodcast(c *gin.Context) {
	podcastID, err := primitive.ObjectIDFromHex(c.Param("podcastID"))
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		return
	}

	var podcast types.Podcast
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, c.ShouldBind(&podcast)) {
		return
	}

	err = database.ReplacePodcast(podcastID, podcast)
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		return
	}

	getPodcast(c)
}

func deletePodcast(c *gin.Context) {
	podcastID, err := primitive.ObjectIDFromHex(c.Param("podcastID"))
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		return
	}

	err = database.DeletePodcast(podcastID)
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		return
	}

	c.Status(http.StatusNoContent)
}
