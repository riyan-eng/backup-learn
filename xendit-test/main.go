package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/uuid"
)

func main() {
	client := resty.New()
    // me
	// client.SetBasicAuth("xnd_development_n20xxyXWX6JLwrkfZXV6fsQjZrb7tsrNr8DV7Bu7gC3lrXPG0PVj1iONOKt8F", "")
	// fais
    client.SetBasicAuth("xnd_development_sM8r21aNXvA9YzZw2jIg6v4t1Iu4kKPKzMf5MOo3L5e2bX1JSV7KubtSwZKR5i", "")
	client.BaseURL = "https://api.xendit.co"

	cust := createCustomer(client)
	createPlan(client, cust)
}

type custResp struct {
	CustomerId  string `json:"id"`
	ReferenceId string `json:"reference_id"`
}

type planResp struct {
	Actions []planActResp `json:"actions"`
}

type planActResp struct {
	Url     string `json:"url"`
	Action  string `json:"action"`
	Method  string `json:"method"`
	UrlType string `json:"url_type"`
}

func createCustomer(c *resty.Client) *custResp {
	res := new(custResp)
	id := uuid.NewString()
	resp, err := c.R().
		SetResult(res).
		SetBody(map[string]interface{}{
			"reference_id":  id,
			"mobile_number": "+628888811111",
			"email":         "merchant@xendit.co",
			"type":          "INDIVIDUAL",
			"individual_detail": map[string]interface{}{
				"given_names": "Mlon Eusk",
			},
		}).
		Post("/customers")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp.StatusCode())
	return res
}

func createPlan(c *resty.Client, cust *custResp) {
	// now := time.Now()
	// startPay := now.UTC().Format("2006-01-02T15:04:05.000Z")

	res := new(planResp)
	timestamp := time.Now().Unix()
	schedule := map[string]interface{}{
		"reference_id":                 fmt.Sprintf("test-%v", timestamp),
		"interval":                     "MONTH",
		"interval_count":               3,
		"retry_interval":               "DAY",
		"retry_interval_count":         1,
		"total_retry":                  3,
		"failed_attempt_notifications": []int{1, 3},
	}
	// if now.Day() > 28 {
	// 	delete(schedule, "anchor_date")
	// }
	resp, err := c.R().
		SetResult(res).
		SetBody(map[string]interface{}{
			"reference_id":     fmt.Sprintf("ref-%v", timestamp),
			"customer_id":      cust.CustomerId,
			"recurring_action": "PAYMENT",
			"currency":         "IDR",
			"amount":           500000,
			"schedule":         schedule,
			"notification_config": map[string]interface{}{
				"locale":              "en",
				"recurring_created":   []string{"EMAIL"},
				"recurring_succeeded": []string{"EMAIL"},
				"recurring_failed":    []string{"EMAIL"},
			},
			"failed_cycle_action":   "STOP",
			"immediate_action_type": "FULL_AMOUNT",
			"description":           "Xendit Recurring Test",
		}).
		Post("/recurring/plans")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(resp.StatusCode())
	fmt.Println(res)
}
