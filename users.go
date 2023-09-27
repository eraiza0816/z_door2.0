package users

import (
	"encoding/json"
	"fmt"
	"os"
)

type Users map[string]string

func LoadUsers() (Users, error) {
	file, err := os.ReadFile("user.conf")
	if err != nil {
		return nil, fmt.Errorf("failed to read user.conf: %w", err)
	}

	var users Users
	err = json.Unmarshal(file, &users)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal user.conf: %w", err)
	}

	return users, nil
}
