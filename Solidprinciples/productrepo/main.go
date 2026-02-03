package productrepo

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

type ProdcutRepo interface {
	AddProduct(product Product)
	GetProduct(name string) (Product, bool)
}

type ProductRepository struct {
	Products map[string]Product
}

func (r *ProductRepository) AddProduct(product Product) {
	r.Products[product.Name] = product
}

func (r *ProductRepository) GetProduct(name string) (Product, bool) {
	if product, ok := r.Products[name]; ok {
		return product, true
	}
	return Product{}, false
}

func NewProductRepository() ProdcutRepo {
	return &ProductRepository{
		Products: make(map[string]Product, 0),
	}
}

type OrderService struct {
	productRepo ProdcutRepo
}

func NewOrderService(productRepo ProdcutRepo) OrderService {
	return OrderService{productRepo: productRepo}
}

func (o *OrderService) PlaceOrder(productName string) {
	if product, found := o.productRepo.GetProduct(productName); found && product.Quantity > 0 {
		println("Order placed for product:", product.Name, "with price:", product.Price)
	} else {
		println("Product not found:", productName)
	}
}
func main() {
	productRepo := NewProductRepository()
	oService := NewOrderService(productRepo)
	productRepo.AddProduct(Product{Name: "Laptop", Price: 999.99})
	productRepo.AddProduct(Product{Name: "Smartphone", Price: 499.99})
	var product Product
	if product, found := productRepo.GetProduct("Laptop"); found {
		println("Product found:", product.Name, "Price:", product.Price)
	} else {
		println("Product not found")
	}

	oService.PlaceOrder(product.Name)
}
