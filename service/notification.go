package service

import (
	"fmt"
	"log"

	"github.com/suavelad/notification_service_go/config"
	"github.com/suavelad/notification_service_go/jobs"
	"github.com/suavelad/notification_service_go/models"
)

func SendNotificationAlert(balanceLog models.FacebookAdBalanceLog, accountObject models.AdAccount, balance float64) {

	var message string
	var sentViaEmail bool

	if balance == config.THRESHOLD_BALANCE {
		log.Println("Balance for account id %s is exactly threshold. Sending alert", accountObject.FbAccountID)
		message = fmt.Sprintf("Your account balance is %d and it is exactly the threshold balance set. Please top up your account", balance)
	} else if balance < config.THRESHOLD_BALANCE {
		log.Println("Balance for account id %s is below threshold. Sending alert", accountObject.FbAccountID)
		message = fmt.Sprintf("Your account balance is %d and it is below the threshold balance set. Please top up your account", balance)
	} else {
		log.Println("Balance for account id %s is above threshold. No alert needed", accountObject.FbAccountID)
		return "Notification sent successfully", nil
	}

	sentViaEmail = true

	alertPayload := map[string]interface{}{
		"ad_account_id":     accountObject.FbAccountID,
		"ad_balance_log_id": balanceLog.ID,
		"sent_via_email":    sentViaEmail,
		"sent_via_phone":    false,
		"sent_via_push":     false,
	}

	err := saveNotificationLog(alertPayload)
	if err != nil {
		return
	}


	//  Send Email
	fmt.Println("Sent to Email Send Job")
	task := jobs.Task{
		Function: SendEmailTask,
		Input:    []interface{}{accountObject.Email, "", "Balance Alert", message},
		Result:   make(chan interface{}),
		Error:    make(chan error),
	}

	// Run the task in the background using a goroutine
	go jobs.ExecuteTask(task)

	return

}
