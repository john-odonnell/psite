package db

import (
	"context"

	"github.com/john-odonnell/psite/v2/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateClient returns a pointer to an already-initialized mongo.Client
// instance.
func CreateClient() (*mongo.Client, error) {
	opts := options.Client().ApplyURI(
		"mongodb+srv://guest:5JfLmOgS7nvCH1Sg@cluster0.7nsmj.mongodb.net/blogdb?retryWrites=true&w=majority",
	)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil
}

// ConnectToDatabase connects an initialized mongo.Client to a named collection
// in a named database.
func ConnectToDatabase(client *mongo.Client, db string, collection string) *mongo.Collection {
	return client.Database(db).Collection(collection)
}

// GetBlogPosts uses an already-configured mongo.Collection to draw all elements
// from the database.
// This is no longer in use, as I've deprecated the blog page, but it's here as
// an example for future DB-based projects.
func GetBlogPosts(collection *mongo.Collection) ([]models.Post, error) {
	cursor, err := collection.Find(context.TODO(), bson.M{}, options.Find())
	if err != nil {
		return nil, err
	}
	defer func() {
		cursor.Close(context.Background())
	}()

	var posts []models.Post
	for cursor.Next(context.TODO()) {
		post := models.Post{}
		err := cursor.Decode(&post)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}
