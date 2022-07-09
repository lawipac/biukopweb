package main

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/smtp"
	"strings"
)

const formPrefix = "/form/"

type formInput struct {
	w http.ResponseWriter
	r *http.Request
}

func formMain(w http.ResponseWriter, r *http.Request) {
	fi := formInput{w, r}
	fi.contactUs()
}

func (m *formInput) contactUs() {
	ret := map[string]interface{}{}
	ret["success"] = true
	//	e := m.r.ParseForm()
	e := m.r.ParseMultipartForm(32 * 1024 * 1024)
	if e != nil {
		return
	}

	for key, values := range m.r.Form { // range over map
		for _, value := range values { // range over []string
			fmt.Println(key, value)
		}
	}

	name := m.r.FormValue("name")
	email := m.r.FormValue("email")
	message := m.r.FormValue("message")

	ret["name"] = name
	ret["email"] = email
	ret["message"] = message

	m.apiV1SendJson(ret)
	go sendContactUs(name, email, message)

}

func (m *formInput) apiV1SendJson(result interface{}) {
	//general setup
	m.w.Header().Set("Content-Type", "application/json;charset=UTF-8")

	out, e := json.Marshal(result)
	if e != nil {
		log.Warn("Cannot convert result to json ", result)
	}
	m.w.Header().Set("Content-Type", "text/json; charset=utf-8")
	//apiV1AddTrackingCookie(w, r, ss) //always the last one to set cookies
	fmt.Fprint(m.w, string(out))
}

func sendContactUs(name string, email string, userInput string) {
	// Sender data.
	from := "mailer@biukop.com.au"
	password := "hpfitsrujgkewcdw"

	// Receiver email address.
	to := []string{
		"patrick@biukop.com.au",
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	raw := `Subject: {name} [ Contact form on Biukop Web ] 
Content-Type: text/plain; charset="UTF-8"

    Dear Manager,

    We received a a form submission from biukop.com.au 
    name  : {name}    
    email : {email}
    message:

    {message}

    Kind Regards
    Biukop Mailing service team.
    `

	raw = strings.Replace(raw, "{name}", name, -1)
	raw = strings.Replace(raw, "{email}", email, -1)
	raw = strings.Replace(raw, "{message}", userInput, -1)

	// Message.
	message := []byte(raw)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
