package mail

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func SendMail(mailreciever string, Message []byte) error {
	file, err := os.OpenFile("./logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	//get environmetal variable
	e := godotenv.Load("./.env")
	if e != nil {
		fmt.Println(e)
		log.Println(e)
	}
	// Sender data.
	from := os.Getenv("FROM")         //sender email address
	password := os.Getenv("PASSWORD") //sender email password

	// smtp server configuration.
	smtpHost := os.Getenv("SMPTHOST") //sender email smpt host
	smtpPort := os.Getenv("SMPTPORT") //sender email smpt port

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Receiver email address.
	to := []string{
		mailreciever,
	}

	// Sending email.
	err = smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, Message)
	if err != nil {
		fmt.Println("Error while sending mail: ", err)
		log.Println("Error while sending mail: ", err)
		return err
	}
	fmt.Println("Email Sent!")
	log.Println("Email Sent!")
	return nil
}
