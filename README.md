# CmdNotify

CmdNotify is Command-line Notification sender application 

## Installation

As library  

```bash
go get github.com/gotourist/cmdnotify
```

## Usage
Add your configurations into .env file in the root of this project like below:

```bash
FROM = xxxxxxxx //sender email address
PASSWORD = xxxxxxxx //sender email password
SMPTHOST = xxxxxxxx //sender email provider host by default it is smtp.gmail.com
SMPTPORT = xxxx //sender email provider port host by default it is 587
APIKEY = xxxxxxxx //vonage apikey
APISECRET = xxxxxxxx //vonage apisecret
```

## Command Mode Usage 1 version

Assuming you've installed the project as above and now you are in $GOPATH/src/github.com/gotourist/cmdnotify/

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

Assuming you've installed the project as above and now you are in $GOPATH/src/github.com/gotourist/cmdnotify/

if you want to send notification as email use this command in linux

```bash
./tsend sendmail -email=example@gmail.com -product=car -cost=30000 -id=11
```

if you want to send notification as sms use this command in linux

```bash
./tsend sendsms -tel=998909998877 -product=car -cost=30000 -id=11
```

if you want to use command in windows change ./tsend command to tsend.exe


