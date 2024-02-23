package utils

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"strings"
)

type EmailTypes struct {
	// ID        int       `json:"id"`
	toEmail      []string
	ccEmail      []string
	subjectEmail string
	messageEmail string
}

// const CONFIG_SMTP_HOST = "smtp.gmail.com"
const CONFIG_SMTP_PORT = 587

// const CONFIG_SENDER_NAME = "Name"
// const CONFIG_AUTH_EMAIL = "XXXXXXXXXX@gmail.com"
// const CONFIG_AUTH_PASSWORD = "XXXXXXXXXX"

func Email(toEmail []string, ccEmail []string, subjectEmail string, messageEmail string) {
	to := toEmail
	// []string{"recipient1@gmail.com", "emaillain@gmail.com"}
	cc := ccEmail
	// []string{"tralalala@gmail.com"}
	subject := "Test mail"
	message := "Hello"

	err := sendMail(to, cc, subject, message)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println("Mail sent!")
}

func sendMail(to []string, cc []string, subject, message string) error {
	body := "From: " + os.Getenv("CONFIG_SENDER_NAME") + "\n" +
		"To: " + strings.Join(to, ",") + "\n" +
		"Cc: " + strings.Join(cc, ",") + "\n" +
		"Subject: " + subject + "\n\n" +
		message

	auth := smtp.PlainAuth("", os.Getenv("CONFIG_AUTH_EMAIL"), os.Getenv("CONFIG_AUTH_PASSWORD"), os.Getenv("CONFIG_SMTP_HOST"))
	smtpAddr := fmt.Sprintf("%s:%d", os.Getenv("CONFIG_SMTP_HOST"), CONFIG_SMTP_PORT)

	err := smtp.SendMail(smtpAddr, auth, os.Getenv("CONFIG_AUTH_EMAIL"), append(to, cc...), []byte(body))
	if err != nil {
		return err
	}

	return nil
}
