package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net"
	"net/mail"
	"net/smtp"
)

func createConn(host, servername string, isSecure bool) (net.Conn, error) {
	if isSecure {
		// TLS config
		tlsconfig := &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         host,
		}

		return tls.Dial("tcp", servername, tlsconfig)
	}

	return net.Dial("tcp", servername)
}

func main() {
	senderEmail := flag.String("sender-email", "", "Sender's email")
	senderPassword := flag.String("sender-password", "", "Sender's password")
	receiverEmail := flag.String("receiver-email", "", "Receiver's email")
	smtpHost := flag.String("host", "", "Smtp host address")
	smtpPort := flag.Int("port", 0, "Smtp port number")
	smtpSecure := flag.Bool("secure", false, "Smtp over TLS")

	flag.Parse()

	from := mail.Address{"", *senderEmail}
	to := mail.Address{"", *receiverEmail}

	subj := "This is the email subject"
	body := "This is an example body.\n With two lines."

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

	// Connect to the SMTP Server
	servername := fmt.Sprintf("%v:%d", *smtpHost, *smtpPort)

	host, _, _ := net.SplitHostPort(servername)

	auth := smtp.PlainAuth("", *senderEmail, *senderPassword, host)

	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)

	conn, err := createConn(host, servername, *smtpSecure)
	if err != nil {
		log.Panic(err)
	}

	c, err := smtp.NewClient(conn, host)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	c.Quit()

	fmt.Println("Email Sent Successfully!")
}
