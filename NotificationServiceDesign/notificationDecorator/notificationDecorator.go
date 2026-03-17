package decorator

import (
	"fmt"
	"time"
)

type INotification interface {
	GetContent() any
}

type NotificationDecorator struct {
	Message string
}

func (d *NotificationDecorator) GetContent() any {
	return d.Message
}

type SimpleNotification struct {
	Notification INotification
}

func (s *SimpleNotification) GetContent() any {
	return fmt.Sprintf("%v", s.Notification.GetContent())
}

type UrgentNotificationWithTimestamp struct {
	Notification INotification
}

func (u *UrgentNotificationWithTimestamp) GetContent() any {
	return fmt.Sprintf("%v %v", u.Notification.GetContent(), time.Now().Format("02-01-2006 15:04:05"))
}
