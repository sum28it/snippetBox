package models

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErrDuplicateEmail     = errors.New("models: email address already in use")
	ErrInvalidCredentials = errors.New("models: invalid user credentials")
)

type Database struct {
	*sql.DB
}

func (db *Database) GetSnippet(id int) (*Snippet, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > CURRENT_TIMESTAMP AND id = $1`

	row := db.QueryRow(stmt, id)

	s := &Snippet{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return s, nil
}

func (db *Database) LatestSnippets() (Snippets, error) {
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > CURRENT_TIMESTAMP ORDER BY created DESC LIMIT 10`

	rows, err := db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	snippets := Snippets{}

	for rows.Next() {
		s := &Snippet{}

		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
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

func (db *Database) InsertSnippet(title, content, expires string) (int, error) {
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP+ ($3 || 'SECOND')::INTERVAL) RETURNING id`

	fmt.Println("Debugging", "database.InsertSnippet72")

	rows, err := db.Query(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	fmt.Println("Debugging", "database.InsertSnippet78")
	var id int

	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			fmt.Println(err)
			return 0, nil
		}
	}
	fmt.Println("Debugging", "database.InsertSnippet84")
	return id, nil
}

func (db *Database) InsertUser(name, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, hashed_password, created)
	VALUES($1, $2, $3, CURRENT_TIMESTAMP)`

	_, err = db.Exec(stmt, name, email, hashedPassword)
	if err != nil {
		pqErr, _ := err.(*pq.Error)
		if pqErr.Code == "23505" {
			return ErrDuplicateEmail
		}
	}
	return err
}

func (db *Database) VerifyUser(email, password string) (int, error) {
	var id int
	var hashedPassword []byte

	row := db.QueryRow("SELECT id, hashed_password FROM users WHERE email = $1", email)
	err := row.Scan(&id, &hashedPassword)
	if err == sql.ErrNoRows {
		return 0, ErrInvalidCredentials
	} else if err != nil {
		return 0, err
	}
	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return 0, ErrInvalidCredentials
	} else if err != nil {
		return 0, err
	}
	return id, nil
}
