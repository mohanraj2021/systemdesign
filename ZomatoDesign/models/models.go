package models

type Menu struct {
	Name  string
	Type  string
	Price float32
}

type Restaurant struct {
	Id    int
	Name  string
	Menus []Menu
}

type User struct {
	Id    int
	Name  string
	Email string
}

type Order struct {
	Id          int
	User        User
	Restaurant  Restaurant
	Items       []Menu
	TotalAmount float32
	Status      string
}

type Cart struct {
	User  User
	Items []Menu
}

func (c *Cart) CreateCart(user User) {
	c.User = user
	c.Items = []Menu{}
}

func (c *Cart) AddItem(item Menu) {
	c.Items = append(c.Items, item)
}

func (c *Cart) RemoveItem(item Menu) {
	for i, v := range c.Items {
		if v.Name == item.Name {
			c.Items = append(c.Items[:i], c.Items[i+1:]...)
		}
	}
}

func (c *Cart) GetTotal() float32 {
	var total float32
	for _, item := range c.Items {
		total += item.Price
	}
	return total
}

func (c *Cart) ClearCart() {
	c.Items = []Menu{}
}
