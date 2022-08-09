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

func (c Client) EnsureDB() error { // EnsureDB creates database file if it doesn't exist
	_, err := ioutil.ReadFile(c.dbPath)
	if errors.Is(err, os.ErrNotExist) {
		return c.createDB()
	}
	return err

}

func (c Client) createDB() error { // Creates database file
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

func (c Client) readDB() (databaseSchema, error) { // reads from database file
	dat, err := ioutil.ReadFile(c.dbPath)
	if err != nil {
		return databaseSchema{}, err
	}
	db := databaseSchema{}
	err = json.Unmarshal(dat, &db)
	return db, err
}

func (c Client) updateDB(db databaseSchema) error { // Updates database file
	dat, err := json.Marshal(db)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(c.dbPath, dat, 0600)
	if err != nil {
		return err
	}
	return nil
}
