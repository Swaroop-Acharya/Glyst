package mocks

import (
	"time"

	"glyst/internal/models"
)

var mockGlyst = models.Glyst{
	ID:      1,
	Title:   "A mock glyst title",
	Content: "A mock glyst content",
	Created: time.Now(),
	Expires: time.Now(),
}

type GlystModel struct{}

func (m *GlystModel) Insert(title, content string, expires int) (int, error) {
	return 2, nil
}

func (m *GlystModel) Get(id int) (models.Glyst, error) {
	switch id {
	case 1:
		return mockGlyst, nil
	default:
		return models.Glyst{}, models.ErrNoRecord
	}
}

func (m *GlystModel) Latest() ([]models.Glyst, error) {
	return []models.Glyst{mockGlyst}, nil
}
