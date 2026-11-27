package main

import "fmt"

type Sender interface {
	Send(message string) error
}

type StatusChecker interface {
	Status() string
}

type Closer interface {
	Close() error
}

type NotificationService interface {
	Sender
	StatusChecker
	Closer
}

type Email struct {
	Addres string
}

type SMS struct {
	PhoneNumber string
}

type Telegram struct {
	Username string
}

func (e Email) Send(message string) error {
	fmt.Printf("Sending email to %s: %s\n", e.Addres, message)
	return nil
}

func (e Email) Status() string {
	return "Email Sucessfully sent"
}

func (e Email) Close() error {
	fmt.Println("Closing email service")
	return nil
}

func (s SMS) Send(message string) error {
	fmt.Printf("Sending SMS to %s: %s\n", s.PhoneNumber, message)
	return nil
}

func (t Telegram) Send(message string) error {
	fmt.Printf("Sendgin Telegram message to @%s: %s\n", t.Username, message)
	return nil
}

func Process(n NotificationService) {
	_ = n.Send("Hello,")
	fmt.Println(n.Status())
	_ = n.Close()
}

func SendNotification(s Sender, message string) error {
	return s.Send(message)
}
func main() {
	email := Email{
		Addres: "user@gmail.com",
	}

	Process(email)

	_ = SendNotification(email, "this is a new message with send notification.")

	sms := SMS{
		PhoneNumber: "+987654321",
	}

	_ = SendNotification(sms, "hello this is my new phone number call me back please")
}
