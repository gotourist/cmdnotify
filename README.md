# CmdNotify

CmdNotify is Command-line Notification sender application 

## Installation

As library  

```bash
go get github.com/gotourist/cmdnotify
```

## Usage
Create .env file and Add your application configuration to your .env file in the root of your project like below:

```bash
FROM = sample@example.com //sender email address
PASSWORD = xxxxxxx //sender email password
SMPTHOST = smpt.gmail.com //sender email provider host
SMPTPORT = xxx //sender email provider port
APIKEY = xxxxxx //vonage apikey
APISECRET = xxxxxx //vonage apisecret
```

## Command Mode Usage 1 version

Assuming you've installed the command as above and now you are in $GOPATH/src/github.com/gotourist/aliftech/

if you want to send notification as email use this command in linux

```bash
./send -type=email -email=example@gmail.com -product=car -cost=30000 -id=11
```

if you want to send notification as sms use this command in linux

```bash
./send -type=sms -tel=998909998877 -product=car -cost=30000 -id=11
```

if you want to use command in windows change ./send command to send.exe

## Command Mode Usage 2 version

Assuming you've installed the command as above and now you are in $GOPATH/src/github.com/gotourist/aliftech/

if you want to send notification as email use this command in linux

```bash
./tsend sendmail -email=example@gmail.com -product=car -cost=30000 -id=11
```

if you want to send notification as sms use this command in linux

```bash
./send sendsms -tel=998909998877 -product=car -cost=30000 -id=11
```

if you want to use command in windows change ./tsend command to tsend.exe


