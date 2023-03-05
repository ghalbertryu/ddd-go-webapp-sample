package db

import (
	"context"
	"m/domain/user"
	"m/infrastructure/repo"
)

type DbRepository struct {
	Queries *repo.Queries
}

func NewDbRepository(queries *repo.Queries) *DbRepository {
	return &DbRepository{Queries: queries}
}

func (dr *DbRepository) Get(id int64) (user.User, error) {
	dbUser, err := dr.Queries.GetUser(context.Background(), 1)
	if err == nil {
		return user.User{uint64(dbUser.ID), dbUser.Account, dbUser.Password}, err
	}

	return user.User{}, err
}
