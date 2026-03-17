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
	alldispatcher := notificationobserver.NewNotificationDispatcher(&emai, &sms, &push)
	loggerObeserver := notificationobserver.NewLoggerObserver(alldispatcher)

	simpleDispatcher := notificationobserver.NewNotificationDispatcher(&emai)
	simpleObserver := notificationobserver.NewLoggerObserver(simpleDispatcher)

	notificationObservableList := notificationobserver.NotificationObeservableList{}
	notificationObservableList.AddObserver(loggerObeserver)
	notificationObservableList.AddObserver(simpleObserver)

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
