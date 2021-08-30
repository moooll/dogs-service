// Package models provide models for endpoints and storage
package models

import (
	"github.com/google/uuid"
)

// Dog is the model for endpoints and storage
type Dog struct {
	ID    uuid.UUID `json:"id" govalid:"id"`
	Name  string    `json:"name" govalid:"req|min:1|max:30|regex:^[a-zA-Z]+$"`
	Breed string    `json:"breed" govalid:"req|max:100|regex:^[a-zA-Z]+$"`
	Color string    `json:"color" govalid:"max:100"`
	Price float32   `json:"price"`
	Age   float32   `json:"age"`
}
