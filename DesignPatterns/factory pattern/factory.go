package main

type Burger interface {
	Cook() string
}

type chickenBurger struct{}

func (c *chickenBurger) Cook() string {
	return "Cooking Chicken Biryani"
}

type muttonBurger struct{}

func (m *muttonBurger) Cook() string {
	return "Cooking Mutton Biryani"
}

type vegBurger struct{}

func (v *vegBurger) Cook() string {
	return "Cooking Veg Biryani"
}

type beefBurger struct{}

func (b *beefBurger) Cook() string {
	return "Cooking Beef Biryani"
}

type italianBurger struct{}

func (i *italianBurger) Cook() string {
	return "Cooking Italian Biryani"
}

type BurgerFactory interface {
	createBurger(burgerType string) Burger
}

type AlluBurgerFactory struct{}

func (f *AlluBurgerFactory) createBurger(burgerType string) Burger {
	switch burgerType {
	case "chicken":
		return &chickenBurger{}
	case "mutton":
		return &muttonBurger{}
	case "veg":
		return &vegBurger{}
	default:
		return nil
	}
}

type MohanBurgerFactory struct{}

func (f *MohanBurgerFactory) createBurger(burgerType string) Burger {
	switch burgerType {
	case "beef":
		return &beefBurger{}
	case "italian":
		return &italianBurger{}
	default:
		return nil
	}
}

func Factory() {
	var factory BurgerFactory

	factory = &AlluBurgerFactory{}

	chickenBurger := factory.createBurger("chicken")
	println(chickenBurger.Cook())

	muttonBurger := factory.createBurger("mutton")
	println(muttonBurger.Cook())

	factory = &MohanBurgerFactory{}

	beefBurger := factory.createBurger("beef")
	println(beefBurger.Cook())

	italianBurger := factory.createBurger("italian")
	println(italianBurger.Cook())

}
