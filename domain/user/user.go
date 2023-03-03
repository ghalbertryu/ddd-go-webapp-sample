package user

type User struct {
	Id   int    `json:"uid"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
