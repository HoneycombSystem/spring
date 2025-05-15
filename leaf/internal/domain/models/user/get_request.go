package user

type GetUserRequest struct {
	Email string `json:"email" validate:"required,email"`
}
type GetUserResponse struct {
	ID int64 `json:"id"`
}

