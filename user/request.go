package user

type SignupRequest struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

type LoginRequest struct {
	Email    string `binding:"required"`
	Password string `binding:"required"`
}