package Utilities

import (
	"GoBasicsPractical/Models"
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

func checkError(msg string, err error) {
	if err != nil {
		fmt.Println(msg, " Reason: ", err)
	}
}

func checkErrorExit(msg string, err error) {
	if err != nil {
		fmt.Println(msg, " Reason: ", err)
		os.Exit(0)
	}
}

func input(msg string) (value string) {
	var err error
	value, err = in.ReadString('\n')
	checkError(msg, err)
	value = strings.TrimSpace(value)
	return
}

func checkNegative(num int) {
	if num < 0 {
		panic("Negative Amount")
	}
}

func checkAmount(num int) {
	if int(num)%10 == 0 || int(num)%50 == 0 {
		return
	}
	panic("Invalid Amount Please Enter Amount in multiples of 10 or 50")
}

func checkLimit(num int) {
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
		value       string
		err         error
	)
	fmt.Print("Enter Your Pin: ")
	value = input("Invalid Pin")
	pin, err = strconv.Atoi(value)
	checkErrorExit("Invalid Pin", err)
	if pin != card.Pin() {
		fmt.Println("Invalid Pin")
		os.Exit(0)
	}
	fmt.Println("Enter Account Type: ")
	accountType = input("Invalid Account Type")
	userDetails := card.UserDetails()
	if !strings.EqualFold(accountType, userDetails.AccountType()) {
		fmt.Println("Invalid Account Type")
		os.Exit(0)
	}
	BankOperations(card)
}

func BankOperations(card *Models.Card) {
	userDetails := card.UserDetails()
	var (
		choice  int
		amount  int
		balance float64
		value   string
		err     error
		yesNo   string
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
		value = input("Invalid Amount")
		amount, err = strconv.Atoi(value)
		checkErrorExit("Invalid Amount. ", err)
		Withdraw(card, amount)
		balance = CheckBalance(card)
	case 2:
		fmt.Println("Enter Amount to be Deposited:")
		value = input("Invalid Amount")
		amount, err = strconv.Atoi(value)
		checkErrorExit("Invalid Amount. ", err)
		Deposit(card, amount)
		balance = CheckBalance(card)
	case 3:
		balance = CheckBalance(card)
		fmt.Println("Do You Wish To Conintue(yes to continue): ")
		yesNo = input("Invalid choice.")
	default:
		log.Fatal("other features not available")
		fmt.Println("Do You Wish To Conintue(yes to continue): ")
		yesNo = input("Invalid choice.")
	}
	fmt.Println("Hello ", userDetails.Username(), ", Your balance: ", balance)
	if strings.EqualFold(yesNo, "yes") {
		BankOperations(card)
	}
}

func Deposit(card *Models.Card, amount int) {
	checkNegative(amount)
	checkLimit(amount)
	userDetails := card.UserDetails()
	userDetails.SetBalance(userDetails.Balance() + float64(amount))
	card.SetUserDetails(userDetails)
}

func Withdraw(card *Models.Card, amount int) {
	checkNegative(amount)
	checkLimit(amount)
	checkAmount(amount)
	userDetails := card.UserDetails()
	balance := userDetails.Balance()
	checkBalance(balance, float64(amount))
	userDetails.SetBalance(balance - float64(amount))
	card.SetUserDetails(userDetails)
}

func CheckBalance(card *Models.Card) float64 {
	userDetails := card.UserDetails()
	return userDetails.Balance()
}
