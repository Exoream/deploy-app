package email

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"gopkg.in/gomail.v2"
)

// SendEmailSMTP sends an email using SMTP with the given template and recipients.
func SendEmailSMTP(to []string, template string, data interface{}) (bool, error) {
	emailHost := os.Getenv("EMAIL_HOST")
	emailFrom := os.Getenv("EMAIL_FROM")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	emailPortStr := os.Getenv("EMAIL_PORT")
	emailPort, _ := strconv.Atoi(emailPortStr)

	m := gomail.NewMessage()
	m.SetHeader("From", emailFrom)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", "Verifikasi Akun Anda")

	emailContent := strings.Replace(template, "{{.verificationLink}}", data.(string), -1)

	// Set HTML body
	m.SetBody("text/html", emailContent)

	d := gomail.NewDialer(emailHost, emailPort, emailFrom, emailPassword)

	// Send email
	if err := d.DialAndSend(m); err != nil {
		return false, err
	}
	return true, nil
}


// SendEmailForPickup sends an email using SMTP with the given template and recipients.
func SendEmailForPickup(to string, template string, data interface{}) (bool, error) {
	emailHost := os.Getenv("EMAIL_HOST")
	emailFrom := os.Getenv("EMAIL_FROM")
	emailPassword := os.Getenv("EMAIL_PASSWORD")
	emailPortStr := os.Getenv("EMAIL_PORT")
	emailPort, _ := strconv.Atoi(emailPortStr)

	m := gomail.NewMessage()
	m.SetHeader("From", emailFrom)
	m.SetHeader("To", to)
	m.SetHeader("Subject", "Detail Pickup Anda")

	// Assert emailData as the correct type
	d, ok := data.(struct {
		UserID      string
		PickupDate  string
		PickupID    string
		Status      string
		TotalPoints int
	})
	if !ok {
		return false, errors.New("Invalid email data type")
	}

	// Replace placeholders in the template with data
	emailContent := template // Assuming 'template' contains the HTML template

	emailContent = strings.Replace(emailContent, "{{.UserID}}", d.UserID, -1)
	emailContent = strings.Replace(emailContent, "{{.PickupID}}", d.PickupID, -1)
	emailContent = strings.Replace(emailContent, "{{.PickupDate}}", d.PickupDate, -1)
	emailContent = strings.Replace(emailContent, "{{.Status}}", d.Status, -1)
	emailContent = strings.Replace(emailContent, "{{.TotalPoints}}", strconv.Itoa(d.TotalPoints), -1)

	// Set HTML body
	m.SetBody("text/html", emailContent)

	dialer := gomail.NewDialer(emailHost, emailPort, emailFrom, emailPassword)

	// Send email
	if err := dialer.DialAndSend(m); err != nil {
		return false, err
	}
	return true, nil
}
