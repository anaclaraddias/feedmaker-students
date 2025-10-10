package handler

import (
	"crypto/sha512"
	"encoding/hex"
	"net/http"

	"github.com/anaclaraddias/feedmaker-students/src/domain/entity"
	"github.com/anaclaraddias/feedmaker-students/src/domain/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateUserHandler struct {
	db *gorm.DB
}

func NewCreateUserHandler(db *gorm.DB) HandlerInterface {
	return &CreateUserHandler{db: db}
}

type CreateStudentRequest struct {
	Name     string          `json:"name" binding:"required"`
	Type     entity.UserType `json:"type" binding:"required"`
	Username string          `json:"username" binding:"required"`
	Password string          `json:"password" binding:"required"`
}

func (handler *CreateUserHandler) Handle(context *gin.Context) {
	context.Writer.Header().Set("Content-Type", "application/json")

	var req CreateStudentRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashPassword := sha512.Sum512([]byte(req.Password))
	password := hex.EncodeToString(hashPassword[:])

	user := &entity.User{
		Name:     req.Name,
		Type:     req.Type,
		Username: req.Username,
		Password: password,
	}

	if err := service.NewCreateUserService(handler.db).Execute(user); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, map[string]interface{}{"id": user.ID, "name": user.Name})
}
