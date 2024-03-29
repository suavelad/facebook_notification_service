package models

import (
	"time"

	"gorm.io/gorm"
)

type AdAccount struct {
	gorm.Model

	Id               int       `json:"id" gorm:"primary_key"`
	Name             string    `json:"name"`
	Email            string    `json:"email" gorm:"unique"`
	BusinessName     string    `json:"business_name"`
	FbAccountId      string    `json:"fb_account_id" gorm:"unique"`
	AdAccountBalance float64   `json:"ad_account_balance"`
	IsActive         bool      `json:"is_active"`
	IsDeleted        bool      `json:"is_deleted"`
	IsVerified       bool      `json:"is_verified"`
	CreatedDate      time.Time `gorm:"autoCreateTime"`
	ModifiedDate     time.Time `json:"modified_date"`
	DeletedDate      time.Time `json:"deleted_date"`
}

type FacebookAdBalanceLog struct {
	gorm.Model

	Id            int     `json:"id" gorm:"primary_key"`
	FbAdAccountId int     `json:"fb_ad_account_id"`
	AmountSpent   float64 `json:"amount_spent"`
	MainBalance   float64 `json:"main_balance"`
	Currency      string  `json:"currency"`
	Threshold     float64 `json:"threshold"`
	SpendCap      float64 `json:"spend_cap"`
}

type NotificationSentLog struct {
	gorm.Model

	Id                     int       `json:"id" gorm:"primary_key"`
	FbAdAccountId          int       `json:"fb_ad_account_id"`
	CreatedDate            time.Time `gorm:"autoCreateTime"`
	SentViaEmail           bool      `json:"sent_via_email"`
	SentViaPhone           bool      `json:"sent_via_phone"`
	SentViaPush            bool      `json:"sent_via_push"`
	FacebookAdBalanceLogId int       `json:"facebook_ad_balance_log_id"`
}
