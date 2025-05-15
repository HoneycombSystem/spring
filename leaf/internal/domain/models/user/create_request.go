package user

type CreateNewUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
type CreateNewUserResponse struct {
	ID int64 `json:"id"`
}

