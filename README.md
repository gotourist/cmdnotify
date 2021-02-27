# CmdNotify Command-line Notification sender application

CmdNotify is Command-line Notification sender application 

## Installation

As library  

```bash
go get github.com/gotourist/cmdnotify
```

## Usage
Create .env file and Add your application configuration to your .env file in the root of your project:

```bash
//to send email
FROM = sample@example.com //sender email address
PASSWORD = xxxxxxx //sender email password
SMPTHOST = smpt.gmail.com //sender email provider host
SMPTPORT = xxx //sender email provider port

//to send sms messages
APIKEY = xxxxxx //vonage apikey
APISECRET = xxxxxx //vonage apisecret
```

Then in your Go app you can do something like

```golang
//get environmetal variable for email
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
```

```golang
//get environmetal variable for sms messages
	e := godotenv.Load("./.env")
	if e != nil {
		fmt.Println(e)
		log.Println(e)
	}
	api_Key := os.Getenv("APIKEY")       //vonage api_key
	api_Secret := os.Getenv("APISECRET") //vonage api_secret
```

## Command Mode Usage

Assuming you've installed the command as above and now you are in $GOPATH/src/github.com/gotourist/aliftech/

if you want to send notification as email use this command in linux

```bash
./send -type=email -email=example@gmail.com -product=car -cost=30000 -id=11
```

if you want to send notification as sms use this command in linux

```bash
./send -type=sms -tel=998909998877 -product=car -cost=30000 -id=11
```

if you want to send notification as email use this command in windows

```bash
send.exe -type=email -email=example@gmail.com -product=car -cost=30000 -id=11
```

if you want to send notification as sms use this command in windows

```bash
send.exe -type=sms -tel=998909998877 -product=car -cost=30000 -id=11
```

