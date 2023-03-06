package repo

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	GDB *Queries
)

func Init() {
	dbx, err := sql.Open("mysql", "root:1234@/sample")
	GDB = New(dbx)
	if err != nil {
		fmt.Println(err)
	}
}
