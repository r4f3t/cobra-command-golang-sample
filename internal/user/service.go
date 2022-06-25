package user

type Service interface {
	GetUser() string
}

type service struct {
	repo string
}

func NewService(repo string) Service {
	return &service{
		repo: repo,
	}
}

func (receiver *service) GetUser() string {
	return "user bilgisi" + receiver.repo
}
