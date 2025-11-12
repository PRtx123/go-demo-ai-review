package service

import (
	"context"
	"database/sql"
	"myapp/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func (s *UserService) GetUser(ctx context.Context, id int) (*repository.User, error) {
	user, err := s.Repo.FindByID(ctx, id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}
	return user, nil
}
