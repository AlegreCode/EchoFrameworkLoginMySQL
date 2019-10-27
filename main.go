package main

import (
	"fmt"
	"html/template"
	"log"

	. "github.com/alegrecode/echo/LoginMySQL/middlewares"

	. "github.com/alegrecode/echo/LoginMySQL/db"

	"github.com/alegrecode/echo/LoginMySQL/controllers/users"

	. "github.com/alegrecode/echo/LoginMySQL/helpers"

	"github.com/Masterminds/sprig"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	e := echo.New()
	e.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
		Getter: middleware.MethodFromForm("_method"),
	}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Static("assets"))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("alegrecode"))))

	templates := make(map[string]*template.Template)
	templates["login.html"] = template.Must(template.New("base").Funcs(sprig.FuncMap()).ParseFiles("views/login.html", "views/base.html", "views/navbar.partial.html", "views/alert.partial.html"))
	templates["register.html"] = template.Must(template.New("base").Funcs(sprig.FuncMap()).ParseFiles("views/register.html", "views/base.html", "views/navbar.partial.html", "views/alert.partial.html"))
	templates["dashboard.html"] = template.Must(template.New("base").Funcs(sprig.FuncMap()).ParseFiles("views/dashboard.html", "views/base.html", "views/navbar.partial.html", "views/alert.partial.html"))

	e.Renderer = &TemplateRegistry{
		Templates: templates,
	}

	DB, _ := GetConexion()
	err = DB.Ping()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("DB is connected.")
	}

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c}
			return next(cc)
		}
	})

	e.GET("/", users.LoginView, IsNotLogged)

	e.POST("/", users.LoginUser, ValidateLogin)

	e.GET("/dashboard", users.DashboardView, IsLogged)

	e.DELETE("/logout", users.LogoutUser, LogoutMiddleware)

	e.GET("/register", users.RegisterView)

	e.POST("/register", users.RegisterUser, ValidateRegister)

	e.Logger.Fatal(e.Start(":5000"))
}
