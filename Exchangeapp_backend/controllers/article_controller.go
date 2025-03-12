package controllers

import (
	"exchangeapp/global"
	"exchangeapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var cache = "articles"

func CreateArticle(ctx *gin.Context) {
	var article models.Article

	if err := ctx.ShouldBindJSON(&article); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := global.Db.AutoMigrate(&article); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}
