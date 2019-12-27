# gogmail

gogmail is a simple go program that sends an email via your gmail account.  
The interesting thing about the app is that it implements the SMTP protocol internally, rather than using a [high-level package](https://golang.org/pkg/net/smtp/).  
Additionally, it will print all on-going SMTP communication from both client and server.  
This can be used as a exploratory learning example for how SMTP works under the hood.

Prerequisites:
 * Make an `.app_pass` file with the [app password](https://support.google.com/accounts/answer/185833) for your account

How to run it:
```
go run *.go
```
