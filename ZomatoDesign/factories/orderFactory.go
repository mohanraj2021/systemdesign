package factories

import (
	"github.com/zomatodesign/managers"
	"github.com/zomatodesign/models"
)

type NowOrderFactory struct {
	User        models.User
	Restaurant  models.Restaurant
	Items       []models.Menu
	TotalAmount float32
	Orders      map[int]models.Order
}

type ScheduledOrderFactory struct {
	User        models.User
	Restaurant  models.Restaurant
	Items       []models.Menu
	TotalAmount float32
	Time        string
	Orders      map[int]models.Order
}

func (of *NowOrderFactory) PlaceOrder(order *models.Order, cart *models.Cart, Restaurant *models.Restaurant) models.Order {
	order.Status = "Placed"
	order.Items = cart.Items
	order.TotalAmount = cart.GetTotal()
	of.Orders[order.Id] = *order
	return *order
}

func (of *NowOrderFactory) CancelOrder(orderID int) bool {
	if _, exists := of.Orders[orderID]; exists {
		order := of.Orders[orderID]
		order.Status = "Cancelled"
		of.Orders[orderID] = order
		return true
	}
	return false
}

func (of *NowOrderFactory) GetOrderStatus(orderID int) string {
	if order, exists := of.Orders[orderID]; exists {
		return order.Status
	}
	return ""
}

func (sof *ScheduledOrderFactory) PlaceOrder(order *models.Order, cart *models.Cart, Restaurant *models.Restaurant) models.Order {
	order.Status = "Scheduled"
	order.Items = cart.Items
	order.TotalAmount = cart.GetTotal()
	sof.Orders[order.Id] = *order
	return *order
}

func (sof *ScheduledOrderFactory) CancelOrder(orderID int) bool {
	if _, exists := sof.Orders[orderID]; exists {
		order := sof.Orders[orderID]
		order.Status = "Cancelled"
		sof.Orders[orderID] = order
		return true
	}
	return false
}

func (sof *ScheduledOrderFactory) GetOrderStatus(orderID int) string {
	if order, exists := sof.Orders[orderID]; exists {
		return order.Status
	}
	return ""
}

func NewOrderFactory(facType string) managers.OrderManager {
	switch facType {
	case "now":
		return &NowOrderFactory{
			Orders: make(map[int]models.Order),
		}
	case "scheduled":
		return &ScheduledOrderFactory{
			Orders: make(map[int]models.Order),
		}
	default:
		return nil
	}
}
