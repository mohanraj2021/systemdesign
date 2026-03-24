package notificationstrategy

import "fmt"

type NotificationStrategy interface {
	SendNotification(notification string) error
}

// This interface is properly defined - it establishes a contract for notification strategies
// However, it could benefit from additional methods like Validate() or GetType() for better extensibility

type NotStrategy struct {
	Strategy []NotificationStrategy
}

func (ns *NotStrategy) SendNotification(notification string) error {

	for _, strategy := range ns.Strategy {
		err := strategy.SendNotification(notification)
		if err != nil {
			return err
		}
	}
	return nil
}

type EmailNotification struct {
	EmailAddress string
}

func (e *EmailNotification) SendNotification(notification string) error {
	// send the notification to the email
	fmt.Printf("mail sent : %v\n", notification)
	return nil
}

type SMSNotification struct {
	PhoneNumber string
}

func (s *SMSNotification) SendNotification(notification string) error {
	// send the notification to the phone number
	fmt.Printf("sms sent : %v\n", notification)
	return nil
}

type PushNotification struct {
	DeviceToken string
}

func (p *PushNotification) SendNotification(notification string) error {
	// send the notification to the device
	fmt.Printf("push notification sent : %v\n", notification)
	return nil
}

type NotificationStrategyList []NotificationStrategy

func (n NotificationStrategyList) SendNotification(notification string) error {
	for _, strategy := range n {
		err := strategy.SendNotification(notification)
		if err != nil {
			return err
		}
	}
	return nil
}

func NewNotificationstrategyList(strategies NotificationStrategy) NotificationStrategyList {
	return NotificationStrategyList{strategies}
}
