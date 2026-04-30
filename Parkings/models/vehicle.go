package models

type Vehicle struct {
	VehicleType   string
	VehicleNumber string
	Size          string
	Is_Parked     bool
}

func (v *Vehicle) VehicleSize() string {
	return v.Size
}

func (v *Vehicle) Type() string {
	return v.VehicleType
}
