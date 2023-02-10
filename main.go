package main

import (
	"GoBasicsPractical/Models"
	"GoBasicsPractical/Utilities"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	err         error
	card        = Models.Card{}
	userDetails = Models.UserDetails{}
	cards       = map[int]Models.Card{}
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	var cardNum int
	if len(os.Args) < 2 {
		fmt.Println("usage: main.go [card-number]")
		os.Exit(0)
	}
	args := os.Args[1:2]
	if len(args) != 1 {
		fmt.Println("usage: main.go [card-number]")
		os.Exit(0)
	}
	cardNum, _ = strconv.Atoi(args[0])
	if cardNum <= 0 {
		fmt.Println("Invalid Card")
		os.Exit(0)
	}
	initialize()
	card, ok := cards[cardNum]
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			var choice string
			fmt.Println("Do You Wish To Conintue(yes to continue): ")
			fmt.Scan(&choice)
			if strings.EqualFold(choice, "yes") {
				Utilities.BankOperations(&card)
			} else {
				fmt.Println("Exiting the System, you can remove your card...")
				os.Exit(0)
			}
		}
	}()
	if ok {
		Utilities.Banking(&card)
		fmt.Println("Exiting the System, you can remove your card...")
	} else {
		fmt.Println("Invalid Card...")
		os.Exit(0)
	}
}

func initialize() {
	err = userDetails.SetUsername("Abhishek Mali")
	checkError(err)
	err = userDetails.SetAccountType("savings")
	checkError(err)
	err = userDetails.SetContact(9429865212)
	checkError(err)
	err = userDetails.SetBalance(11000)
	checkError(err)
	err = card.SetNumber(398653)
	checkError(err)
	err = card.SetPin(1000)
	checkError(err)
	err = card.SetUserDetails(userDetails)
	checkError(err)
	cards[card.Number()] = card
	err = userDetails.SetUsername("Kishan Mehta")
	checkError(err)
	err = userDetails.SetAccountType("savings")
	checkError(err)
	err = userDetails.SetContact(9876543210)
	checkError(err)
	err = userDetails.SetBalance(5000)
	checkError(err)
	err = card.SetNumber(102023)
	checkError(err)
	err = card.SetPin(1002)
	checkError(err)
	err = card.SetUserDetails(userDetails)
	checkError(err)
	cards[card.Number()] = card
	err = userDetails.SetUsername("Heema Goratela")
	checkError(err)
	err = userDetails.SetAccountType("savings")
	checkError(err)
	err = userDetails.SetContact(1234567890)
	checkError(err)
	err = userDetails.SetBalance(20000)
	checkError(err)
	err = card.SetNumber(202310)
	checkError(err)
	err = card.SetPin(2100)
	checkError(err)
	err = card.SetUserDetails(userDetails)
	checkError(err)
	cards[card.Number()] = card
}
