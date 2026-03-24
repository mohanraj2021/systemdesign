package decorator

import (
	"time"

	"plugplay.com/notification/priority"
)

type INotification interface {
	GetContent() priority.NotificationPriority
}

type NotificationDecorator struct {
	Notification priority.NotificationPriority // ← VIOLATES SRP
}

func (d *NotificationDecorator) GetContent() priority.NotificationPriority {
	return d.Notification
}

type SimpleNotification struct {
	Notification INotification
}

func (s *SimpleNotification) GetContent() priority.NotificationPriority {
	content := s.Notification.GetContent()
	return content
}

type UrgentNotificationWithTimestamp struct {
	Notification INotification
}

func (u *UrgentNotificationWithTimestamp) GetContent() priority.NotificationPriority {
	content := u.Notification.GetContent()
	content.Timestamp = int(time.Now().Unix())
	return content
}
