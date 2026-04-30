package managers

import (
	"errors"

	"github.com/parking/models"
)

type IVehicle interface {
	AddVehicle(vehicle models.Vehicle) error
	EditVehicle(vehicle models.Vehicle) error
	DeleteVehicle(vehicle models.Vehicle) error
}

type Iparking interface {
	IParkingManage
	ParkingLevelManagement
	ParkingCrud
}

type IParkingManage interface {
	ParkVehicle(*models.Vehicle) error
	UnparkVehicle(*models.Vehicle) error
}

type ParkingLevelManagement interface {
	AvailableSlots(*models.Vehicle) []models.ParkingSpot
	TotalSlots() int
}

type ParkingCrud interface {
	AddParkingLevel(models.Parking)
	RemoveParkingLevel(level int) error
	EditParkingLevel(level int, parking models.Parking) error
}

type ParkingManager struct {
	ParkingSpots []models.Parking
}

type VehicleManager struct {
	Vehicles map[string]models.Vehicle
}

func (vm *VehicleManager) AddVehicle(vehicle models.Vehicle) error {
	if _, ok := vm.Vehicles[vehicle.VehicleNumber]; ok {
		return errors.New("vehicle already exist")
	}
	vm.Vehicles[vehicle.VehicleNumber] = vehicle
	return nil
}

func (vm *VehicleManager) EditVehicle(vehicle models.Vehicle) error {
	if _, ok := vm.Vehicles[vehicle.VehicleNumber]; ok {
		vm.Vehicles[vehicle.VehicleNumber] = vehicle
		return nil
	}
	return errors.New("vehicle not found")
}

func (vm *VehicleManager) DeleteVehicle(vehicle models.Vehicle) error {
	delete(vm.Vehicles, vehicle.VehicleNumber)
	return nil
}

func (pm *ParkingManager) AvailableSlots(*models.Vehicle) []models.ParkingSpot {
	var availableSpots []models.ParkingSpot
	for _, level := range pm.ParkingSpots {
		for _, spot := range level.ParkingSpots {
			if !spot.IsOccupied {
				availableSpots = append(availableSpots, spot)
			}
		}
	}
	return availableSpots
}

func (pm *ParkingManager) ParkVehicle(vehicle *models.Vehicle) error {
	for i, level := range pm.ParkingSpots {
		for j, spot := range level.ParkingSpots {
			if !spot.IsOccupied {
				pm.ParkingSpots[i].ParkingSpots[j].IsOccupied = true
				vehicle.Is_Parked = true
				return nil
			}
		}
	}
	return errors.New("no available parking spots")
}

func (pm *ParkingManager) UnparkVehicle(vehicle *models.Vehicle) error {
	for i, level := range pm.ParkingSpots {
		for j, spot := range level.ParkingSpots {
			if spot.IsOccupied {
				pm.ParkingSpots[i].ParkingSpots[j].IsOccupied = false
				vehicle.Is_Parked = false
				return nil
			}
		}
	}
	return errors.New("vehicle not found in parking")
}

func (pm *ParkingManager) TotalSlots() int {
	return len(pm.ParkingSpots)
}

func (pm *ParkingManager) TotalAvailableSlots() int {
	count := 0
	for _, level := range pm.ParkingSpots {
		for _, spot := range level.ParkingSpots {
			if !spot.IsOccupied {
				count++
			}
		}
	}
	return count
}

func (pm *ParkingManager) AddParkingLevel(level models.Parking) {
	pm.ParkingSpots = append(pm.ParkingSpots, level)
}

func (pm *ParkingManager) RemoveParkingLevel(level int) error {
	for i, parking := range pm.ParkingSpots {
		if parking.Level == level {
			pm.ParkingSpots = append(pm.ParkingSpots[:i], pm.ParkingSpots[i+1:]...)
			return nil
		}
	}
	return errors.New("parking level not found")
}

func (pm *ParkingManager) EditParkingLevel(level int, parking models.Parking) error {
	for i, p := range pm.ParkingSpots {
		if p.Level == level {
			pm.ParkingSpots[i] = parking
			return nil
		}
	}
	return errors.New("parking level not found")
}

func NewParkingManager() Iparking {
	return &ParkingManager{
		ParkingSpots: make([]models.Parking, 0),
	}
}

func NewVehicleManager() IVehicle {
	return &VehicleManager{
		Vehicles: make(map[string]models.Vehicle, 0),
	}
}
