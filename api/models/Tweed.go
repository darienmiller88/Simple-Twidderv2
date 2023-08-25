package models

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/kamva/mgm/v3"
)

type Tweed struct{
	mgm.DefaultModel `bson:",inline"`
	Name    string   `json:"name"    bson:"name"`
	Content string   `json:"content" bson:"content"`
}

func (t Tweed) Validate() error{
	return validation.ValidateStruct(
		&t,
		validation.Field(&t.Name, validation.Required, validation.Length(3, 30)),
		validation.Field(&t.Content, validation.Required, validation.Length(2, 100)),
	)
}