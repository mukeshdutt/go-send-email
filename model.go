package main

type mailerModel struct {
	ToAddress []string `json:"toAddress"`
	CcAddress []string `json:"ccAddress"`
	Subject   string   `json:"subject"`
	EmailBody string   `json:"emailBody"`
	IsSent    bool     `json:"isSent"`
}
