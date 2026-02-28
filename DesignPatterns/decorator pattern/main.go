package main

import "fmt"

type ICharacter interface {
	getPowerUp() int
}

type Character struct {
	Name  string
	Age   int
	Power int
}

func (c *Character) getPowerUp() int {
	c.Power = c.Power + 1
	fmt.Printf("Base %s Power: %d \n", c.Name, c.Power)
	return c.Power
}

type FirePowerUp struct {
	Character ICharacter
}

func (f *FirePowerUp) getPowerUp() int {
	basePower := f.Character.getPowerUp()
	return basePower + 2
}

type WaterPowerUp struct {
	Character ICharacter
}

func (w *WaterPowerUp) getPowerUp() int {
	basePower := w.Character.getPowerUp()
	return basePower + 3
}

type EarthPowerUp struct {
	Character ICharacter
}

func (e *EarthPowerUp) getPowerUp() int {
	basePower := e.Character.getPowerUp()
	return basePower + 4
}

func main() {
	// Base character
	character := &Character{Name: "John", Age: 30, Power: 0}
	character2 := &Character{Name: "Alice", Age: 25, Power: 0}
	// Wrap base with fire decorator
	firePowerUp := &FirePowerUp{Character: character}
	firePowerUp2 := &FirePowerUp{Character: character2}

	// Wrap fire with water decorator (chaining)
	waterPowerUp := &WaterPowerUp{Character: firePowerUp}
	waterPowerUp2 := &WaterPowerUp{Character: firePowerUp2}

	earthPowerUp := &EarthPowerUp{Character: waterPowerUp}
	earthPowerUp2 := &EarthPowerUp{Character: waterPowerUp2}
	fmt.Println("Earth + Water + Fire Power Total:", earthPowerUp.getPowerUp())
	fmt.Println("Earth + Water + Fire Power Total:", earthPowerUp2.getPowerUp())
}
