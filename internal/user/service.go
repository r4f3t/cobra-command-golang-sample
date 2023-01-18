package user

type Service interface {
	GetUser() string
}

type service struct {
	userRepository UserRepository
}

func NewService(userRepo UserRepository) Service {
	return &service{
		userRepository: userRepo,
	}
}

func (receiver *service) GetUser() string {
	return "user info :" + receiver.userRepository.GetUserFromDb()
}
