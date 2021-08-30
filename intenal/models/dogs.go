// Package models provide models for endpoints and storage
package models

import (
	"github.com/google/uuid"
)

// Dog is the model for endpoints and storage
type Dog struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Breed string    `json:"breed"`
	Color string    `json:"color"`
	Price float32   `json:"price"`
	Age   float32   `json:"age"`
}
