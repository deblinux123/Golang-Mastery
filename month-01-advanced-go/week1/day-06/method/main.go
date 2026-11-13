package main

import (
	"fmt"
	"log"
)

type Product struct {
	Name  string
	Price float64
}

type BankAccount struct {
	Owner   string
	Balance float64
}

func (b *BankAccount) Display() string {
	return fmt.Sprintf("Owner: %s\nBalance:%.2f\n", b.Owner, b.Balance)
}

func (b *BankAccount) GetBalance() float64 {
	return b.Balance
}

func (b *BankAccount) CanWithdraw(amount float64) (bool, error) {
	if amount > b.Balance || amount <= 0 {
		return false, fmt.Errorf("Invlaid amount")
	}

	b.Balance -= amount

	return true, nil
}

func (p *Product) Display() {
	fmt.Println("Product:", p.Name)
	fmt.Println("Price:", p.Price)
}

func main() {

	product := Product{
		Name:  "Laptop",
		Price: 1200,
	}

	product.Display()

	account := BankAccount{
		Owner:   "Farid",
		Balance: 100,
	}

	fmt.Println(account.Display())
	fmt.Println(account.GetBalance())

	with, err := account.CanWithdraw(10)

	if err != nil {
		log.Println(err)
	}

	if with {
		log.Println("Withdraw successfully.")
	}
	fmt.Println("After withdraw:", account.GetBalance())
}
