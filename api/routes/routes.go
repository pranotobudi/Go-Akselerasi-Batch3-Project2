package routes

import (
	"log"

	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project2/api/handler"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project2/api/repository"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project2/api/service"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project2/auth"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project2/database"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project2/helper"
	"github.com/pranotobudi/Go-Akselerasi-Batch3-Project2/middleware"

	"github.com/labstack/echo"
)

type NewsRoutes struct{}

func (r NewsRoutes) Route() []helper.Route {
	log.Println("INSIDE NewsRoutes.Route")
	db := database.GetDBInstance()
	repository.InitDBTable(db)
	repository.DBSeed(db)
	// db.AutoMigrate(User{}, Role{}, Permission{}, RolePermission{}, movie.Genre{}, movie.Movie{}, movie.GenreMovie{}, movie.MovieReview{})
	repo := repository.NewRepository(db)
	newsService := service.NewServices(repo)
	authService := auth.NewAuthService()
	newsHandler := handler.NewHandler(newsService, authService)

	return []helper.Route{
		{
			Method:  echo.POST,
			Path:    "/register/authors",
			Handler: newsHandler.AuthorRegistrationSendEmail,
		},
		{
			Method:  echo.POST,
			Path:    "/register/readers",
			Handler: newsHandler.ReaderRegistrationSendEmail,
		},
		{
			Method:  echo.POST,
			Path:    "/register/admins",
			Handler: newsHandler.AdminRegistrationSendEmail,
		},
		{
			Method:  echo.GET,
			Path:    "/register/confirmation",
			Handler: newsHandler.RegisterConfirmation,
		},
		{
			Method:  echo.POST,
			Path:    "/login",
			Handler: newsHandler.UserLogin,
		},

		// CRUD NEWS

		{
			Method:  echo.POST,
			Path:    "/news",
			Handler: newsHandler.AddNews,
			// Middleware: []echo.MiddlewareFunc{
			// 	middleware.JwtMiddleWare(),
			// 	middleware.RoleAccessMiddleware("admin"),
			// },
		},
		{
			Method:  echo.GET,
			Path:    "/news",
			Handler: newsHandler.GetAllNews,
			// Middleware: []echo.MiddlewareFunc{
			// 	middleware.JwtMiddleWare(),
			// 	middleware.RoleAccessMiddleware("admin"),
			// },
		},
		{
			Method:  echo.GET,
			Path:    "/news/:id",
			Handler: newsHandler.GetNews,
			// Middleware: []echo.MiddlewareFunc{
			// 	middleware.JwtMiddleWare(),
			// 	middleware.RoleAccessMiddleware("admin"),
			// },
		},
		{
			Method:  echo.PUT,
			Path:    "/news/:id",
			Handler: newsHandler.UpdateNews,
			// Middleware: []echo.MiddlewareFunc{
			// 	middleware.JwtMiddleWare(),
			// 	middleware.RoleAccessMiddleware("admin"),
			// },
		},
		{
			Method:  echo.DELETE,
			Path:    "/news/:id",
			Handler: newsHandler.DeleteNews,
			// Middleware: []echo.MiddlewareFunc{
			// 	middleware.JwtMiddleWare(),
			// 	middleware.RoleAccessMiddleware("admin"),
			// },
		},

		{
			Method:  echo.GET,
			Path:    "/news/category/:id",
			Handler: newsHandler.GetNewsByCategoryID,
			// Middleware: []echo.MiddlewareFunc{
			// 	middleware.JwtMiddleWare(),
			// 	middleware.RoleAccessMiddleware("admin"),
			// },
		},
		{
			Method:  echo.GET,
			Path:    "/news/trending",
			Handler: newsHandler.GetAllTrendingNews,
			// Middleware: []echo.MiddlewareFunc{
			// 	middleware.JwtMiddleWare(),
			// 	middleware.RoleAccessMiddleware("admin", "reader"),
			// },
		},
		{
			Method:  echo.GET,
			Path:    "/news/highlight",
			Handler: newsHandler.GetAllHighlightNews,
			Middleware: []echo.MiddlewareFunc{
				middleware.JwtMiddleWare(),
				// middleware.RoleAccessMiddleware("author"),
			},
		},

		// CRUD AUTHOR
		{
			Method:  echo.POST,
			Path:    "/author",
			Handler: newsHandler.AddAuthor,
			// Middleware: []echo.MiddlewareFunc{
			// 	middleware.JwtMiddleWare(),
			// 	middleware.RoleAccessMiddleware("admin"),
			// },
		},
		{
			Method:  echo.GET,
			Path:    "/author/:id",
			Handler: newsHandler.GetAuthor,
			// Middleware: []echo.MiddlewareFunc{
			// 	middleware.JwtMiddleWare(),
			// 	middleware.RoleAccessMiddleware("admin"),
			// },
		},
		{
			Method:  echo.PUT,
			Path:    "/author/:id",
			Handler: newsHandler.UpdateAuthor,
			// Middleware: []echo.MiddlewareFunc{
			// 	middleware.JwtMiddleWare(),
			// 	middleware.RoleAccessMiddleware("admin"),
			// },
		},
		{
			Method:  echo.DELETE,
			Path:    "/author/:id",
			Handler: newsHandler.DeleteAuthor,
			// Middleware: []echo.MiddlewareFunc{
			// 	middleware.JwtMiddleWare(),
			// 	middleware.RoleAccessMiddleware("admin"),
			// },
		},

		// CRUD READER
		{
			Method:  echo.POST,
			Path:    "/reader",
			Handler: newsHandler.AddReader,
			// Middleware: []echo.MiddlewareFunc{
			// 	middleware.JwtMiddleWare(),
			// 	middleware.RoleAccessMiddleware("admin"),
			// },
		},
		{
			Method:  echo.GET,
			Path:    "/reader/:id",
			Handler: newsHandler.GetReader,
			// Middleware: []echo.MiddlewareFunc{
			// 	middleware.JwtMiddleWare(),
			// 	middleware.RoleAccessMiddleware("admin"),
			// },
		},
		{
			Method:  echo.PUT,
			Path:    "/reader/:id",
			Handler: newsHandler.UpdateReader,
			// Middleware: []echo.MiddlewareFunc{
			// 	middleware.JwtMiddleWare(),
			// 	middleware.RoleAccessMiddleware("admin"),
			// },
		},
		{
			Method:  echo.DELETE,
			Path:    "/reader/:id",
			Handler: newsHandler.DeleteReader,
			// Middleware: []echo.MiddlewareFunc{
			// 	middleware.JwtMiddleWare(),
			// 	middleware.RoleAccessMiddleware("admin"),
			// },
		},

		{
			Method: echo.POST,
			Path:   "/news/share",
			// Handler: newsHandler.AddNewsShare,
			Middleware: []echo.MiddlewareFunc{
				middleware.JwtMiddleWare(),
				middleware.RoleAccessMiddleware("author"),
			},
		},
		{
			Method: echo.POST,
			Path:   "/news/comment",
			// Handler: userHandler.GetAllUsers,
			Middleware: []echo.MiddlewareFunc{
				middleware.JwtMiddleWare(),
				middleware.RoleAccessMiddleware("admin"),
			},
		},
		{
			Method:  echo.GET,
			Path:    "/news/statistic",
			Handler: newsHandler.GetStatistic,
			// Middleware: []echo.MiddlewareFunc{
			// 	middleware.JwtMiddleWare(),
			// 	middleware.RoleAccessMiddleware("admin"),
			// },
		},
		{
			Method: echo.PUT,
			Path:   "/users/password",
			// Handler: userHandler.GetAllUsers,
			Middleware: []echo.MiddlewareFunc{
				middleware.JwtMiddleWare(),
				middleware.RoleAccessMiddleware("admin"),
			},
		},
		{
			Method: echo.PUT,
			Path:   "/users/profile",
			// Handler: userHandler.GetAllUsers,
			Middleware: []echo.MiddlewareFunc{
				middleware.JwtMiddleWare(),
				middleware.RoleAccessMiddleware("admin"),
			},
		},
	}
}
