package postgresql

import (
	"database/sql"
	"strconv"
	"time"

	"github.com/jiangzhifang/snippetbox/pkg/models"
)

type SnippetModel struct {
	DB *sql.DB
}

// func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
func (m *SnippetModel) Insert(title, content string, expires string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES($1, $2, $3, $4) RETURNING ID`

	var id int

	now := time.Now()
	days, err := strconv.Atoi(expires)
	if err != nil {
		panic("error")
	}

	err = m.DB.QueryRow(stmt, title, content, now, now.AddDate(0, 0, days)).Scan(&id)
	if err != nil {
		return 0, err
	}

	// The ID returned has the type int64, so we convert it to an int type
	// before returning.
	return id, nil

}

func (m *SnippetModel) Get(id int) (*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > CURRENT_TIMESTAMP AND id = $1`

	s := &models.Snippet{}
	err := m.DB.QueryRow(stmt, id).Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err == sql.ErrNoRows {
		return nil, models.ErrNoRecord
	} else if err != nil {
		return nil, err
	}

	return s, nil
}

// This will return the 10 most recently created snippets.
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > CURRENT_TIMESTAMP ORDER BY created DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	snippets := []*models.Snippet{}
	for rows.Next() {
		s := &models.Snippet{}
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return snippets, nil
}
