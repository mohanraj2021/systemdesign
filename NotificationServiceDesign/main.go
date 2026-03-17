package main

import (
	"time"

	decorator "plugplay.com/notification/notificationDecorator"
	notificationobserver "plugplay.com/notification/notificationObserver"
	notificationstrategy "plugplay.com/notification/notificationStrategy"
)

func main() {
	emai := notificationstrategy.EmailNotification{}
	sms := notificationstrategy.SMSNotification{}
	push := notificationstrategy.PushNotification{}

	loggerObeserver := notificationobserver.LoggerObserver{
		NotificatioStrat: []notificationstrategy.NotificationStrategy{&emai, &sms, &push},
	}

	simpleObserver := notificationobserver.LoggerObserver{
		NotificatioStrat: []notificationstrategy.NotificationStrategy{&emai},
	}

	notificationObservableList := notificationobserver.NotificationObeservableList{}
	notificationObservableList.AddObserver(&loggerObeserver)
	notificationObservableList.AddObserver(&simpleObserver)

	simpleNotification := decorator.SimpleNotification{
		Message: "Hello, this is a simple notification",
	}

	urgentNotification := decorator.UrgentNotificationWithTimestamp{
		Message:   "Hello, this is a urgent notification",
		Timestamp: time.Now(),
	}
	notificationObservableList.Notify(&simpleNotification)
	notificationObservableList.Notify(&urgentNotification)
}
