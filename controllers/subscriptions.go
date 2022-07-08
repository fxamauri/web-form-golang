package controllers

import "github.com/fxamauri/web-form-golang/db"

type Subscription struct {
	Name  string
	Email string
}

func Create(name string, email string) error {
	s := Subscription{Name: name, Email: email}

	return db.Insert("newsletter", s)
}
