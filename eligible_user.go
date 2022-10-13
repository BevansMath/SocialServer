package main

import (
	"errors"
	"fmt"
)

// Test cases for invalid user inputs
func eligibleUser(email, password string, age int) error {
	if email == "" {
		return errors.New("email cannot be empty")
	}
	if password == "" {
		return errors.New("password cannot be empty")
	}
	const leastAge = 18
	if age < 18 {
		return fmt.Errorf("age must be at least %v years old", leastAge)
	}
	return nil
}
