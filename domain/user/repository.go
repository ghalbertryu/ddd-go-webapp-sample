package user

type UserRepository interface {
	Get(int64) (User, error)
	Add(User) error
	Update(User) error
}
