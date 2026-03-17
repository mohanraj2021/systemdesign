package notificationstrategy

import "fmt"

type NotificationStrategy interface {
	SendNotification(notification any) error
}

type EmailNotificationStrategy interface {
	SendNotification(notification any) error
}

type EmailNotification struct {
	EmailAddress string
}

func (e *EmailNotification) SendNotification(notification any) error {
	// send the notification to the email
	fmt.Printf("mail sent : %v\n", notification)
	return nil
}

type SMSNotificationStrategy interface {
	SendNotification(notification any) error
}

type SMSNotification struct {
	PhoneNumber string
}

func (s *SMSNotification) SendNotification(notification any) error {
	// send the notification to the phone number
	fmt.Printf("sms sent : %v\n", notification)
	return nil
}

type PushNotificationStrategy interface {
	SendNotification(notification any) error
}

type PushNotification struct {
	DeviceToken string
}

func (p *PushNotification) SendNotification(notification any) error {
	// send the notification to the device
	fmt.Printf("push notification sent : %v\n", notification)
	return nil
}
