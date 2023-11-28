package db

import (
	"context"
	"log"

	"github.com/pageza/mock-UrbanEats-Foodtruck/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

// Connect establishes a connection to the database
func Connect() {
	cfg := config.LoadConfig()
	clientOptions := options.Client().ApplyURI(cfg.MongoURI)
	var err error
	Client, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = Client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB")
}
func GetMenuItems() ([]MenuItem, error) {
	var menuItems []MenuItem
	// Database logic to fetch menu items
	// Example: Fetching data from MongoDB
	collection := Client.Database("yourDatabaseName").Collection("menuItems")
	cursor, err := collection.Find(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var menuItem MenuItem
		err := cursor.Decode(&menuItem)
		if err != nil {
			return nil, err
		}
		menuItems = append(menuItems, menuItem)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return menuItems, nil
}

// GetCollection returns a handle to a collection in the database
func GetCollection(collectionName string) *mongo.Collection {
	return Client.Database("urbaneats").Collection(collectionName)
}
