package service

type LoginService interface {
	Login(username string, password string) bool
}

type loginService struct {
	authorizedUsername string
	authorizedPassword string
}

func (l *loginService) Login(username string, password string) bool {
	return l.authorizedUsername == username && l.authorizedPassword == password
}

func NewLoginService() LoginService{
	return &loginService{
		authorizedUsername: "babay",
		authorizedPassword: "babay12345",
	}
}

