package messages

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
)

func SendMail(msg string, recipient []string) {
	// Message.
	message := []byte("This is a test email message.")

	// Authentication.
	auth := smtp.PlainAuth("", From_mail, Mail_password, SMTP_Host)

	fmt.Println(auth)
	// Sending email.
	if err := smtp.SendMail(fmt.Sprintf("%s:%d", SMTP_Host, 587), auth, From_mail, recipient, message); err != nil {
		log.Printf("Error sending mail %v", err)
		return
	}

	log.Println("Email Sent Successfully!")
}

func SendMails(msg string, recipients []string) {
	for _, v := range recipients {
		if err := SendSSLMail(msg, v); err != nil {
			log.Printf("Error sending mail %v", err)
		}
	}
}

func SendSSLMail(msg string, recipient string) error {
	from := mail.Address{Name: "", Address: From_mail}
	to := mail.Address{Name: "", Address: recipient}

	subj := "Test Mail"
	body := msg

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subj

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Authentication.
	auth := smtp.PlainAuth("", From_mail, Mail_password, SMTP_Host)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         SMTP_Host,
	}

	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", SMTP_Host, 465), tlsconfig)
	if err != nil {
		log.Printf("Error sending mail %v", err)
		return err
	}

	c, err := smtp.NewClient(conn, SMTP_Host)
	if err != nil {
		log.Printf("Error sending mail %v", err)
		return err
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Printf("Error sending mail %v", err)
		return err
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Printf("Error sending mail %v", err)
		return err
	}

	if err = c.Rcpt(to.Address); err != nil {
		log.Printf("Error sending mail %v", err)
		return err
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Printf("Error sending mail %v", err)
		return err
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Printf("Error sending mail %v", err)
		return err
	}

	err = w.Close()
	if err != nil {
		log.Printf("Error sending mail %v", err)
		return err
	}

	if err = c.Quit(); err != nil {
		return err
	}
	return nil
}
