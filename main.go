package main

import (
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	ID       uint64 `bson:"id,omitempty"`
	Email    string `bson:"email,omitempty"`
	Name     string `bson:"name"`
	Password string `bson:"password"`
}

const MONGO_URI = "mongodb://1234:1234@127.0.0.1:27017/"

func initMongo(uri string) *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	return client
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/users", getUsersHandler)
	router.GET("/users/:email", getUserHandler)
	router.POST("/users", addUserHandler)

	router.Run("localhost:8080")
	return router
}

var client *mongo.Client

func main() {
	client = initMongo(MONGO_URI)
	defer client.Disconnect(context.Background())

	r := setupRouter()
	r.Run()

	// podcastsCollection := client.Database("test").Collection("users")

	// user := User{
	// 	ID:       0,
	// 	Name:     "Hong",
	// 	Password: "1234",
	// }

	// podcastResult, err := podcastsCollection.InsertOne(context.TODO(), user)

	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(podcastResult.InsertedID)

}

// func main3() {
// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MONGO_URI))
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer client.Disconnect(ctx)

// 	database := client.Database("quickstart")
// 	podcastsCollection := database.Collection("podcasts")
// 	episodesCollection := database.Collection("episodes")

// 	var episodes []Episode
// 	cursor, err := episodesCollection.Find(ctx, bson.M{"duration": bson.D{{"$gt", 25}}})
// 	if err != nil {
// 		panic(err)
// 	}
// 	if err = cursor.All(ctx, &episodes); err != nil {
// 		panic(err)
// 	}

// 	podcast := Podcast{
// 		Title:  "The Polyglot Developer",
// 		Author: "Nic Raboy",
// 		Tags:   []string{"development", "programming", "coding"},
// 	}
// 	insertResult, err := podcastsCollection.InsertOne(ctx, podcast)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(insertResult.InsertedID)
// 	fmt.Println(episodes)

// }

// func main2() {
// 	client, err := mongo.NewClient(options.Client().ApplyURI(MONGO_URI))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
// 	err = client.Connect(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer client.Disconnect(ctx)

// 	err = client.Ping(ctx, readpref.Primary())
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	databases, err := client.ListDatabaseNames(ctx, bson.M{})
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	fmt.Println(databases)

// }
