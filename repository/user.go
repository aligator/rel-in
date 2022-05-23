package repository

import (
	"context"
	"github.com/go-rel/rel"
	"rel-in/entity"
)

type UserRepository interface {
	FindAll(limit int) ([]entity.User, error)
}

type userRepository struct {
	repo rel.Repository
}

func NewUserRepository(repo rel.Repository) UserRepository {
	return &userRepository{repo}
}

func (ur userRepository) FindAll(limit int) ([]entity.User, error) {
	var (
		query = rel.Select()
	)

	if limit > 0 {
		query = query.Limit(limit)
	}

	users := []entity.User{}

	err := ur.repo.FindAll(context.Background(), &users, query)
	if err != nil {
		return nil, err
	}
	return users, nil
}
