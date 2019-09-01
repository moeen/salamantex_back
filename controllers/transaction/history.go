package transaction

import (
	"gihub.com/moeen/salamantex_back/models"
	"github.com/gin-gonic/gin"
)

func History(context *gin.Context) {
	userID, _ := context.Get("user_id")

	var transactions []models.Transaction

	models.GetDB().Where(&models.Transaction{SourceUser: userID.(uint)}).Find(&transactions)
	context.JSON(200, transactions)
}