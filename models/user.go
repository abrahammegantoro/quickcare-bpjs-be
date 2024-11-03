package models

type User struct {
	Model
	Name     string `bson:"name"`
	Email    string `bson:"email"`
	Username string `bson:"username"`
	Password string `bson:"password"`
}
