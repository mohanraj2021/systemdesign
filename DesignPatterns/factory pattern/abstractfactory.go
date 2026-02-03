package main

type Pizza interface {
	Cook() string
}

type GarlicBread interface {
	CookGarlicBread() string
}

type chickenPizza struct{}

func (c *chickenPizza) Cook() string {
	return "Cooking Chicken Biryani"
}

type muttonPizza struct{}

func (m *muttonPizza) Cook() string {
	return "Cooking Mutton Biryani"
}

type vegPizza struct{}

func (v *vegPizza) Cook() string {
	return "Cooking Veg Biryani"
}

type beefPizza struct{}

func (b *beefPizza) Cook() string {
	return "Cooking Beef Biryani"
}

type italianPizza struct{}

func (i *italianPizza) Cook() string {
	return "Cooking Italian Biryani"
}

type PizzaFactory interface {
	createPizza(PizzaType string) Pizza
}

type GarlicBreadFactory interface {
	createGarlicBread(garlicBreadType string) GarlicBread
}

type garlicBread struct{}

func (g *garlicBread) CookGarlicBread() string {
	return "Cooking Garlic Bread"
}

type AlluPizzaFactory struct{}

func (f *AlluPizzaFactory) createGarlicBread(GarlicBreadType string) GarlicBread {
	switch GarlicBreadType {
	case "normal":
		return &garlicBread{}
	default:
		return nil
	}
}
func (f *AlluPizzaFactory) createPizza(PizzaType string) Pizza {
	switch PizzaType {
	case "chicken":
		return &chickenPizza{}
	case "mutton":
		return &muttonPizza{}
	case "veg":
		return &vegPizza{}
	default:
		return nil
	}
}

type MohanPizzaFactory struct{}

func (f *MohanPizzaFactory) createPizza(PizzaType string) Pizza {
	switch PizzaType {
	case "beef":
		return &beefPizza{}
	case "italian":
		return &italianPizza{}
	default:
		return nil
	}
}

func AbstractFactory() {
	// var factory PizzaFactory

	factory := &AlluPizzaFactory{}
	// garlicBreadFactory := factory.(GarlicBreadFactory)

	chickenPizza := factory.createPizza("chicken")
	println(chickenPizza.Cook())

	muttonPizza := factory.createPizza("mutton")
	println(muttonPizza.Cook())

	garlicBread := factory.createGarlicBread("normal")
	println(garlicBread.CookGarlicBread())

	mfactory := &MohanPizzaFactory{}

	beefPizza := mfactory.createPizza("beef")
	println(beefPizza.Cook())

	italianPizza := mfactory.createPizza("italian")
	println(italianPizza.Cook())

}
