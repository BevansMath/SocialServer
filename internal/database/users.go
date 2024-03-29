package database

import (
	"errors"
	"time"
)

//user required fields
type User struct {
	CreatedAt time.Time `json:"createdAt"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
}

// Create User
func (c Client) CreateUser(
	email, password, name string, age int,
) (User, error) {
	db, err := c.readDB()
	if err != nil {
		return User{}, err
	}
	if _, ok := db.Users[email]; ok {
		return User{}, errors.New("user already exists")
	}
	user := User{
		CreatedAt: time.Now().UTC(),
		Email:     email,
		Password:  password,
		Name:      name,
		Age:       age,
	}
	db.Users[email] = user
	err = c.updateDB(db)
	if err != nil {
		return User{}, err
	}
	return user, err
}

// Update user
func (c Client) UpdateUser(
	email, password, name string,
	age int,
) (User, error) {
	db, err := c.readDB()
	if err != nil {
		return User{}, err
	}
	user, ok := db.Users[email]
	if !ok {
		return User{}, errors.New("user doesn't exist")

	}
	user.Password = password
	user.Name = name
	user.Age = age
	db.Users[email] = user
	err = c.updateDB(db)
	if err != nil {
		return User{}, err

	}
	return user, err

}

// Get User
func (c Client) GetUser(email string) (User, error) {
	db, err := c.readDB()
	if err != nil {
		return User{}, err

	}
	user, ok := db.Users[email]
	if !ok {
		return User{}, errors.New("user doesn't exist")

	}
	return user, nil
}

// Delete User
func (c Client) DeleteUser(email string) error {
	db, err := c.readDB()
	if err != nil {
		return err

	}
	delete(db.Users, email)
	err = c.updateDB(db)
	if err != nil {
		return err
	}
	return nil

}
