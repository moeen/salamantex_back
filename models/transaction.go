package models

import (
	"gihub.com/moeen/salamantex_back/constants"
	"time"
)

type Transaction struct {
	ID         uint                       `gorm:"primary_key" json:"id"`
	Amount     float64                    `json:"amount"`
	Type       constants.TransactionType  `json:"type"`
	SourceUser uint                       `json:"source_user"`
	TargetUser uint                       `json:"target_user"`
	Created    time.Time                  `json:"created"`
	Processed  *time.Time                 `json:"processed"`
	State      constants.TransactionState `json:"state"`
}
