package main

import (
	"fmt"
)

type Prodct struct {
	ID       int
	Name     string
	Price    float64
	Stock    int
	IsActive bool
}

type Inventory struct {
	Products map[int]*Prodct
}

func (i *Inventory) AddProduct(p Prodct) error {

	if p.ID <= 0 {
		return fmt.Errorf("invalid product ID: %d", p.ID)
	}

	if p.Name == "" {
		return fmt.Errorf("Name connot be empty")
	}

	if p.Price <= 0 {
		return fmt.Errorf("Price must grater than zero and cannot be negative")
	}

	if p.Stock < 0 {
		return fmt.Errorf("stock cannot be empty or negative.")
	}

	if _, ok := i.Products[p.ID]; ok {
		return fmt.Errorf("product with ID %d already exists", p.ID)
	}

	i.Products[p.ID] = &p

	return nil

}

func (i *Inventory) GetProduct(id int) (*Prodct, error) {
	product, ok := i.Products[id]

	if !ok {
		return nil, fmt.Errorf("product with id %d not found", id)
	}

	return product, nil
}

func (i *Inventory) UpdateStock(id int, quantity int) error {

	if quantity < 0 {
		return fmt.Errorf("stock cannot be negative")
	}

	product, err := i.GetProduct(id)

	if err != nil {
		return err
	}

	(*product).Stock = quantity

	return nil
}

func (i *Inventory) DeleteProduct(id int) error {
	if id <= 0 {
		return fmt.Errorf("Invalid id: %d", id)
	}

	if _, ok := i.Products[id]; !ok {
		return fmt.Errorf("product with id %d not found.", id)
	}

	delete(i.Products, id)

	return nil
}

func (i *Inventory) TotalValue() float64 {
	total := 0.0

	for _, product := range i.Products {
		total += product.Price * float64(product.Stock)
	}

	return total
}

func (i *Inventory) ListAllProduct() {

	if len(i.Products) == 0 {
		fmt.Println("Inventory is empty.")
		return
	}

	for id, product := range i.Products {
		fmt.Printf(
			"ID: %d | Name: %s | Price: %.2f | Stock: %d\n",
			id,
			product.Name,
			product.Price,
			product.Stock,
		)
	}
}

func NewInventory() *Inventory {
	return &Inventory{
		Products: make(map[int]*Prodct),
	}
}
func main() {
	inventory := NewInventory()

	products := []Prodct{
		{ID: 1, Name: "Laptop", Price: 123.33, Stock: 43, IsActive: true},
		{ID: 2, Name: "Mouse", Price: 25, Stock: 20, IsActive: true},
		{ID: 3, Name: "Keyboard", Price: 75, Stock: 10, IsActive: true},
	}

	newProduct := Prodct{
		ID:    2,
		Name:  "Mac Book",
		Price: 1200,
		Stock: 5,
	}

	for _, items := range products {
		inventory.AddProduct(items)
	}

	fmt.Println("Result:")
	inventory.ListAllProduct()

	err := inventory.AddProduct(newProduct)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Product added.")
	}

	pro, err := inventory.GetProduct(4)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(*pro)

	if err := inventory.UpdateStock(4, 100); err != nil {

		fmt.Println("Error:", err)
	} else {
		fmt.Println("Stocke updated successfully.")
	}
	inventory.ListAllProduct()

	err = inventory.DeleteProduct(-48)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Item deleted.")
	}

	inventory.ListAllProduct()

	total := inventory.TotalValue()

	fmt.Println("Total Valuel:", total)
}
