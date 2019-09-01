package transaction

import (
	"gihub.com/moeen/salamantex_back/controllers"
	"gihub.com/moeen/salamantex_back/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Info(context *gin.Context) {
	var errorResponse controllers.ErrorResponse

	userID, _ := context.Get("user_id")
	txIDString := context.Param("id")

	txID, err := strconv.Atoi(txIDString)
	if err != nil {
		errorResponse.Errors = append(errorResponse.Errors, "Invalid TX id")
		context.JSON(400, errorResponse)
		return
	}

	var tx models.Transaction

	if err := models.GetDB().Where(models.Transaction{ID: uint(txID), SourceUser: userID.(uint)}).Find(&tx).Error; err != nil {
		errorResponse.Errors = append(errorResponse.Errors, "TX not found")
		context.JSON(404, errorResponse)
		return
	}

	context.JSON(200, controllers.Response{
		"state": tx.State,
	})
}
