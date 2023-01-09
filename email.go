package main

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
)

// created constans for server details, but you can configure in config file
// when you create production application
const (
	HOST = "smtp.gmail.com"
	PORT = "587"
)

// saved the credentials into the database
var (
	FROM   = os.Getenv("MAIL")
	PASSWD = os.Getenv("PASSWD")
)

// created a structure to feed the required details to send the email
type Mail struct {
	To      []string
	Cc      []string
	Bcc     []string
	Subject string
	Body    string
}

// prepared the email body, where golang gives options to create a template
// with To, CC, From, Bcc, Attachment and body then put prepared string into
// then send data into body parameters
func buildBody(mail Mail) string {

	var msg string
	msg = fmt.Sprintf("From: %s\r\n", FROM)

	if len(mail.To) > 0 {
		msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	}
	if len(mail.Cc) > 0 {
		msg += fmt.Sprintf("Cc: %s\r\n", strings.Join(mail.Cc, ";"))
	}
	if len(mail.Bcc) > 0 {
		msg += fmt.Sprintf("Bcc: %s\r\n", strings.Join(mail.Bcc, ";"))
	}
	if mail.Subject != "" {
		msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	}
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)
	return msg
}

// Finally, called the smtp function to send the email
func sendEmail(mail Mail) bool {

	auth := smtp.PlainAuth("", FROM, PASSWD, HOST)
	addr := fmt.Sprintf("%s:%s", HOST, PORT)
	msg := buildBody(mail)
	err := smtp.SendMail(addr, auth, FROM, mail.To, []byte(msg))

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}
