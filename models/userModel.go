package models

type User struct {
FirstName string `bson:"FirstName"`
Email string `bson:"Email"`
Username string `bson:"Username"`
}