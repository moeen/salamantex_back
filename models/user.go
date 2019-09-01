package models

import (
	"database/sql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name            string         `gorm:"type:varchar(512)"`
	Description     string         `gorm:"type:varchar(1000)"`
	Email           string         `gorm:"type:varchar(1000);unique"`
	Password        string         `gorm:"type:varchar(100)"`
	BitcoinAddress  sql.NullString `gorm:"type:varchar(34);unique"`
	BitcoinBalance  float64
	EthereumAddress sql.NullString `gorm:"type:varchar(42);unique"`
	EthereumBalance float64
	MaxTxAmount     uint
}
