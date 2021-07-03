package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project2/api/routes"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project2/helper"
)

func DefineApiRoutes(e *echo.Echo) {
	handlers := []helper.Handler{
		routes.NewsRoutes{},
	}

	var routes []helper.Route
	for _, handler := range handlers {
		// log.Println("WE'RE HERE routes: ", handler)
		routes = append(routes, handler.Route()...)
	}
	api := e.Group("/api")
	for _, route := range routes {
		switch route.Method {
		case echo.POST:
			{
				api.POST(route.Path, route.Handler, route.Middleware...)
			}
		case echo.GET:
			{
				api.GET(route.Path, route.Handler, route.Middleware...)
			}
		case echo.PUT:
			{
				api.PUT(route.Path, route.Handler, route.Middleware...)
			}
		case echo.DELETE:
			{
				api.DELETE(route.Path, route.Handler, route.Middleware...)
			}
		case echo.PATCH:
			{
				api.PATCH(route.Path, route.Handler, route.Middleware...)
			}
		}

	}
}
