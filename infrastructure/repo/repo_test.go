package repo

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
)

func TestCreateUser(t *testing.T) {
	dbx, err := sql.Open("mysql", "root:1234@/sample")
	queries := New(dbx)
	if err != nil {
		fmt.Println(err)
	}

	ctx := context.Background()
	//queries.CreateUser(ctx, CreateUserParams{"xyz", "zyz123"})
	user, _ := queries.GetUser(ctx, 1)
	fmt.Printf("%q", user)
}
