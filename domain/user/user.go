package user

import "m/infrastructure/repo"

type User struct {
	Id       int64  `json:"uid"`
	Account  string `json:"acc"`
	Password string `json:"pwd"`
}

func NewUser(account string) *User {
	return &User{Account: account}
}

func ConvertUser(userPo repo.User) *User {
	return &User{userPo.ID, userPo.Account, userPo.Password}
}
