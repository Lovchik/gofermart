package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gofermart/internal/server/models"
	"gofermart/internal/server/storage"
	"gofermart/internal/server/utils"
	"net/http"
)

type Service struct {
	WebServer *gin.Engine
	Store     storage.Storage
}

func (s *Service) Refresh(c *gin.Context) {
	var response models.Response
	header := c.GetHeader("Authorization")
	if !utils.IsValidToken(header, "refresh") {
		log.Info("Неверный токен")
		c.JSON(http.StatusUnauthorized, response.ErrorResponse("Неверный токен"))
		return
	}
	tokens, err := s.Store.Refresh(header)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
	}
	c.JSON(http.StatusOK, response.NewWithMessage(tokens, ""))
	return
}

func (s *Service) Login(c *gin.Context) {
	var response models.Response
	var credentials models.LoginRequest
	err := c.ShouldBind(&credentials)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Неверные параметры тела запроса"))
		return
	}
	err = utils.Validate().Struct(credentials)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse("Должны быть заполнены логин и пароль"))
		return
	}
	tokens, err := s.Store.Login(credentials)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.NewWithMessage(tokens, "Пользователь успешно авторизирован"))
	return
}
