package database

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Post Fields
type Post struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UserEmail string    `json:"userEmail"`
	Text      string    `json:"text"`
}

// Create posts
func (c Client) CreatePost(
	userEmail, text string,
) (Post, error) {
	db, err := c.readDB()
	if err != nil {
		return Post{}, err
	}
	if _, ok := db.Users[userEmail]; !ok {
		return Post{}, errors.New("user doesn't exist")

	}
	id := uuid.New().String()
	post := Post{
		ID:        id,
		CreatedAt: time.Now().UTC(),
		UserEmail: userEmail,
		Text:      text,
	}
	db.Posts[id] = post
	err = c.updateDB(db)
	if err != nil {
		return Post{}, err
	}
	return post, err

}

// Get posts
func (c Client) GetPosts(userEmail string) ([]Post, error) {
	db, err := c.readDB()
	if err != nil {
		return nil, err

	}
	posts := []Post{}
	for _, post := range db.Posts {
		if post.UserEmail == userEmail {
			posts = append(posts, post)
		}

	}
	return posts, nil

}

// Delete posts
func (c Client) DeletePost(id string) error {
	db, err := c.readDB()
	if err != nil {
		return err
	}
	delete(db.Posts, id)
	err = c.updateDB(db)
	if err != nil {
		return err
	}
	return nil
}
