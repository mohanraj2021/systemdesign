package models

type Parking struct {
	Level        int
	ParkingSpots []ParkingSpot
}

type ParkingSpot struct {
	SpotNumber int
	IsOccupied bool
	SlotType   string
	Level      int
}

func (p *Parking) TotalSlots() int {
	return len(p.ParkingSpots)
}
