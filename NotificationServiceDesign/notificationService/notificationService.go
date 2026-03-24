package notificationservice

import (
	notificationobserver "plugplay.com/notification/notificationObserver"
)

func NewService() *notificationobserver.NotificationObeservableList {
	// Allstrategy := notificationstrategy.NotStrategy{
	// 	Strategy: make([]notificationstrategy.NotificationStrategy, 0),
	// }
	// alldispatcher := notificationobserver.NewNotificationDispatcher(&Allstrategy)
	// loggerObeserver := notificationobserver.NewLoggerObserver(alldispatcher)

	// onlyMailStrategy := notificationstrategy.NotStrategy{
	// 	Strategy: make([]notificationstrategy.NotificationStrategy, 0),
	// }
	// simpleDispatcher := notificationobserver.NewNotificationDispatcher(&onlyMailStrategy)
	// simpleObserver := notificationobserver.NewLoggerObserver(simpleDispatcher)

	observable := notificationobserver.NotificationObeservableList{}
	// observable.AddObserver(loggerObeserver)
	// observable.AddObserver(simpleObserver)

	return &observable
}
