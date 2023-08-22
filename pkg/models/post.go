package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Post represents a blog post, as modeled in the database.
// This is no longer in use, as I've deprecated the blog page, but it's here as
// an example for future DB-based projects.
type Post struct {
	Title string             `json:"title" bson:"title"`
	Body  string             `json:"body" bson:"body"`
	Date  primitive.DateTime `json:"date" bson:"date"`
}
