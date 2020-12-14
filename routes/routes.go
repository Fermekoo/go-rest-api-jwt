package routes

import (
	"net/http"

	"github.com/rs/cors"

	"github.com/Fermekoo/blog-dandi/controllers"
	"github.com/Fermekoo/blog-dandi/middleware"
	"github.com/labstack/echo"
)

func Init() *echo.Echo {

	e := echo.New()

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"OPTIONS", "GET", "POST", "PUT", "DELETE"},
		AllowedHeaders: []string{"*"},
		Debug:          true,
	})

	e.Use(echo.WrapMiddleware(corsMiddleware.Handler))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello this is echo")
	})

	e.GET("/users", controllers.FetchAllUser, middleware.IsAuthenticated)
	e.GET("/users", controllers.FetchAllUser)
	e.POST("/users", controllers.StoreUser)
	e.PUT("/users", controllers.UpdateUser)
	e.DELETE("/users", controllers.DeleteUser)
	e.POST("/upload", controllers.Upload)

	e.POST("/login", controllers.CheckLogin)

	return e
}
