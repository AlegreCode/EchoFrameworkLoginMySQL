package models

import (
	"errors"
	"database/sql"
	"fmt"

	. "github.com/alegrecode/echo/LoginMySQL/db"

	"github.com/gookit/validate"

	"github.com/labstack/echo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id,omitempty" form:"id,omitempty"`
	Name     string `json:"name" form:"name" validate:"required|minLen:3|maxLen:50"`
	Lastname string `json:"lastname" form:"lastname" validate:"required|minLen:3|maxLen:50"`
	Email    string `json:"email" form:"email" validate:"required|email"`
	Age      string `json:"age" form:"age" validate:"required|min:1|max:150"`
	Password string `json:"password" form:"password" validate:"required|minLen:4"`
}

func (f User) ConfigValidation(v *validate.Validation) {
	v.StopOnError = false
	v.WithScenes(validate.SValues{
		"login":    []string{"Email", "Password"},
		"register": []string{"Name", "Lastname", "Email", "Age", "Password"},
	})
}

func (f User) Messages() map[string]string {
	return validate.MS{
		"required":    "The {field} field is required.",
		"Email.email": "Email not valid.",
	}
}

func SaveUser(c echo.Context) sql.Result {
	user := new(User)
	if err := c.Bind(user); err != nil {
		fmt.Println(err.Error())
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hash)

	sql := "INSERT INTO users (name, lastname, email, age, password) VALUES (?,?,?,?,?);"
	stmt, _ := DB.Prepare(sql)
	defer stmt.Close()

	result, _ := stmt.Exec(user.Name, user.Lastname, user.Email, user.Age, user.Password)
	return result
}

func GetSingleUser(emailForm string) (User, error) {
	var user User
	var id, name, lastname, email, age, password string

	sql := "SELECT id, name, lastname, email, age, password FROM users WHERE email=?;"
	err := DB.QueryRow(sql, emailForm).Scan(&id, &name, &lastname, &email, &age, &password)
	if err != nil {
		return user, errors.New("User not found.")
	}

	user = User{ID: id, Name: name, Lastname: lastname, Email: email, Age: age, Password: password}
	return user, nil
}
