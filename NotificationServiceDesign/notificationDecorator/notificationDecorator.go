package decorator

import (
	"fmt"
	"time"
)

type INotification interface {
	GetContent() any
}

type NotificationDecorator struct {
	notification INotification
}

func (d *NotificationDecorator) GetContent() any {
	return d.notification.GetContent()
}

type SimpleNotification struct {
	Message string
}

func (s *SimpleNotification) GetContent() any {
	return s.Message
}

type UrgentNotificationWithTimestamp struct {
	Message   string
	Timestamp time.Time
}

func (u *UrgentNotificationWithTimestamp) GetContent() any {
	return fmt.Sprintf("%v %v", u.Message, u.Timestamp)
}
