package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id       primitive.ObjectID `bson:"id,omitempty"`
	Email    string             `bson:"email,omitempty"`
	Name     string             `bson:"name"`
	Password string             `bson:"password"`
}
