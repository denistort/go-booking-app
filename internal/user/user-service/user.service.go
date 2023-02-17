package userService

type UserService struct {
}

func New() *UserService {
	return &UserService{}
}

func (r *UserService) SignIn() {}

func (r *UserService) SignUp() {}
