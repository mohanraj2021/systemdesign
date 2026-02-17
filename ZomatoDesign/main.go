package main

import (
	"fmt"
	"math/rand/v2"

	"github.com/zomatodesign/app"
	"github.com/zomatodesign/factories"
	"github.com/zomatodesign/models"
)

func main() {
	application := app.NewApp()
	application.InitializeRestaurants()

	user := models.User{
		Id:    1,
		Name:  "Mohan",
		Email: "mohan@zomato.com",
	}

	user1 := models.User{
		Id:    2,
		Name:  "Alice",
		Email: "alice@zomato.com",
	}
	appUser := application.UserManager.CreateUser(&user)
	appUser1 := application.UserManager.CreateUser(&user1)
	restaurant := application.RestaurantManager.GetRestaurant("McDonalds")

	order := application.OrderManager.PlaceOrder(&models.Order{
		Id:          rand.Int(),
		User:        appUser,
		Restaurant:  application.RestaurantManager.GetRestaurant("McDonalds"),
		Items:       []models.Menu{{Name: "McRib", Type: "Burger", Price: 5.99}},
		TotalAmount: 5.99,
	}, &models.Cart{}, &restaurant)

	application.OrderManager = factories.NewOrderFactory("scheduled")

	order1 := application.OrderManager.PlaceOrder(&models.Order{
		Id:          rand.Int(),
		User:        appUser1,
		Restaurant:  application.RestaurantManager.GetRestaurant("McDonalds"),
		Items:       []models.Menu{{Name: "McRib", Type: "Burger", Price: 5.99}},
		TotalAmount: 5.99,
	}, &models.Cart{}, &restaurant)

	application.NotificationService.Notify(appUser, "Your order has been placed")

	application.NotificationService.Notify(appUser1, "Your order has been placed")
	fmt.Println(application.OrderManager.GetOrderStatus(order.Id))
	fmt.Println(application.OrderManager.GetOrderStatus(order1.Id))
}
