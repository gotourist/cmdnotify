package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/gotourist/cmdnotify/mail"
	"github.com/gotourist/cmdnotify/sms"
)

type Client struct {
	Id        int
	TelNumber string
	Email     string
}

type Purchase struct {
	Id      int
	Product string
	Cost    int
}

func main() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
		fmt.Println(err)
	}
	log.SetOutput(file)
	var client Client
	var purchase Purchase

	//  2 version of command-line notification sender
	sendmail := flag.NewFlagSet("sendmail", flag.ExitOnError)
	email := sendmail.String("email", "", "An email of client")
	mailpurchaseid := sendmail.Int("id", 0, "Purchase id")
	mailproduct := sendmail.String("product", "", "Product name")
	mailcost := sendmail.Int("cost", 0, "Cost of a product")

	sendsms := flag.NewFlagSet("sendsms", flag.ExitOnError)
	telnumber := sendsms.String("tel", "", "Telephone number of client")
	smspurchaseid := sendsms.Int("id", 0, "Purchase id")
	smsproduct := sendsms.String("product", "", "Product name")
	smscost := sendsms.Int("cost", 0, "Cost of a product")

	if len(os.Args) < 2 {
		fmt.Println("expected 'sendmail' or 'sendsms' subcommands")
		log.Println("expected 'sendmail' or 'sendsms' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "sendmail":
		sendmail.Parse(os.Args[2:])
		client.Email = *email
		purchase.Product = *mailproduct
		purchase.Cost = *mailcost
		purchase.Id = *mailpurchaseid
		if len(sendmail.Args()) > 0 {
			fmt.Println("Undefined command: ", sendmail.Args())
			log.Println("Undefined command: ", sendmail.Args())
			os.Exit(1)
		} else if client.Email != "" {
			//make message
			t, _ := template.ParseFiles("template.html")
			var body bytes.Buffer
			mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
			body.Write([]byte(fmt.Sprintf("Subject: Notification mail about your purchase \n%s\n\n", mimeHeaders)))
			t.Execute(&body, Purchase{
				Id:      purchase.Id,
				Product: purchase.Product,
				Cost:    purchase.Cost,
			})
			err = mail.SendMail(client.Email, body.Bytes())
			if err != nil {
				log.Fatal(err)
				fmt.Println(err)
			}
		} else {
			fmt.Println("Client email is not included")
			log.Println("Client email is not included")
		}
	case "sendsms":
		sendsms.Parse(os.Args[2:])
		client.TelNumber = *telnumber
		purchase.Product = *smsproduct
		purchase.Cost = *smscost
		purchase.Id = *smspurchaseid
		if len(sendmail.Args()) > 0 {
			fmt.Println("Undefined command: ", sendsms.Args())
			log.Println("Undefined command: ", sendsms.Args())
			os.Exit(1)
		} else if client.TelNumber != "" {
			text := fmt.Sprintf(`Notification message about your purchase:
Purchase Id: %d
Product: %s
Cost: %d $`, purchase.Id, purchase.Product, purchase.Cost)
			err = sms.SendMessage(client.TelNumber, text)
			if err != nil {
				log.Fatal(err)
				fmt.Println(err)
			}
		} else {
			fmt.Println("Client tel number is not included")
			log.Println("Client tel number is not included")
		}
	default:
		log.Println("expected 'sendmail' or 'sendsms' subcommands")
		fmt.Println("expected 'sendmail' or 'sendsms' subcommands")
		os.Exit(1)
	}

	//  1 version of command-line notification sender
	// 	notificationtype := flag.String("type", "", "Notification type")
	// 	email := flag.String("email", "", "An email of client")
	// 	telnumber := flag.String("tel", "", "Telephone number of client")
	// 	purchaseid := flag.Int("id", 0, "Purchase id")
	// 	product := flag.String("product", "", "Product name")
	// 	cost := flag.Int("cost", 0, "Cost of a product")
	// 	flag.Parse()
	// 	client.Email = *email
	// 	client.TelNumber = *telnumber
	// 	purchase.Product = *product
	// 	purchase.Cost = *cost
	// 	purchase.Id = *purchaseid
	// 	switch *notificationtype {
	// 	case "email":
	// 		if client.Email != "" {
	// 			//make message
	// 			t, _ := template.ParseFiles("template.html")
	// 			var body bytes.Buffer
	// 			mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	// 			body.Write([]byte(fmt.Sprintf("Subject: Notification mail about your purchase \n%s\n\n", mimeHeaders)))
	// 			t.Execute(&body, Purchase{
	// 				Id:      purchase.Id,
	// 				Product: purchase.Product,
	// 				Cost:    purchase.Cost,
	// 			})
	// 			err = mail.SendMail(client.Email, body.Bytes())
	// 			if err != nil {
	// 				log.Fatal(err)
	// 				fmt.Println(err)
	// 			}
	// 		} else {
	// 			fmt.Println("Client email is not included")
	// 			log.Println("Client email is not included")
	// 		}
	// 	case "sms":
	// 		if client.TelNumber != "" {
	// 			text := fmt.Sprintf(`Notification message about your purchase:
	// Purchase Id: %d
	// Product: %s
	// Cost: %d $`, purchase.Id, purchase.Product, purchase.Cost)
	// 			err = sms.SendMessage(client.TelNumber, text)
	// 			if err != nil {
	// 				log.Fatal(err)
	// 				fmt.Println(err)
	// 			}
	// 		} else {
	// 			fmt.Println("Client tel number is not included")
	// 			log.Println("Client tel number is not included")
	// 		}
	// 	default:
	// 		fmt.Println("notification type is not selected")
	// 		log.Println("notification type is not selected")
	// 	}
}
