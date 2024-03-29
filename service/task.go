package service

import (
	"github.com/suavelad/notification_service_go/config"
)

func sendBalanceAlert() {
	accountIDs := GetAllAccountIds()

	if len(accountIDs) == 0 || accountIDs == nil {
		return
	}

	for _, accountID := range accountIDs {
		accountId := accountID.FbAccountId
		balancePayload := GetUserBalance(accountId)

		if balancePayload == nil || balancePayload["mainBalance"] == nil {
			continue
		}
		balance := balancePayload["mainBalance"].(float64)

		balanceLog, err := saveAdBalanceLog(balancePayload)
		if err != nil {
			return
		}

		if balance > config.THRESHOLD_BALANCE {
			continue
		} else {
			SendNotificationAlert(balanceLog, accountID, balance)
		}

	}

	return "Alert sent successfully"

}
