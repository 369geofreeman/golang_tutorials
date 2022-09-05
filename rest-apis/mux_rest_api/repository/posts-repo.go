package repository

import (
	"context"
	"log"

	"../entity"
	"cloud.google.com/go/firestore"
)

type PostRpository interface {
	// return referance to the object and and error
	Save(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type repo struct {
}

// NewPostRepository
func NewPostRepository() PostRepository {
	return &repo{}
}

const (
	projectId      string = "golang-rest-api"
	collectionName string = "Posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)

	if err != nil {
		log.Printf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}

	defer client.Close()
	_, _, err = clinet.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {{
		log.Printf("Failed adding a new post: %v", err)
		return nil, err
	}
	return post, nil
}


func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := firestore.NewClient(ctx, projectId)

	if err != nil {
		log.Printf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}
	
	defer client.Close()
	var posts []entity.Post
	iterator := client.Collection(collectionName).Documents(ctx)
	for {
		doc, err := iterator.Next()
		if err != nil {
			log.Printf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}
		post := entity.Post{
			ID: doc.Data()["ID"].(int),
			Title: doc.Data()["Title"].(string),
			Text: doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}