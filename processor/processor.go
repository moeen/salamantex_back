package processor

import (
	"fmt"
	"gihub.com/moeen/salamantex_back/constants"
	"gihub.com/moeen/salamantex_back/models"
	"log"
	"strconv"
	"time"
)

func ProcessTransaction(txString string) {
	txID, err := strconv.Atoi(txString)
	if err != nil {
		log.Println(err)
		return
	}

	var tx models.Transaction
	if err := models.GetDB().First(&tx, txID).Error; err != nil {
		log.Println(err)
		return
	}

	var sourceUser models.User
	if err := models.GetDB().First(&sourceUser, tx.SourceUser).Error; err != nil {
		log.Println(err)
		return
	}

	var targetUser models.User
	if err := models.GetDB().First(&targetUser, tx.TargetUser).Error; err != nil {
		log.Println(err)
		return
	}

	if tx.Type == constants.Bitcoin {
		if tx.Amount > sourceUser.BitcoinBalance {
			tx.State = constants.Rejected
		} else {
			tx.State = constants.Approved
			sourceUser.BitcoinBalance = sourceUser.BitcoinBalance - tx.Amount
			targetUser.BitcoinBalance = targetUser.BitcoinBalance + tx.Amount
		}
	} else if tx.Type == constants.Ethereum {
		if tx.Amount > sourceUser.EthereumBalance {
			tx.State = constants.Rejected
		} else {
			tx.State = constants.Approved
			sourceUser.EthereumBalance = sourceUser.EthereumBalance - tx.Amount
			targetUser.EthereumBalance = targetUser.EthereumBalance + tx.Amount
		}
	}
	now := time.Now()
	tx.Processed = &now

	models.GetDB().Save(&tx)
	models.GetDB().Save(&targetUser)
	models.GetDB().Save(&sourceUser)

	log.Println(fmt.Sprintf("TX %v has been proceeded with status %v", tx.ID, tx.State))
}
