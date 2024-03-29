package main

import (
	"errors"
	"testing"
)

// Unit Testing
func TestEligibleUser(t *testing.T) {
	var tests = []struct {
		email       string
		password    string
		age         int
		expectedErr error
	}{
		{
			email:       "test@example.com",
			password:    "12345",
			age:         18,
			expectedErr: nil,
		},
		{
			email:       "",
			password:    "qwerty18",
			age:         18,
			expectedErr: errors.New("email cannot be empty"),
		},
		{
			email:       "test2@example.com",
			password:    "",
			age:         18,
			expectedErr: errors.New("password cannot be empty"),
		},
		{
			email:       "test3@example.com",
			password:    "dirtyqwerty",
			age:         16,
			expectedErr: errors.New("age must be at least 18 years old"),
		},
	}

	for _, tt := range tests {
		err := eligibleUser(tt.email, tt.password, tt.age)
		errString := ""
		expectedErrString := ""
		if err != nil {
			errString = err.Error()
		}
		if tt.expectedErr != nil {
			expectedErrString = tt.expectedErr.Error()
		}
		if errString != expectedErrString {
			t.Errorf("got %s, want %s", errString, expectedErrString)
		}
	}
}
