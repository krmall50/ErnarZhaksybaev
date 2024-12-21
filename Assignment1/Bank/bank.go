package Bank

import (
	"fmt"
)

type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

func (ba *BankAccount) Deposit(amount float64) {
	if amount < 0 {
		fmt.Println("Amount must be positive")
		return
	}
	ba.Balance += amount
	fmt.Printf("Deposited: %.2f\n Balance: %.2f", amount, ba.Balance)
}

func (ba *BankAccount) Withdraw(amount float64) {
	if amount <= 0 {
		fmt.Println("Withdrawal amount must be positive.")
		return
	}
	if amount > ba.Balance {
		fmt.Println("Insufficient balance.")
		return
	}
	ba.Balance -= amount
	fmt.Printf("Withdrew: %.2f. New Balance: %.2f\n", amount, ba.Balance)
}

func (ba *BankAccount) GetBalance() {
	fmt.Printf("Current Balance: %.2f\n", ba.Balance)
}

func Transaction(account *BankAccount, transactions []float64) {
	for _, amount := range transactions {
		if amount > 0 {
			account.Deposit(amount)
		} else {
			account.Withdraw(-amount)
		}
	}
}
