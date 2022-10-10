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
	if err := smtp.SendMail(fmt.Sprintf("%s:%d", SMTP_Host, 465), auth, From_mail, recipient, message); err != nil {
		log.Printf("Error sending mail %v", err)
		return
	}

	log.Println("Email Sent Successfully!")
}

func SendMails(subject, msg string, recipients []string) {
	mailwg.Add(len(recipients))
	for _, v := range recipients {
		go func(recipient string) {
			defer mailwg.Done()
			SendSSLMail(subject, msg, recipient)
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

/*
	{
		"headline":"hello",
		"body": "hello body",
		"platforms": [{"name":"facebook"}, {"name":"telegram"}],
		"mails": [{"address":"salemododa2@gmail.com"}, {"address" :"robtyler0701@gmail.com"}]
	}
*/
