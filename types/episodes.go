package types

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.opentelemetry.io/otel/attribute"
)

const EpisodeID string = "episodeID"

// Episode is the type representing a podcast episode
type Episode struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	PodcastID   primitive.ObjectID `json:"podcastID,omitempty" bson:"p,omitempty"`
	Title       string             `json:"title,omitempty" bson:"t,omitempty"`
	Description string             `json:"description,omitempty" bson:"d,omitempty"`
	Duration    int                `json:"duration,omitempty" bson:"u,omitempty"`
}

func (e Episode) String() string {
	return fmt.Sprintf("%#v", e)
}

func (e Episode) Attributes() (attrs []attribute.KeyValue) {
	attrs = append(attrs, attribute.String("id", e.ID.Hex()))
	attrs = append(attrs, attribute.String("title", e.Title))
	attrs = append(attrs, attribute.String("description", e.Description))
	attrs = append(attrs, attribute.Int("duration", int(e.Duration)))

	return
}

type Episodes []Episode

func (e *Episodes) Init() {
	if *e == nil {
		*e = make(Episodes, 0)
	}
}

type EpisodeResponse struct {
	Episode Episode `json:"episode"`
}

type EpisodesResponse struct {
	Episodes Episodes `json:"episodes"`
}
