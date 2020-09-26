package emailer

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/my/repo/go/src/github.com/joho/godotenv"
)

type smtpServer struct {
	host string
	port string
}

// Address URI to smtp server
func (s *smtpServer) Address() string {
	return s.host + ":" + s.port
}

// Birthday struct
type Birthday struct {
	Name          string
	Email         string
	JokeSetup     string
	JokePunchLine string
}

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

// SendBdayEmail function
func SendBdayEmail(info Birthday) string {
	// Sender data.
	from := goDotEnvVariable("EMAIL")
	password := goDotEnvVariable("EMAIL-PASS")

	// Receiver Email address.
	to := []string{
		from, info.Email,
	}
	// smtp server configuration.
	smtpServer := smtpServer{host: "smtp.gmail.com", port: "587"}
	// Message.
	strMessage := fmt.Sprintf("Happy birthday %s! Here is a joke to get your day started right!\n\n\n%s\n\n%s\n\n\n\n\nBest Regards,\n\nAustin", info.Name, info.JokeSetup, info.JokePunchLine)
	msg := "From: " + from + "\n" +
		"To: " + info.Email + "\n" +
		"Subject: Happy Birthday\n\n" + strMessage
	// fmt.Println(strMessage)
	message := []byte(msg)
	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpServer.host)
	// Sending Email.
	err := smtp.SendMail(smtpServer.Address(), auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return "Email Err"
	}
	fmt.Println("Email Sent!")
	return "Success"
}
