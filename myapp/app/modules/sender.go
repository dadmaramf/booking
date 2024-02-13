package modules

import "fmt"

type UserMailer struct {
	Mailer
}

func (u UserMailer) SendReport(reported_id string, email []string) error {
	err := u.Attach("/home/alexandra/hu/myproject/myapp/app/views/UserMailer/SendReport.html")
	if err != nil {
		fmt.Println(err)
	}
	return u.Send(H{
		"subject":     "a signature has been reported",
		"to":          []string(email),
		"reported_id": reported_id,
	})
}
