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
		os.Exit(1)

	}
	fmt.Println("database ensured!")
}
