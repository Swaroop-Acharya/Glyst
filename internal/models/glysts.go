package models

import (
	"database/sql"
	"errors"
	"time"
)

type GlystModelInterface interface {
	Insert(title string, content string, expires int) (int, error)
	Get(id int) (Glyst, error)
	Latest() ([]Glyst, error)
}

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
	// Write the SQL statement we want to execute. I've split it over two lines
	// for readability (which is why it's surrounded with backquotes instead
	// of normal double quotes).
	stmt := `INSERT INTO glysts (title,content,created,expires)
		VALUES(?,?,UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP,INTERVAL ? DAY))`

	// Use the Exec() method on the embedded connection pool to execute the
	// statement. The first parameter is the SQL statement, followed by the
	// values for the placeholder parameters: title, content and expiry in
	// that order. This method returns a sql.Result type, which contains some
	// basic information about what happened when the statement was executed
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, nil
	}

	// Use the LastInsertId() method on the result to get the ID of our
	// newly inserted record in the snippets table.
	id, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}

	// The ID returned has the type int64, so we convert it to an int type
	// before returning.
	return int(id), nil
}

// This will return a specific snippet based on its id
func (m *GlystModel) Get(id int) (Glyst, error) {
	stmt := `SELECT id, title, content, created, expires FROM glysts
	where expires > UTC_TIMESTAMP() and id = ?`

	var g Glyst
	err := m.DB.QueryRow(stmt, id).Scan(&g.ID, &g.Title, &g.Content, &g.Created, &g.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Glyst{}, ErrNoRecord
		} else {
			return Glyst{}, err
		}
	}

	return g, nil
}

// This will return the 10 most recently created snippets.
func (m *GlystModel) Latest() ([]Glyst, error) {
	// Write the SQL statement we want to execute.
	stmt := `SELECT id, title, content, created, expires FROM glysts
	where expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

	// Use the Query() method on the connection pool to execute our
	// SQL statement. This returns a sql.Rows resultset containing the result of
	// our query.
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	// We defer rows.Close() to ensure the sql.Rows resultset is
	// always properly closed before the Latest() method returns. This defer
	// statement should come *after* you check for an error from the Query()
	// method. Otherwise, if Query() returns an error, you'll get a panic
	// trying to close a nil resultset.
	defer rows.Close()

	// Initialize an empty slice to hold the Snippet structs.
	var glysts []Glyst

	// Use rows.Next to iterate through the rows in the resultset. This
	// prepares the first (and then each subsequent) row to be acted on by the
	// rows.Scan() method. If iteration over all the rows completes then the
	// resultset automatically closes itself and frees-up the underlying
	// database connection.
	for rows.Next() {
		// Create a new zeroed Snippet struct.
		var g Glyst

		// Use rows.Scan() to copy the values from each field in the row to the
		// new Snippet object that we created. Again, the arguments to row.Scan()
		// must be pointers to the place you want to copy the data into, and the
		// number of arguments must be exactly the same as the number of
		// columns returned by your statement.
		err = rows.Scan(&g.ID, &g.Title, &g.Content, &g.Created, &g.Expires)
		if err != nil {
			return nil, err
		}
		// Append it to the slice of snippets.
		glysts = append(glysts, g)
	}

	// When the rows.Next() loop has finished we call rows.Err() to retrieve any
	// error that was encountered during the iteration. It's important to
	// call this - don't assume that a successful iteration was completed
	// over the whole resultset.
	if err = rows.Err(); err != nil {
		return nil, err
	}

	// If everything went OK then return the Snippets slice.
	return glysts, nil
}
