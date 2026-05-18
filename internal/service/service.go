package service

type Services struct {
	Users *UserService
}
type UserService struct {
	store UserStore
}

func NewUserService(serv UserStore) *UserService {
	return &UserService{
		store: serv,
	}
}
