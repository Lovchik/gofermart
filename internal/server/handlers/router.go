package handlers

import (
	"github.com/gin-gonic/gin"
	"gofermart/internal/server/models"
	"gofermart/internal/server/utils"
	"net/http"
)

func UserRegister(router *gin.RouterGroup, s *Service) {
	router.POST("/refresh", s.Refresh)
	router.POST("/login", s.Login)
	router.POST("/register", s.RegisterUser)
	//router.GET("/orders", AuthMiddleware(), s.GetOrders)
	//router.GET("/balance", AuthMiddleware(), s.GetBalance)
	//router.POST("/withdraw", AuthMiddleware(), s.Withdraw)
	//router.GET("/withdrawals", AuthMiddleware(), s.GetWithdrawals)
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var response models.Response
		token := c.GetHeader("Authorization")
		if token == "" && !utils.IsValidToken(token, "access") {
			c.JSON(http.StatusUnauthorized, response.ErrorResponse("Неверный токен"))
			c.Abort()
			return
		}
		userID, err := utils.GetUserId(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, response.ErrorResponse("Ошибка получения пользователя"))
			return
		}
		c.Set("user_id", userID)
		c.Next()
	}
}
