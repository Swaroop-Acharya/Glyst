package models

import (
	"database/sql"
	"time"
)

// Define a Snippet type to hold the data for an individual snippet. Notice how
// the fields of the struct correspond to the fields in our MySQL snippets
// table?

type Glyst struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

// Define a SnippetModel type which wraps a sql.DB connection pool
type GlystModel struct {
	DB *sql.DB
}

// This will insert a new snippet into the database.
func (m *GlystModel) Insert(title string, content string, expires int) (int, error) {
	return 0, nil
}

// This will return a specific snippet based on its id
func (m *GlystModel) Get(id int) (Glyst, error) {
	return Glyst{}, nil
}

// This will return the 10 most recently created snippets.
func (m *GlystModel) Latest() ([]Glyst, error) {
	return nil, nil
}
