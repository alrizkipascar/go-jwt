package models

type LoginRequest struct {
	Email    string `json:"email"`
	Number   int64  `json:"number"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Number int64  `json:"number"`
	Token  string `json:"token"`
}

type LoginResponseEmail struct {
	Email string `json:"email"`
	Token string `json:"token"`
}
