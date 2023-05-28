package http

type userRepository interface {
	Add(username string, email string, password string, role string) error
	CheckPasswordByEmail(email string, password string) (bool, error)
	GetUsernameByEmail(email string) (string, error)
}

type authService interface {
	GetUserJWT(email string) (string, error)
	GetUserEmailByJWT(raw string) (string, error)
}
