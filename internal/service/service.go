package service

type Services struct {
	Users *UserService
	Posts *PostService
}
type UserService struct {
	store UserStore
}
type PostService struct {
	store PostStore
}

func NewUserService(serv UserStore) *UserService {
	return &UserService{
		store: serv,
	}
}

func NewPostService(service PostStore) *PostService {
	return &PostService{
		store: service,
	}
}
