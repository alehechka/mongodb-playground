package types

import "go.mongodb.org/mongo-driver/bson/primitive"

// Episode is the type representing a podcast episode
type Episode struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Podcast     primitive.ObjectID `json:"podcastID,omitempty" bson:"p,omitempty"`
	Title       string             `json:"title,omitempty" bson:"t,omitempty"`
	Description string             `json:"description,omitempty" bson:"d,omitempty"`
	Duration    int32              `json:"duration,omitempty" bson:"du,omitempty"`
}
