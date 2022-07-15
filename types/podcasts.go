package types

import (
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Podcast is the type representing a podcast
type Podcast struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title    string             `json:"title,omitempty" bson:"t,omitempty"`
	Author   string             `json:"author,omitempty" bson:"a,omitempty"`
	Tags     Tags               `json:"tags,omitempty" bson:"g,omitempty"`
	Episodes Episodes           `json:"episodes,omitempty" bson:"-"`
}

type Tags []string

func (t *Tags) ParseTags(rawTags string) {
	if len(rawTags) > 0 {
		*t = strings.Split(rawTags, ",")
	}
}

type Podcasts []Podcast

func (p *Podcasts) Init() {
	if *p == nil {
		*p = make(Podcasts, 0)
	}
}

type PodcastsResponse struct {
	Podcasts Podcasts `json:"podcasts"`
}

type PodcastResponse struct {
	Podcast Podcast `json:"podcast"`
}
