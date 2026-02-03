package app

import (
	"github.com/zomatodesign/factories"
	"github.com/zomatodesign/managers"
	"github.com/zomatodesign/models"
)

type App struct {
	RestaurantManager managers.RestaurantManager
	OrderManager      managers.OrderManager
	UserManager       managers.UserManager
}

func NewApp() *App {
	return &App{
		RestaurantManager: *factories.NewRestaurantFactory(),
		OrderManager:      factories.NewOrderFactory("now"),
		UserManager:       factories.NewUserFactory(),
	}
}

func (a *App) InitializeRestaurants() {
	a.RestaurantManager.SetRestaurant("McDonalds", models.Restaurant{
		Id:    1,
		Name:  "McDonalds",
		Menus: []models.Menu{{Name: "McRib", Type: "Burger", Price: 5.99}},
	})
	a.RestaurantManager.SetRestaurant("BurgerKing", models.Restaurant{
		Id:    2,
		Name:  "BurgerKing",
		Menus: []models.Menu{{Name: "Big Mac", Type: "Burger", Price: 5.99}},
	})
	a.RestaurantManager.SetRestaurant("KFC", models.Restaurant{
		Id:    3,
		Name:  "KFC",
		Menus: []models.Menu{{Name: "Chicken", Type: "Burger", Price: 5.99}},
	})

}
