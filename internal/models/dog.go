// Package models provide models for endpoints and storage
package models

import (
	"github.com/google/uuid"
)

// Dog is the model for endpoints and storage
type Dog struct {
	// ID describes dog id
	ID uuid.UUID `json:"id" govalid:"id"`
	// Name describes dog name
	Name string `json:"name" govalid:"req|min:1|max:30|regex:^[a-zA-Z]+$"`
	// Breed describes dog's breed
	Breed string `json:"breed" govalid:"req|max:100|regex:^[a-zA-Z]+$"`
	// Color descibes dog's color
	Color string `json:"color" govalid:"max:100"`
	// Price describes dog's price
	Price float32 `json:"price"`
	// Age describes dog's age
	Age float32 `json:"age"`
}
