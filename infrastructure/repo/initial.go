package repo

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var (
	GLog log.Logger // 全局log
	GDB  *Queries   // 全局 DB
)

func InitRepository() {
	db, err := sql.Open("mysql", "root:1234@/sample")
	GDB = New(db)
	if err != nil {
		fmt.Println(err)
	}

	//queries.CreateUser(context.Background(), CreateUserParams{"xyz", "zyz123"})
}
