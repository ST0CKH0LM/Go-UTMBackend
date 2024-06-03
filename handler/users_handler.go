package handler

import (
	"log"

	"github.com/gin-gonic/gin"
	"gitlab.com/Std217/test/usecase"
)

type UsersHandler struct {
	usersUsecase usecase.UsersUsecase
}

func NewUsersHandler(usecase usecase.UsersUsecase) *UsersHandler {
	return &UsersHandler{usersUsecase: usecase}
}

func (handler *UsersHandler) GetAllUser(c *gin.Context) {
	products, err := handler.usersUsecase.GetAllUser()
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(200, gin.H{"data": products})
}
