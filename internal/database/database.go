package database

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type Client struct { // For API database interactions
	dbPath string
}

type databaseSchema struct {
	Users map[string]User `json:"users"`
	Posts map[string]Post `json:"posts"`
}

func NewClient(dbPath string) Client { // Initializes a new client, returns client struct
	return Client{
		dbPath: dbPath,
	}
}

func (c Client) EnsureDB() error {
	_, err := ioutil.ReadFile(c.dbPath)
	if errors.Is(err, os.ErrNotExist) {
		return c.createDB()
	}
	return err

}

func (c Client) createDB() error {
	dat, err := json.Marshal(databaseSchema{
		Users: make(map[string]User),
		Posts: make(map[string]Post),
	})
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(c.dbPath, dat, 0600)
	if err != nil {
		return err
	}
	return nil
}
