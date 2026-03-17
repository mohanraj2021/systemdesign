package notificationobserver

import (
	"fmt"

	decorator "plugplay.com/notification/notificationDecorator"
	notificationstrategy "plugplay.com/notification/notificationStrategy"
)

type NotificationObservable interface {
	RemoveObserver(observer NotificationObeserver)
	AddObserver(observer NotificationObeserver)
	Notify(notification decorator.INotification)
}

type NotificationObeserver interface {
	Update(notification decorator.INotification)
}

type NotificationObeservableList struct {
	observers []NotificationObeserver
}

type LoggerObserver struct {
	NotificatioStrat []notificationstrategy.NotificationStrategy
}

func (l *LoggerObserver) Update(notification decorator.INotification) {
	// log the notification
	msg := fmt.Sprintf("Received notification: %v\n", notification.GetContent())
	for _, strategy := range l.NotificatioStrat {
		strategy.SendNotification(msg)
	}
}

func (n *NotificationObeservableList) RemoveObserver(observer NotificationObeserver) {
	// remove the observer from the list of observers
	for i, obs := range n.observers {
		if obs == observer {
			n.observers = append(n.observers[:i], n.observers[i+1:]...)
			break
		}
	}
}

func (n *NotificationObeservableList) AddObserver(observer NotificationObeserver) {
	// add the observer to the list of observers
	n.observers = append(n.observers, observer)
}

func (n *NotificationObeservableList) Notify(notification decorator.INotification) {
	// notify all the observers
	for _, obs := range n.observers {
		obs.Update(notification)
	}
}
