package transaction

import (
	"gihub.com/moeen/salamantex_back/config"
	"gihub.com/moeen/salamantex_back/constants"
	"gihub.com/moeen/salamantex_back/controllers"
	"gihub.com/moeen/salamantex_back/models"
	"gihub.com/moeen/salamantex_back/redis"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

type SendModel struct {
	To     string                    `json:"to"`
	Type   constants.TransactionType `json:"type"`
	Amount float64                   `json:"amount"`
}

func Send(context *gin.Context) {
	var jsonData SendModel
	var errorResponse controllers.ErrorResponse

	if err := context.ShouldBindJSON(&jsonData); err != nil {
		errorResponse.Errors = append(errorResponse.Errors, "Body is not valid")
		context.JSON(400, errorResponse)
		return
	}

	userID, _ := context.Get("user_id")

	var sourceUser models.User
	models.GetDB().First(&sourceUser, userID)

	var targetUser models.User
	if err := models.GetDB().Where(&models.User{Email: jsonData.To}).Find(&targetUser).Error; err != nil {
		errorResponse.Errors = append(errorResponse.Errors, "User not found")
		context.JSON(404, errorResponse)
		return
	}

	if jsonData.Type == constants.Bitcoin {
		if !targetUser.BitcoinAddress.Valid || !sourceUser.BitcoinAddress.Valid {
			errorResponse.Errors = append(errorResponse.Errors, "Sender or receptor does not have Bitcoin address")
			context.JSON(404, errorResponse)
			return
		}
		if jsonData.Amount > sourceUser.BitcoinBalance {
			errorResponse.Errors = append(errorResponse.Errors, "You don't have enough balance")
			context.JSON(404, errorResponse)
			return
		}
	} else if jsonData.Type == constants.Ethereum {
		if !targetUser.EthereumAddress.Valid || !sourceUser.EthereumAddress.Valid {
			errorResponse.Errors = append(errorResponse.Errors, "Sender or receptor does not have Ethereum address")
			context.JSON(404, errorResponse)
			return
		}
		if jsonData.Amount > sourceUser.EthereumBalance {
			errorResponse.Errors = append(errorResponse.Errors, "You don't have enough balance")
			context.JSON(404, errorResponse)
			return
		}
	} else {
		errorResponse.Errors = append(errorResponse.Errors, "Invalid currency type")
		context.JSON(400, errorResponse)
		return
	}

	tx := models.Transaction{
		Amount:     jsonData.Amount,
		Created:    time.Now(),
		Type:       jsonData.Type,
		SourceUser: sourceUser.ID,
		TargetUser: targetUser.ID,
		State:      constants.Pending,
	}

	if err := models.GetDB().Create(&tx).Error; err != nil {
		log.Println(err)
		errorResponse.Errors = append(errorResponse.Errors, "Something bad happened, please try again later.")
		context.JSON(500, errorResponse)
		return
	}
	_, err := redis.GetRedis().LPush(config.GetConfig().Redis.TxQueue, tx.ID).Result()
	if err != nil {
		log.Println(err)
		errorResponse.Errors = append(errorResponse.Errors, "Something bad happened, please try again later.")
		context.JSON(500, errorResponse)
		return
	}
	context.JSON(200, controllers.Response{
		"id":          tx.ID,
		"type":        tx.Type,
		"amount":      tx.Amount,
		"source_user": sourceUser.Email,
		"target_user": targetUser.Email,
		"state":       tx.State,
	})
	return
}
