package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-swagger/go-swagger/examples/generated/models"
	"gorm.io/gorm/utils"
)

func Register(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPwd, err := utils.HashPassword(user.Password)
}
