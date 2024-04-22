package route

import (
	"atmail/internal/http/handler"
	"atmail/internal/http/middleware"

	"github.com/gin-gonic/gin"
)

type UserRoute struct {
	handler handler.UserHandler
}

func NewUserRoute(userHandler handler.UserHandler) *UserRoute {
	return &UserRoute{
		handler: userHandler,
	}
}

func (u *UserRoute) Setup(router *gin.RouterGroup) {
	router.Use(middleware.AuthHandler)
	router.GET("users", u.handler.GetAll)
	router.GET("users/:id", u.handler.Get)
	router.POST("users", u.handler.Create)
	router.PUT("users/:id", u.handler.Update)
	router.DELETE("users/:id", u.handler.Delete)
}
