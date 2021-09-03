package storage

import (
	"github.com/google/uuid"
	"github.com/moooll/dogs-service/internal/models"
)

// GetSession retrieves refresh session from the db
func (st *Storage) GetSession(refresh uuid.UUID) (s models.Session, err error) {
	err = st.conn.QueryRow(st.ctx,
		"select * from refresh_sessions where refresh_token = $1",
		refresh).Scan(&s.ID, &s.UserID, &s.RefreshToken, &s.Fingerprint, &s.ExpiresAt)
	if err != nil {
		return models.Session{}, err
	}

	return s, nil
}

// CreateSession creates new row in refresh_session
func (st *Storage) CreateSession(s *models.Session) error {
	row, err := st.conn.Query(st.ctx,
		`insert into refresh_sessions(id, user_id, refresh_token, fingerprint, expires_at)
		values ($1, $2, $3, $4, $5)`, s.ID, s.UserID, s.RefreshToken, s.Fingerprint, s.ExpiresAt)
	if err != nil {
		return err
	}

	row.Close()

	return nil
}

// DeleteSession deletes refresh session from the DB by refresh token
func (st *Storage) DeleteSession(refresh uuid.UUID) error {
	row, err := st.conn.Query(st.ctx,
		"delete from refresh_sessions where refresh_token = $1",
		refresh)
	if err != nil {
		return err
	}

	row.Close()

	return nil
}

// UpdateSession updates refresh session by s.ID
func (st *Storage) UpdateSession(s *models.Session) error {
	row, err := st.conn.Query(st.ctx,
		"update refresh_sessions set fingerprint = $1, refresh_token = $2, expires_at = $3 where id = $4",
		s.Fingerprint,
		s.RefreshToken,
		s.ExpiresAt,
		s.ID)
	if err != nil {
		return err
	}

	row.Close()

	return nil
}
