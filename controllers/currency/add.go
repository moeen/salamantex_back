package currency

import (
	"database/sql"
	"gihub.com/moeen/salamantex_back/constants"
	"gihub.com/moeen/salamantex_back/controllers"
	"gihub.com/moeen/salamantex_back/models"
	"github.com/gin-gonic/gin"
)

type AddCurrencyModel struct {
	Address string                    `json:"address"`
	Type    constants.TransactionType `json:"type"`
}

func AddCurrency(context *gin.Context) {
	var jsonData AddCurrencyModel
	var errorResponse controllers.ErrorResponse

	if err := context.ShouldBindJSON(&jsonData); err != nil {
		errorResponse.Errors = append(errorResponse.Errors, "Body is not valid")
		context.JSON(400, errorResponse)
		return
	}

	userID, _ := context.Get("user_id")

	var user models.User
	models.GetDB().First(&user, userID)

	if jsonData.Type == constants.Bitcoin {
		if err := controllers.GetValidate().Var(jsonData.Address, "required,btc_addr"); err != nil {
			errorResponse.Errors = append(errorResponse.Errors, "Invalid Bitcoin address")
			context.JSON(400, errorResponse)
			return
		}

		if user.BitcoinAddress.String != "" {
			errorResponse.Errors = append(errorResponse.Errors, "Bitcoin address already exists")
			context.JSON(400, errorResponse)
			return
		}

		user.BitcoinAddress = sql.NullString{
			String: jsonData.Address,
			Valid:  true,
		}
		// 10 BTC to test
		user.BitcoinBalance = 10
		if err := models.GetDB().Save(&user).Error; err != nil {
			errorResponse.Errors = append(errorResponse.Errors, "Duplicate address")
			context.JSON(400, errorResponse)
			return
		}

		context.JSON(200, controllers.Response{
			"type":    jsonData.Type,
			"address": user.BitcoinAddress.String,
		})
		return
	} else if jsonData.Type == constants.Ethereum {
		if err := controllers.GetValidate().Var(jsonData.Address, "required,eth_addr"); err != nil {
			errorResponse.Errors = append(errorResponse.Errors, "Invalid Ethereum address")
			context.JSON(400, errorResponse)
			return
		}

		if user.EthereumAddress.Valid {
			errorResponse.Errors = append(errorResponse.Errors, "Ethereum address already exists")
			context.JSON(400, errorResponse)
			return
		}

		user.EthereumAddress = sql.NullString{
			String: jsonData.Address,
			Valid:  true,
		}
		// // 10 ETH to test
		user.EthereumBalance = 10
		if err := models.GetDB().Save(&user).Error; err != nil {
			errorResponse.Errors = append(errorResponse.Errors, "Duplicate address")
			context.JSON(400, errorResponse)
			return
		}

		context.JSON(200, controllers.Response{
			"type":    jsonData.Type,
			"address": user.EthereumAddress.String,
		})
		return
	} else {
		errorResponse.Errors = append(errorResponse.Errors, "Invalid currency type")
		context.JSON(400, errorResponse)
		return
	}
}
