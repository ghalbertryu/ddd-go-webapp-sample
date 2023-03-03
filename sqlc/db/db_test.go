package db

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCreateUser(t *testing.T) {
	ctx := context.Background()
	dbx, err := sql.Open("mysql", "root:1234@/sample")
	queries := New(dbx)
	if err != nil {
		fmt.Println(err)
	}

	queries.CreateUser(ctx, CreateUserParams{"xyz", "zyz123"})

}
