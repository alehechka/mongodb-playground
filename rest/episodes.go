package rest

import (
	"net/http"

	"github.com/alehechka/go-utils/ginshared"
	"github.com/alehechka/mongodb-playground/database"
	"github.com/alehechka/mongodb-playground/types"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func getPodcastEpisodes(c *gin.Context) {
	podcastID, err := primitive.ObjectIDFromHex(c.Param("podcastID"))
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		return
	}

	findEpisode := types.Episode{
		Title:       c.Query("filter[title]"),
		Description: c.Query("filter[description]"),
	}

	episodes, err := database.FindPodcastEpisodes(podcastID, findEpisode)
	if ginshared.ShouldAbortWithError(c)(http.StatusInternalServerError, err) {
		return
	}

	c.JSON(http.StatusOK, types.EpisodesResponse{Episodes: episodes})
}

func getPodcastEpisode(c *gin.Context) {
	podcastID, err := primitive.ObjectIDFromHex(c.Param("podcastID"))
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		return
	}

	episodeID, err := primitive.ObjectIDFromHex(c.Param("episodeID"))
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		return
	}

	episode, err := database.FindPodcastEpisode(podcastID, types.Episode{ID: episodeID})
	if ginshared.ShouldAbortWithError(c)(http.StatusNotFound, err) {
		return
	}

	c.JSON(http.StatusOK, types.EpisodeResponse{Episode: episode})
}

func createPodcastEpisode(c *gin.Context) {
	podcastID, err := primitive.ObjectIDFromHex(c.Param("podcastID"))
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		return
	}

	var episode types.Episode
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, c.ShouldBind(&episode)) {
		return
	}

	episodeID, err := database.InsertPodcastEpisode(podcastID, episode)
	if ginshared.ShouldAbortWithError(c)(http.StatusInternalServerError, err) {
		return
	}

	c.Params = append(c.Params, gin.Param{Key: "episodeID", Value: episodeID.Hex()})
	getPodcastEpisode(c)
}

func replacePodcastEpisode(c *gin.Context) {
	podcastID, err := primitive.ObjectIDFromHex(c.Param("podcastID"))
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		return
	}

	episodeID, err := primitive.ObjectIDFromHex(c.Param("episodeID"))
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		return
	}

	var episode types.Episode
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, c.ShouldBind(&episode)) {
		return
	}

	err = database.ReplacePodcastEpisode(podcastID, episodeID, episode)
	if ginshared.ShouldAbortWithError(c)(http.StatusInternalServerError, err) {
		return
	}

	getPodcastEpisode(c)
}

func deletePodcastEpisode(c *gin.Context) {
	podcastID, err := primitive.ObjectIDFromHex(c.Param("podcastID"))
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		return
	}

	episodeID, err := primitive.ObjectIDFromHex(c.Param("episodeID"))
	if ginshared.ShouldAbortWithError(c)(http.StatusBadRequest, err) {
		return
	}

	err = database.DeletePodcastEpisode(podcastID, episodeID)
	if ginshared.ShouldAbortWithError(c)(http.StatusNotFound, err) {
		return
	}

	c.Status(http.StatusNoContent)
}
