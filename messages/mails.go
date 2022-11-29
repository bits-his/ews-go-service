package messages

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"

	"gopkg.in/gomail.v2"
)

func SendMails(subject, msg string, recipients []string) {
	mailwg.Add(len(recipients))
	for _, v := range recipients {
		go func(recipient string) {
			defer mailwg.Done()
			SendMail(subject, msg, recipient)
			log.Printf("Mail sent to %s, successfully", recipient)
		}(v)
	}
	mailwg.Wait()
}

func SendSSLMail(subject, msg string, recipient string) {
	to := mail.Address{Name: "", Address: recipient}

	Mail_subject = subject
	Mail_body = msg

	container := NewContainer()
	container.m.Lock()
	// Setup headers
	//headers = make(map[string]string)
	container.Headers["From"] = from.String()
	container.Headers["To"] = to.String()
	container.Headers["Subject"] = Mail_subject
	defer container.m.Unlock()

	// Setup message
	message := ""
	for k, v := range container.Headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + Mail_body

	conn, err := tls.Dial("tcp", fmt.Sprintf("%s:%d", SMTP_Host, 465), tlsconfig)
	if err != nil {
		log.Printf("Error sending mail %v", err)
		return
	}

	c, err := smtp.NewClient(conn, SMTP_Host)
	if err != nil {
		log.Printf("Error sending mail %v", err)
		return
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Printf("Error sending mail %v", err)
		return
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Printf("Error sending mail %v", err)
		return
	}

	if err = c.Rcpt(to.Address); err != nil {
		log.Printf("Error sending mail %v", err)
		return
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Printf("Error sending mail %v", err)
		return
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Printf("Error sending mail %v", err)
		return
	}

	err = w.Close()
	if err != nil {
		log.Printf("Error sending mail %v", err)
		return
	}

	if err = c.Quit(); err != nil {
		return
	}
}

func SendMail(subject, msg string, recipient string) {

	m := gomail.NewMessage()
	m.SetHeader("From", "EWS ALERT DISPATCHER <"+From_mail+">")
	m.SetHeader("To", recipient)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", msg)

	d := gomail.NewPlainDialer(SMTP_Host, 465, From_mail, Mail_password)
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}
