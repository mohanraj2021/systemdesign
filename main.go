package main

import (
	"fmt"

	"github.com/zomatodesign/app"
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
	appUser := application.UserManager.CreateUser(&user)
	restaurant := application.RestaurantManager.GetRestaurant("McDonalds")

	order := application.OrderManager.PlaceOrder(&models.Order{
		User:        appUser,
		Restaurant:  application.RestaurantManager.GetRestaurant("McDonalds"),
		Items:       []models.Menu{{Name: "McRib", Type: "Burger", Price: 5.99}},
		TotalAmount: 5.99,
	}, &models.Cart{}, &restaurant)

	fmt.Println(application.OrderManager.GetOrderStatus(order.Id))
}
