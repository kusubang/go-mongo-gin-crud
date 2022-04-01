package services

import (
	"context"
	"errors"
	"go-mongodb/models"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceMongo struct {
	collection *mongo.Collection
}

func NewUserServiceMongo(collection *mongo.Collection) UserService {
	return &UserServiceMongo{collection}
}

func (u *UserServiceMongo) GetAllUsers() ([]models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	var users []models.User
	defer cancel()

	cursor, err := u.collection.Find(ctx, bson.M{})

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

func (u *UserServiceMongo) GetAUser(userId string) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.User
	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)
	err := u.collection.FindOne(ctx, bson.M{"id": objId}).Decode(&user)

	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (u *UserServiceMongo) CreateAUser(user *models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	d := u.collection.FindOne(ctx, bson.M{"email": user.Email})

	defer cancel()
	if err := d.Err(); err == nil {
		return nil, errors.New("User already exists")
	}

	newUser := models.User{
		Id:       primitive.NewObjectID(),
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	_, err := u.collection.InsertOne(ctx, newUser)
	if err != nil {
		return nil, err
	}
	return &newUser, nil
}

func (u *UserServiceMongo) EditAUser(userId string, user *models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	objId, _ := primitive.ObjectIDFromHex(userId)

	update := bson.M{"email": user.Email, "name": user.Name}
	result, err := u.collection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

	if err != nil {
		return nil, err
	}
	//get updated user details
	var updatedUser models.User
	if result.MatchedCount == 1 {
		err := u.collection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedUser)
		if err != nil {
			return nil, err
		}
	}
	return &updatedUser, errors.New("fail to update")
}

func (u *UserServiceMongo) DeleteAUser(userId string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(userId)

	result, err := u.collection.DeleteOne(ctx, bson.M{"id": objId})

	if err != nil {
		return err
	}

	if result.DeletedCount < 1 {
		return errors.New("User with specified ID not found!")
	}

	return nil
}
