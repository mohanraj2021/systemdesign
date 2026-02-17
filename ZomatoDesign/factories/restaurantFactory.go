package factories

import (
	"sync"

	"github.com/zomatodesign/managers"
	"github.com/zomatodesign/models"
)

var once sync.Once

type RestaurantFactory struct {
	restaurantMap map[string]models.Restaurant
}

func (rf *RestaurantFactory) GetRestaurant(name string) models.Restaurant {
	return rf.restaurantMap[name]
}
func (rf *RestaurantFactory) SetRestaurant(name string, restaurant models.Restaurant) models.Restaurant {
	rf.restaurantMap[name] = restaurant
	return restaurant
}
func (rf *RestaurantFactory) UpdateRestaurant(name string, restaurant models.Restaurant) models.Restaurant {
	rf.restaurantMap[name] = restaurant
	return restaurant
}
func (rf *RestaurantFactory) DeleteRestaurant(name string) bool {
	delete(rf.restaurantMap, name)
	return true
}
func (rf *RestaurantFactory) SearchRestaurant(name string) []models.Restaurant {
	var results []models.Restaurant
	for key, restaurant := range rf.restaurantMap {
		if key == name {
			results = append(results, restaurant)
		}
	}
	return results
}

func NewRestaurantFactory() managers.RestaurantManager {
	var restaurantFactory managers.RestaurantManager
	once.Do(func() {
		restaurantFactory = &RestaurantFactory{
			restaurantMap: make(map[string]models.Restaurant),
		}
	})
	return restaurantFactory
}
