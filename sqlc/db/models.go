// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.2

package db

import (
	"time"
)

// user
type User struct {
	// user id
	ID int64 `json:"id"`
	// account
	Account string `json:"account"`
	// password
	Password string `json:"password"`
	// create_time
	CreateTime time.Time `json:"create_time"`
	// update_time
	UpdateTime time.Time `json:"update_time"`
}
