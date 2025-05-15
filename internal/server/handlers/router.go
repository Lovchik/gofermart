package handlers

import (
	"github.com/gin-gonic/gin"
)

func AuthenticationRegister(router *gin.RouterGroup, s *Service) {
	router.POST("/refresh", s.Refresh)
	router.POST("/login", s.Login)
}
