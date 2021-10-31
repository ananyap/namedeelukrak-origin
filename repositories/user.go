package repositories

type User struct {
	UserId   int    `db:"user_id" json:"user_id"`
	Username string `db:"user_nanme" json:"user_name"`
	Passwd   string `db:"passwd" json:"passwd"`
	Email    string `db:"email" json:"email"`
}

type UserSignupRequest struct {
	Username string `json:"user_name"`
	Passwd   string `json:"passwd"`
	Email    string `json:"email"`
}

type UserLoginRequest struct {
	Username string `json:"user_name"`
	Passwd   string `json:"passwd"`
}

type UserRepository interface {
	Signup(userSignup UserSignupRequest) (*User, error)
	Login(userLogin User) (*User, error)
}
