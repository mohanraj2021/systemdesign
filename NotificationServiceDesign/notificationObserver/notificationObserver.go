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

type NotificationDispatcher interface {
	Dispatch(message string) error
}

type NotificationObeserver interface {
	Update(notification decorator.INotification)
}

type Dispatcher struct {
	strategies notificationstrategy.NotificationStrategyList
}

func NewNotificationDispatcher(strategies *notificationstrategy.NotStrategy) *Dispatcher {
	return &Dispatcher{
		strategies: notificationstrategy.NewNotificationstrategyList(strategies),
	}
}

func (d *Dispatcher) Dispatch(message string) error {
	return d.strategies.SendNotification(message)
}

type NotificationObeservableList struct {
	observers []NotificationObeserver
}

type LoggerObserver struct {
	dispatcher NotificationDispatcher
}

func NewLoggerObserver(dispatcher NotificationDispatcher) *LoggerObserver {
	return &LoggerObserver{dispatcher: dispatcher}
}

func (l *LoggerObserver) Update(notification decorator.INotification) {
	msg := fmt.Sprintf("Received notification: %v\n", notification.GetContent())
	l.dispatcher.Dispatch(msg)
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
