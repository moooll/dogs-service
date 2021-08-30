package storage

import (
	"errors"
	"sync"

	"github.com/google/uuid"
	"github.com/moooll/dogs-service/intenal/models"
)

type Storage struct {
	s  map[uuid.UUID]models.Dog
	mu *sync.Mutex
}

func NewStorage() *Storage {
	m := make(map[uuid.UUID]models.Dog)
	return &Storage{
		m,
		&sync.Mutex{},
	}	
}
func (st *Storage) Create(d models.Dog) error {
	st.mu.Lock()
	st.s[d.ID] = d
	st.mu.Unlock()
	return nil
}

func (st *Storage) Read(id uuid.UUID) (models.Dog, error) {
	st.mu.Lock()
	dog, ok := st.s[id]
	if !ok {
		return models.Dog{}, errors.New("dog not found")
	}

	return dog, nil
}

func (st *Storage) ReadAll() (dogs []models.Dog, err error) {
	for _, d := range st.s {
		dogs = append(dogs, d)
	}
	return dogs, nil
}

func (st *Storage) Update(d models.Dog) (models.Dog, error) {
	st.mu.Lock()
	st.s[d.ID] = d
	st.mu.Unlock()
	return d, nil
}

func (st *Storage) Delete(id uuid.UUID) error {
	delete(st.s, id)
	return nil
}
