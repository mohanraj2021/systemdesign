package managers

import "github.com/zomatodesign/models"

type OrderManager interface {
	PlaceOrder(order *models.Order, cart *models.Cart, Restaurant *models.Restaurant) models.Order
	CancelOrder(orderID int) bool
	GetOrderStatus(orderID int) string
	// PayForOrder(orderID int,payType string) bool
}

type RestaurantManager interface {
	ResaurantReader
	ResaurantWriter
}

type ResaurantReader interface {
	GetRestaurant(name string) models.Restaurant
	SearchRestaurant(name string) []models.Restaurant
}

type ResaurantWriter interface {
	SetRestaurant(name string, restaurant models.Restaurant) models.Restaurant
	UpdateRestaurant(name string, restaurant models.Restaurant) models.Restaurant
	DeleteRestaurant(name string) bool
}

type UserManager interface {
	UserReader
	UserWriter
}

type UserReader interface {
	GetUser(id int) models.User
}

type UserWriter interface {
	CreateUser(user *models.User) models.User
	UpdateUser(id int, user *models.User) models.User
	DeleteUser(id int) bool
}

type NotificationService interface {
	Notify(user models.User, message string)
}
