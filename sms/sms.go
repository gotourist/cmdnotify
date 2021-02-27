package sms

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/vonage/vonage-go-sdk"
)

func SendMessage(clientnumber string, messageInfo string) error {
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
	api_Key := os.Getenv("APIKEY")       //vonage api_key
	api_Secret := os.Getenv("APISECRET") //vonage api_secret
	auth := vonage.CreateAuthFromKeySecret(api_Key, api_Secret)
	smsClient := vonage.NewSMSClient(auth)
	response, resperr, err := smsClient.Send("Vonage APIs", clientnumber, messageInfo, vonage.SMSOpts{Type: "unicode"})

	if response.Messages[0].Status == "0" {
		fmt.Println("Message sent")
		log.Println("Message sent")
		return nil
	}
	fmt.Println("Error occured while sending message:")
	fmt.Println(err)
	fmt.Println(resperr)
	log.Println("Error occured while sending message:")
	log.Println(err)
	log.Println(resperr)
	return err
}
