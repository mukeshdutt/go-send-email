
## Dynamic Email Sender using Golang

### Purponse
Send the information to the email and the information should be dynamic. Hence, we have used mongodb as a database.

### Tech Stacks
- Golang
- MongoDB

<br>

### Application Flow
---
Fetch all the ready to send email information from mongodb and execute email method to send.

<br>

<!-- ### Models
---
```go
type Model struct {
	ToAddress []string `json:"toAddress"`
	CcAddress []string `json:"ccAddress"`
	Subject   string   `json:"subject"`
	EmailBody string   `json:"emailBody"`
	IsSent    bool     `json:"isSent"`
}
``` -->