package main

import (
	"fmt"

	"github.com/wuriyanto48/goemail"
)

func main() {
	authEmail := "wuriyanto007@gmail.com"
	authPassword := "your email password"
	authHost := "smtp.gmail.com"
	address := "smtp.gmail.com:587"
	to := []string{"wuriyanto48@yahoo.co.id"}
	from := "wuriyanto007@gmail.com"
	subject := "Golang email"
	body := "Golang email sent..."
	email := goemail.New(to, address, from, subject, body, authEmail, authPassword, authHost)

	emailData := struct {
		Username string
		URL      string
	}{
		Username: "Wuriyanto",
		URL:      "wuriyanto.com",
	}

	err := execute(email, "template.html", emailData)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("email sent")

}

func execute(u goemail.EmailSender, fileName string, data interface{}) error {
	err := u.SetTemplate(fileName, data)
	if err != nil {
		return err
	}

	err = u.Send()
	if err != nil {
		return err
	}

	return nil
}
