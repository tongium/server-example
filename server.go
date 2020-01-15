package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func main() {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", func(c echo.Context) error {
		fmt.Println(`{"severity":"panic","timestamp":"2020-01-16T13:45:41.329Z","user_id":"1","uri":"/post/admin/cross-overrule-data","method":"POST","ip":"::1","message":{"portDate":"2021-01-08"}}`)
		fmt.Println(`{"severity":"info","timestamp":"2020-01-16T13:45:41.329Z","user_id":"2","uri":"/post/admin/cross-overrule-data","method":"POST","ip":"::1","message":{"portDate":"2021-01-08"}}`)
		fmt.Println(`{"severity":"debug","timestamp":"2020-01-16T13:45:41.329Z","user_id":"3","uri":"/post/admin/cross-overrule-data","method":"POST","ip":"::1","message":{"portDate":"2021-01-08"}}`)
		fmt.Println(`{"severity":"warn","timestamp":"2020-01-16T13:45:41.329Z","user_id":"4","uri":"/post/admin/cross-overrule-data","method":"POST","ip":"::1","message":{"portDate":"2021-01-08"}}`)
		return c.String(http.StatusOK, "Hello world")
	})

	// Start server
	e.Logger.Panic(e.Start(":" + os.Getenv("PORT")))
}
