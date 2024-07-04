package models

import "encoding/json"

type Registration struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (r *Registration) Save() (User, error) {
	return User{ID: 1, login: "admin"}, nil
}

func NewRegistrationFromJSON(data []byte) (*Registration, error) {
	r := &Registration{}
	err := json.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}
	return r, nil
}
