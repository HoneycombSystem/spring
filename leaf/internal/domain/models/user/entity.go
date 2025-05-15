package user

type User struct {
	ID        int64  `json:"id"`
	Email	 string `json:"email"`
}

type UserAuth struct {
	ID        int64  `json:"id"`
	Password string `json:"password"`
}
