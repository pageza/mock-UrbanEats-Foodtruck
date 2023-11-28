package db

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CustomerFeedback struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	CustomerName string             `bson:"customerName"`
	Feedback     string             `bson:"feedback"`
	Rating       int                `bson:"rating"`
	Date         time.Time          `bson:"date"`
}
