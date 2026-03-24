package main

import (
	"container/heap"
	"math/rand/v2"

	"plugplay.com/notification/enums"
	decorator "plugplay.com/notification/notificationDecorator"
	notificationobserver "plugplay.com/notification/notificationObserver"
	notificationservice "plugplay.com/notification/notificationService"
	notificationstrategy "plugplay.com/notification/notificationStrategy"
	"plugplay.com/notification/priority"
)

func main() {
	v := 100
	notificationService := notificationservice.NewService()
	notificationQuee := make(priority.NotificationPriorityList, 0)
	heap.Init(&notificationQuee)
	heap.Push(&notificationQuee, priority.NotificationPriority{ID: rand.IntN(v), OrderId: rand.Int(), Priority: enums.CriticalPriority, Type: string(enums.SimpleDecoratorType), Message: "Hello, this is a simple notification", Status: enums.FailedStatus, Channels: []enums.ChannelType{enums.EmailChannelType}})
	heap.Push(&notificationQuee, priority.NotificationPriority{ID: rand.IntN(v), OrderId: rand.Int(), Priority: enums.CriticalPriority, Type: string(enums.SimpleDecoratorType), Message: "Hello, this is a simple notification", Status: enums.FailedStatus, Channels: []enums.ChannelType{enums.EmailChannelType}})
	heap.Push(&notificationQuee, priority.NotificationPriority{ID: rand.IntN(v), OrderId: rand.Int(), Priority: enums.HighPriority, Type: string(enums.SimpleDecoratorType), Message: "Hello, this is a simple notification", Status: enums.PendingStatus, Channels: []enums.ChannelType{enums.PushChannelType}})
	heap.Push(&notificationQuee, priority.NotificationPriority{ID: rand.IntN(v), OrderId: rand.Int(), Priority: enums.MediumPriority, Type: string(enums.UrgentDecoratorType), Message: "Hello, this is a simple notification", Status: enums.SentStatus, Channels: []enums.ChannelType{enums.SMSChannelType}})
	heap.Push(&notificationQuee, priority.NotificationPriority{ID: rand.IntN(v), OrderId: rand.Int(), Priority: enums.LowPriority, Type: string(enums.UrgentDecoratorType), Message: "Hello, this is a simple notification", Status: enums.DeliveredStatus, Channels: []enums.ChannelType{enums.EmailChannelType}})
	heap.Push(&notificationQuee, priority.NotificationPriority{ID: rand.IntN(v), OrderId: rand.Int(), Priority: enums.HighPriority, Type: string(enums.UrgentDecoratorType), Message: "Hello, this is a simple notification", Status: enums.PendingStatus, Channels: []enums.ChannelType{enums.PushChannelType, enums.EmailChannelType}})

	for _, notification := range notificationQuee {
		var notDecorator = decorator.NotificationDecorator{
			Notification: notification,
		}
		switch notification.Status {
		case enums.PendingStatus:
			switch notification.Type {
			case string(enums.SimpleDecoratorType):
				strategies := BuildNotificationStrategyList(notification.Channels)
				dispatcher := notificationobserver.NewNotificationDispatcher(&notificationstrategy.NotStrategy{Strategy: strategies})
				loggerObeserver := notificationobserver.NewLoggerObserver(dispatcher)
				decorators := decorator.SimpleNotification{
					Notification: &notDecorator,
				}
				notificationService.AddObserver(loggerObeserver)
				notificationService.Notify(&decorators)
				notificationService.RemoveObserver(loggerObeserver)
			case string(enums.UrgentDecoratorType):
				strategies := BuildNotificationStrategyList(notification.Channels)
				dispatcher := notificationobserver.NewNotificationDispatcher(&notificationstrategy.NotStrategy{Strategy: strategies})
				loggerObeserver := notificationobserver.NewLoggerObserver(dispatcher)
				urgentDecorator := decorator.UrgentNotificationWithTimestamp{
					Notification: &notDecorator,
				}
				notificationService.AddObserver(loggerObeserver)
				notificationService.Notify(&urgentDecorator)
				notificationService.RemoveObserver(loggerObeserver)

			}
		case enums.SentStatus:
		case enums.FailedStatus:

		case enums.DeliveredStatus:

		}
	}

}

func BuildNotificationStrategyList(channels []enums.ChannelType) []notificationstrategy.NotificationStrategy {

	strategies := []notificationstrategy.NotificationStrategy{}
	for _, chanel := range channels {
		switch chanel {
		case enums.EmailChannelType:
			strategies = append(strategies, &notificationstrategy.EmailNotification{EmailAddress: "mohan@plugplay.com"})
		case enums.SMSChannelType:
			strategies = append(strategies, &notificationstrategy.SMSNotification{PhoneNumber: "1234567890"})
		case enums.PushChannelType:
			strategies = append(strategies, &notificationstrategy.PushNotification{DeviceToken: "1234567890"})
		}
	}
	return strategies
}
