package tokens

import (
	"time"

	"github.com/google/uuid"
	"github.com/moooll/dogs-service/internal/models"
)

// NewSession creates new session with provided refresh token
func NewSession(userID uuid.UUID, token string) models.Session {
	id := uuid.New()
	fingerprint := uuid.New()
	return models.Session{
		ID:           id,
		UserID:       userID,
		Fingerprint:  fingerprint,
		RefreshToken: token,
		ExpiresAt:    time.Now().Add(43200 * time.Minute),
	}
}

// NewResfreshToken generates new refresh string token
func NewResfreshToken() string {
	return uuid.NewString()
}
