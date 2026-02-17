package factories

import (
	"github.com/zomatodesign/managers"
	"github.com/zomatodesign/models"
)

type SmsNotificationService struct {
}

func (ns *SmsNotificationService) Notify(user models.User, message string) {
	println("Sending notification to", user.Name, "with message:", message)
}

type EmailNotificationService struct {
}

func (ns *EmailNotificationService) Notify(user models.User, message string) {
	println("Sending notification to", user.Name, "with message:", message)
}

func NotificationServiceFactory(service string) managers.NotificationService {
	switch service {
	case "sms":
		return &SmsNotificationService{}
	case "email":
		return &EmailNotificationService{}
	default:
		return &SmsNotificationService{}
	}
}
