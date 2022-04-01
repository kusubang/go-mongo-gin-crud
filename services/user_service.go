package services

import (
	"context"
	"go-mongodb/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllUsers(collection *mongo.Collection) ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var users []models.User
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var aUser models.User
		err := cursor.Decode(&aUser)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, aUser)
	}

	return users, nil
}
