package user

type User struct {
	Id       uint64 `json:"uid"`
	Account  string `json:"acc"`
	Password string `json:"pwd"`
}

func NewUser(account string) *User {
	return &User{Account: account}
}
