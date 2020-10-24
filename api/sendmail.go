package handler

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

//Feedback is feedback struct
type Feedback struct {
	Name    string
	Email   string
	Message string
}

//Handler is the default handler
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/api/sendmail" || r.Method != "POST" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	var fb Feedback
	err := json.NewDecoder(r.Body).Decode(&fb)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, body, err := SendMail(fb)
	if err != nil {
		println("Error sending Email: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(res)
	w.Header().Set("content-type", "application/json")
	w.Write([]byte(body))
	return
}

//SendMail sends the email using Sendgrid
func SendMail(f Feedback) (res int, out string, err error) {
	from := mail.NewEmail(f.Name, f.Email)
	subject := "[Feedback] on CKAD tutorials"
	to := mail.NewEmail("Liptan Biswas", "me@liptanbiswas.com")
	message := mail.NewSingleEmail(from, subject, to, "", f.Message)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	r, err := client.Send(message)
	if err != nil {
		return r.StatusCode, r.Body, err
	}
	return r.StatusCode, r.Body, nil
}
