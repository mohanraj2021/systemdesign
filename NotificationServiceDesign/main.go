package main

import (
	decorator "plugplay.com/notification/notificationDecorator"
	notificationobserver "plugplay.com/notification/notificationObserver"
	notificationstrategy "plugplay.com/notification/notificationStrategy"
)

func main() {
	emai := notificationstrategy.EmailNotification{}
	sms := notificationstrategy.SMSNotification{}
	push := notificationstrategy.PushNotification{}
	list := notificationstrategy.NewNotificationstrategyList(&emai, &sms, &push)
	loggerObeserver := notificationobserver.LoggerObserver{
		NotificatioStrat: list,
	}

	simpleObserver := notificationobserver.LoggerObserver{
		NotificatioStrat: notificationstrategy.NewNotificationstrategyList(&emai),
	}

	notificationObservableList := notificationobserver.NotificationObeservableList{}
	notificationObservableList.AddObserver(&loggerObeserver)
	notificationObservableList.AddObserver(&simpleObserver)

	notificationDecorator := decorator.NotificationDecorator{
		Message: "Hello, this is a simple notification",
	}

	simpleDecorator := decorator.SimpleNotification{
		Notification: &notificationDecorator,
	}

	urgentNotification := decorator.UrgentNotificationWithTimestamp{
		Notification: &simpleDecorator,
	}
	notificationObservableList.Notify(&notificationDecorator)
	notificationObservableList.Notify(&urgentNotification)
}
