package Utilities

import (
	"GoBasicsPractical/Models"
	"fmt"
	"log"
	"os"
	"strings"
)

func checkNegative(num float64) {
	if num < 0 {
		panic("Negative Amount")
	}
}

func checkLimit(num float64) {
	if num > 25000 {
		panic("Exceeded Limit")
	}
}

func checkBalance(balance, amount float64) {
	if amount > balance {
		panic("Insufficient Balance")
	}
}

func Banking(card *Models.Card) {
	var (
		pin         int
		accountType string
	)
	fmt.Print("Enter Your Pin: ")
	fmt.Scan(&pin)
	if pin != card.Pin() {
		fmt.Println("Invalid Pin")
		os.Exit(0)
	}
AGAIN:
	fmt.Println("Enter Account Type: ")
	fmt.Scan(&accountType)
	userDetails := card.UserDetails()
	if !strings.EqualFold(accountType, userDetails.AccountType()) {
		goto AGAIN
	}
	BankOperations(card)
}

func BankOperations(card *Models.Card) {
	userDetails := card.UserDetails()
	var (
		choice int
		amount float64
	)
	fmt.Println("-----Welcome ", userDetails.Username(), " -----")
	fmt.Println("1. Withdraw Money")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Check Balance")
	fmt.Println("Enter Your Choice")
	fmt.Scan(&choice)
	switch choice {
	case 1:
		fmt.Println("Enter Amount to be Withdraw:")
		fmt.Scan(&amount)
		Withdraw(card, amount)
		amount = CheckBalance(card)
	case 2:
		fmt.Println("Enter Amount to be Deposited:")
		fmt.Scan(&amount)
		Deposit(card, amount)
		amount = CheckBalance(card)
	case 3:
		amount = CheckBalance(card)
	default:
		log.Fatal("other features not available")
	}
	fmt.Println("Hello ", userDetails.Username(), ", Your balance: ", amount)
}

func Deposit(card *Models.Card, amount float64) {
	checkNegative(amount)
	userDetails := card.UserDetails()
	userDetails.SetBalance(userDetails.Balance() + amount)
	card.SetUserDetails(userDetails)
}

func Withdraw(card *Models.Card, amount float64) {
	checkNegative(amount)
	checkLimit(amount)
	userDetails := card.UserDetails()
	balance := userDetails.Balance()
	checkBalance(balance, amount)
	userDetails.SetBalance(balance - amount)
	card.SetUserDetails(userDetails)
}

func CheckBalance(card *Models.Card) float64 {
	userDetails := card.UserDetails()
	return userDetails.Balance()
}
