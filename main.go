package main

import (
	"GoBasicsPractical/Models"
	"GoBasicsPractical/Utilities"
	"bufio"
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

func messageError(msg string, err error) {
	if err != nil {
		fmt.Println(fmt.Sprint(msg, err))
	}
}

func main() {
	var (
		value   string
		err     error
		cardNum int
		card    Models.Card
		ok      bool
	)
	in := bufio.NewReader(os.Stdin)
	initialize()
	fmt.Println("Welcome to ATM: ")
	fmt.Println("Please Insert Your Card: ")
	for {

		//fmt.Scan(&cardNum)
		value, err = in.ReadString('\n')
		checkError(err)
		value = strings.TrimSpace(value)
		cardNum, err = strconv.Atoi(value)
		messageError("Invalid Card Reason:", err)
		card, ok = cards[cardNum]
		if ok {
			break
		}
		fmt.Println("Invalid Card Please insert Valid Card")
	}
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			var (
				choice string
			)
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
