package user

type UserService struct {
	DB interface{}
}

type Repository interface {
	CreateUser(user string) error
}

func NewUserService(db interface{}) Repository {
	return &UserService{
		DB: db,
	}
}
