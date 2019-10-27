package users

import (
	"fmt"
	"net/http"

	. "github.com/alegrecode/echo/LoginMySQL/middlewares"
	. "github.com/alegrecode/echo/LoginMySQL/models"

	"github.com/labstack/echo"
)

func LoginView(c echo.Context) error {
	flash := c.(*CustomContext).GetFlash()
	return c.Render(http.StatusOK, "login.html", map[string]interface{}{
		"flash": flash,
	})
}

func LoginUser(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "/dashboard")
}

func DashboardView(c echo.Context) error {
	auth := c.(*CustomContext).Auth()
	return c.Render(http.StatusOK, "dashboard.html", map[string]interface{}{
		"auth": auth,
	})
}

func LogoutUser(c echo.Context) error {
	return c.Redirect(http.StatusMovedPermanently, "/")
}

func RegisterView(c echo.Context) error {
	flash := c.(*CustomContext).GetFlash()
	return c.Render(http.StatusOK, "register.html", map[string]interface{}{
		"flash": flash,
	})
}

func RegisterUser(c echo.Context) error {
	result := SaveUser(c)
	fmt.Println(result.LastInsertId())
	c.(*CustomContext).SetFlash("done", "User registred successfully.")
	return c.Redirect(http.StatusMovedPermanently, "/")
}
