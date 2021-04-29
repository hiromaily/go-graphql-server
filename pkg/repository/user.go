package repository

import (
	"encoding/json"
	"io/ioutil"

	"github.com/hiromaily/go-graphql-server/pkg/user"
)

type userMap struct {
	repo map[string]user.UserType
}

func NewUserMapRepo() (user.User, error) {
	data, err := importJSONDataFromFile("./assets/user.json")
	if err != nil {
		return nil, err
	}
	return &userMap{
		repo: data,
	}, nil
}

func (u *userMap) Fetch(id string) user.UserType {
	return u.repo[id]
}

// Helper function to import json from file to map
func importJSONDataFromFile(fileName string) (map[string]user.UserType, error) {
	var data map[string]user.UserType
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(content, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
