package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ContactFormSubmission struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Name      string             `bson:"name"`
	Email     string             `bson:"email"`
	Message   string             `bson:"message"`
	Submitted time.Time          `bson:"submitted"`
}
