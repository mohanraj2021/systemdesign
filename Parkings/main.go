package main

import (
	"github.com/parking/managers"
	"github.com/parking/models"
)

type App struct {
	VehicleManage managers.IVehicle
	ParkingManage managers.Iparking
}

func NewApp() *App {
	return &App{
		VehicleManage: &managers.VehicleManager{Vehicles: make(map[string]models.Vehicle)},
		ParkingManage: &managers.ParkingManager{ParkingSpots: []models.Parking{}},
	}
}
func main() {

	app := NewApp()
	vehicle := models.Vehicle{
		VehicleType:   "Car",
		VehicleNumber: "KA-01-AB-1234",
		Size:          "Medium",
		Is_Parked:     false,
	}
	err := app.VehicleManage.AddVehicle(vehicle)
	if err != nil {
		panic(err)
	}
	parkingLot1 := models.Parking{
		Level: 1,
		ParkingSpots: []models.ParkingSpot{
			{SpotNumber: 1, IsOccupied: false, SlotType: "Medium", Level: 1},
			{SpotNumber: 2, IsOccupied: false, SlotType: "Large", Level: 1},
			{SpotNumber: 3, IsOccupied: false, SlotType: "Small", Level: 1},
		},
	}
	parkingLot2 := models.Parking{
		Level: 2,
		ParkingSpots: []models.ParkingSpot{
			{SpotNumber: 1, IsOccupied: false, SlotType: "Medium", Level: 2},
			{SpotNumber: 2, IsOccupied: false, SlotType: "Large", Level: 2},
			{SpotNumber: 3, IsOccupied: false, SlotType: "Small", Level: 2},
		},
	}
	app.ParkingManage.AddParkingLevel(parkingLot1)
	app.ParkingManage.AddParkingLevel(parkingLot2)
	app.ParkingManage.ParkVehicle(&vehicle)

}
