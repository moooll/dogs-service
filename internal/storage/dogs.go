// Package storage contains functions for working with storage, like reading, creating, deleting and updating
package storage

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/moooll/dogs-service/internal/models"
)

// Storage constains map storage and mutex
type Storage struct {
	conn *pgxpool.Pool
	ctx  context.Context
}

// NewStorage returns new initialized storage
func NewStorage(ctx context.Context, conn *pgxpool.Pool) *Storage {
	return &Storage{
		ctx:  ctx,
		conn: conn,
	}
}

// Create creates new record in a map
func (st *Storage) Create(d models.Dog) error {
	row, err := st.conn.Query(st.ctx, "insert into dogs(id, name, breed, color, price, age) values ($1, $2, $3, $4, $5, $6)",
		d.ID,
		d.Name,
		d.Breed,
		d.Color,
		d.Age,
		d.Price)
	if err != nil {
		return err
	}

	row.Close()
	return nil
}

// Read reads record by id and returns error if not found
func (st *Storage) Read(id uuid.UUID) (d models.Dog, err error) {
	err = st.conn.QueryRow(st.ctx, "select * from dogs where id = $1", id).Scan(&d.ID,
		&d.Name,
		&d.Breed,
		&d.Color,
		&d.Age,
		&d.Price)
	if err != nil {
		return models.Dog{}, err
	}

	return d, nil
}

// ReadAll reads all records
func (st *Storage) ReadAll() (dogs []models.Dog, err error) {
	rows, err := st.conn.Query(st.ctx, "select * from dogs")
	if err != nil {
		return []models.Dog{}, err
	}

	d := models.Dog{}
	for rows.Next() {
		er := rows.Scan(&d.ID, &d.Name, &d.Breed, &d.Color, &d.Age, &d.Price)
		if er != nil {
			return []models.Dog{}, er
		}

		dogs = append(dogs, d)
	}
	return dogs, nil
}

// Update updates record to what's provided in d
func (st *Storage) Update(d models.Dog) (err error) {
	row, err := st.conn.Query(st.ctx, `insert into dogs(id, name, breed, color, age, price) 
		values ($1, $2, $3, $4, $5, $6) on conflict (id) do update set 
		name = $2, breed = $3, color = $4, age = $5, price = $6`,
		d.ID,
		d.Name,
		d.Breed,
		d.Color,
		d.Age,
		d.Price)
	if err != nil {
		return err
	}

	row.Close()

	return nil
}

// Delete deletes a record by id
func (st *Storage) Delete(id uuid.UUID) error {
	row, err := st.conn.Query(st.ctx, "delete from dogs where id = $1", id)
	if err != nil {
		return err
	}

	row.Close()
	return nil
}
