package managers

// import "github.com/zomatodesign/models"

// // OrderManager - Acceptable as is, but could be split further if needed
// type OrderManager interface {
// 	PlaceOrder(order *models.Order, cart *models.Cart, Restaurant *models.Restaurant) models.Order
// 	CancelOrder(orderID int) bool
// 	GetOrderStatus(orderID int) string
// }

// // Split RestaurantManager into read and write interfaces
// type RestaurantReader interface {
// 	GetRestaurant(name string) models.Restaurant
// 	SearchRestaurant(name string) []models.Restaurant
// }

// type RestaurantWriter interface {
// 	SetRestaurant(name string, restaurant models.Restaurant) models.Restaurant
// 	UpdateRestaurant(name string, restaurant models.Restaurant) models.Restaurant
// 	DeleteRestaurant(name string) bool
// }

// // Composite interface for full functionality
// type RestaurantManager interface {
// 	RestaurantReader
// 	RestaurantWriter
// }

// // Split UserManager into read and write interfaces
// type UserReader interface {
// 	GetUser(id int) models.User
// }

// type UserWriter interface {
// 	CreateUser(user *models.User) models.User
// 	UpdateUser(id int, user *models.User) models.User
// 	DeleteUser(id int) bool
// }

// // Composite interface for full functionality
// type UserManager interface {
// 	UserReader
// 	UserWriter
// }

// // NotificationService - Already follows ISP (single method)
// type NotificationService interface {
// 	Notify(user models.User, message string)
// }
