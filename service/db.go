package service

import (
	"github.com/suavelad/notification_service_go/config"
	"github.com/suavelad/notification_service_go/models"
)

func saveNotificationLog(alertPayload map[string]interface{}) {

	alert := &models.NotificationSentLog{
		FbAdAccountId:          alertPayload["ad_account_id"].(string),
		FacebookAdBalanceLogId: alertPayload["ad_balance_log_id"].(int),
		SentViaEmail:           alertPayload["sent_via_email"].(bool),
		SentViaPhone:           alertPayload["sent_via_phone"].(bool),
		SentViaPush:            alertPayload["sent_via_push"].(bool),
	}

	if err := config.DB.Create(alert).Error; err != nil {
		return err
	}

	return nil

}

func saveAdBalanceLog(balanceLogPayload map[string]interface{}) {

	adLogs := &models.FacebookAdBalanceLog{
		FbAdAccountId: balanceLogPayload["account_id"].(string),
		MainBalance:   balanceLogPayload["main_balance"].(int),
		Currency:      balanceLogPayload["currency"].(string),
		AmountSpent:   balanceLogPayload["amount_spent"].(string),
		SpendCap:      balanceLogPayload["spent_cap"].(string),
		Threshold:     balanceLogPayload["threshold"].(string),
	}

	if err := config.DB.Create(adLogs).Error; err != nil {
		return _, err
	}

	return adLogs, nil

}
