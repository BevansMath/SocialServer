package main

import (
	"fmt"
	"log"
	"os"

	"github.com/BevansMath/SocialServer/internal/database"
)

func main() {
	c := database.NewClient("db.json")
	err := c.EnsureDB
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

	updatedUser, err := c.UpdatedUser("test@example.com", "password", "jane doe", 29)
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
		log.Fatal("User not retrieved")

	}
	fmt.Println("Confirmed removal of user...")
}
