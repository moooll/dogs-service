// Package storage contains functions for working with storage, like reading, creating, deleting and updating
package storage

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/moooll/dogs-service/internal/models"
)

// Storage constains map storage and mutex
type Storage struct {
	s  map[uuid.UUID]models.Dog
	mu *sync.Mutex
}

// NewStorage returns new initialized storage
func NewStorage() *Storage {
	m := make(map[uuid.UUID]models.Dog)
	return &Storage{
		m,
		&sync.Mutex{},
	}
}

// Create creates new record in a map
func (st *Storage) Create(d models.Dog) error {
	st.mu.Lock()
	st.s[d.ID] = d
	st.mu.Unlock()
	return nil
}

// Read reads record by id and returns error if not found
func (st *Storage) Read(id uuid.UUID) (models.Dog, error) {
	st.mu.Lock()
	dog, ok := st.s[id]
	if !ok {
		return models.Dog{}, errors.New("dog not found")
	}

	return dog, nil
}

// ReadAll reads all records
func (st *Storage) ReadAll() (dogs []models.Dog, err error) {
	for _, d := range st.s {
		dogs = append(dogs, d)
	}
	return dogs, nil
}

// Update updates record to what's provided in d
func (st *Storage) Update(d models.Dog) (models.Dog, error) {
	st.mu.Lock()
	st.s[d.ID] = d
	st.mu.Unlock()
	return d, nil
}

// Delete deletes a record by id
func (st *Storage) Delete(id uuid.UUID) error {
	delete(st.s, id)
	return nil
}
