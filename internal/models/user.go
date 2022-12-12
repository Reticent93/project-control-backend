package models

import "golang.org/x/crypto/bcrypt"

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (u *User) PasswordMatches(hashPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(hashPassword))
	if err != nil {
		switch {
		case err == bcrypt.ErrMismatchedHashAndPassword:
			//invalid password
			return false, nil
		default:
			//something else went wrong
			return false, err
		}
	}
	return true, nil
}
