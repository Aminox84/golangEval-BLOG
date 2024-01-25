package web

import (
	"blog/internal/database"
	"blog/internal/repositories"
	"blog/internal/routes"
	"blog/internal/services"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer(router *gin.Engine) *Server {
	return &Server{router}
}

func (s *Server) Run(addr string) error {
	return s.router.Run(addr)
}

func (s *Server) RegisterRoutes() {
	database.Init()

	userService := services.NewUserService(repositories.NewUserRepository(database.DB))
	articleService := services.NewArticleService(repositories.NewArticleRepository(database.DB))

	userRoutes := routes.NewUserRoutes(userService)
	userRoutes.RegisterRoutes(s.router)

	articleRoutes := routes.NewArticleRoutes(articleService)
	articleRoutes.RegisterRoutes(s.router)

}
