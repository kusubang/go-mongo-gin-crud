package main

type User struct {
	ID       uint64 `bson:"id,omitempty"`
	Email    string `bson:"email,omitempty"`
	Name     string `bson:"name"`
	Password string `bson:"password"`
}
