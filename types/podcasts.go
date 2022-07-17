package types

import (
	"encoding/json"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const PodcastID string = "podcastID"

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

func (t Tags) String() string {
	return strings.Join(t, ",")
}

type Podcasts []Podcast

func (p *Podcasts) Init() {
	if *p == nil {
		*p = make(Podcasts, 0)
	}
}

func (p Podcast) String() string {
	b, _ := json.Marshal(p)
	return string(b)
}

type PodcastsResponse struct {
	Podcasts Podcasts `json:"podcasts"`
}

type PodcastResponse struct {
	Podcast Podcast `json:"podcast"`
}
