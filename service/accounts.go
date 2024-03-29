package service

import (
	config "command-line-arguments/Users/sunnex/Documents/Assessment/notification_service_go/config/base.go"
	models "command-line-arguments/Users/sunnex/Documents/Assessment/notification_service_go/models/account.go"
	"fmt"

	"github.com/gin-gonic/gin"
)


func GetAllAccountIds() {
	accountIds := []models.AdAccount{}
	config.DB.Where("is_active = ?", true).Where("is_verified = ?", true).Find(&accountIds).Pluck("id", &accountIds)
	return accountIds

}


func GetFbAdUrl(accountID string) {
	resp, err := http.Get(url)
	if err != nil {
		return nil
	}
	defer resp.Body.Close()
	return fmt.Sprintf("%s/%s/%s?fields=amount_spent,balance,spend_cap&access_token=%s|%s", config.baseURL,config.BASE_VERSION, accountID, config.appID, config.appSecret)
}


func GetUserBalance(accountId string) {
	url := GetFbAdUrl(accountId)

	resp, err := http.Get(url)
	if err != nil {
		
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer resp.Body.Close()

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	amountSpent := data["amount_spent"].(float64)
	spendCap := data["spend_cap"].(float64)
	balance := data["balance"].(float64)
	returnedAccountID := data["id"].(string)

	if returnedAccountID == nil || returnedAccountID != accountId {
		fmt.Println("Error: Account ID mismatch")
		return nil
	}
	var mainBalance := (spendCap - amountSpent) / 100

	payload := map[string]interface{}{
		"main_balance": mainBalance,
		"account_id":   returnedAccountID,
		"currency":     "USD",
		"threshold":    config.THRESHOLD_BALANCE,
		"amount_spent": amountSpent,
		"spend_cap":    spendCap,
		"balance":      balance,
	}

	return payload
}

