package repositories

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
)

type UserRepositoryDB struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return UserRepositoryDB{db}
}

func (r UserRepositoryDB) Signup(userSignup UserSignupRequest) (*User, error) {
	query := "INSERT users (user_name, passwd, email) values (?, ?, ?)"
	result, err := r.db.Exec(query, userSignup.Username, userSignup.Passwd, userSignup.Email)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())
	}

	user := User{
		UserId:   int(id),
		Username: userSignup.Username,
		Passwd:   userSignup.Passwd,
		Email:    userSignup.Email,
	}

	return &user, nil
}

func (r UserRepositoryDB) Login(userLogin User) (*User, error) {
	return nil, nil
}
