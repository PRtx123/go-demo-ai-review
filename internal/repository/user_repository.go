package repository

import (
	"context"
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

type User struct {
	ID   int
	Name string
	Age  int
}

func (r *UserRepository) FindByID(ctx context.Context, id int) (*User, error) {
	var u User
	err := r.DB.QueryRowContext(ctx, "SELECT id, name, age FROM users WHERE id = ?", id).
		Scan(&u.ID, &u.Name, &u.Age)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
