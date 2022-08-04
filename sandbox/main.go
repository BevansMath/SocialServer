package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BevansMath/SocialServer/internal/database"
)

func main() {
	c := database.NewClient("db.json")
	err := c.EnsureDB()
	if err != nil {
		log.Fatal(err)
		os.Exit(1) // upon testing this code, an exit code of one was given.

	}
	user, err := c.CreateUser("test@example.com", "password", "Jane Doe", 18)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	fmt.Println("user created", user)

	updatedUser, err := c.UpdateUser("test@example.com", "password", "jane doe", 29)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)

	}
	fmt.Println("user updated", updatedUser)

	gotUser, err := c.GetUser("test@example.com")
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("got user", gotUser)

	err = c.DeleteUser("test@example.com")
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("user deleted")

	_, err = c.GetUser("test@example.com")
	if err == nil {
		log.Fatal("User unable to be retrieved")

	}
	fmt.Println("Confirmed removal of user...")

	user, err = c.CreateUser("test@example.com", "password", "john doe", 18)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user recreated", user)

	post, err := c.CreatePost("test@example.com", "The cat has FLEAS!!")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("post created", post)

	secondPost, err := c.CreatePost("test@example.com", "The fleas are EVERYWHERE!!!")
	if err != nil {
		log.Fatal(err)

	}
	fmt.Println("another post created", secondPost)

	posts, err := c.GetPosts("test@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted first post", posts)

	err = c.DeletePost(secondPost.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("deleted second post", posts)

	posts, err = c.GetPosts("test@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("got posts", posts)

	err = c.DeleteUser("test@example.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("user redeleted")
}
