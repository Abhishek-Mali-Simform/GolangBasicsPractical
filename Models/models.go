package Models

import (
	"errors"
	"fmt"
	"math"
)

type UserDetails struct {
	username    string
	accountType string
	contact     int
	balance     float64
}

type Card struct {
	number      int
	pin         int
	userDetails UserDetails
}

func checkLimit(num, limit int, err string) error {
	if int(math.Log10(float64(num))) != limit {
		return errors.New("invalid " + err)
	}
	return nil
}

func (userDetails *UserDetails) SetUsername(username string) error {
	if username == "" {
		return errors.New("username required")
	}
	userDetails.username = username
	return nil
}

func (userDetails *UserDetails) Username() string {
	return userDetails.username
}

func (userDetails *UserDetails) SetAccountType(accountType string) error {
	if accountType == "" {
		return errors.New("account type required")
	}
	flag := 0
	switch accountType {
	case "savings", "current":
		flag = 0
	default:
		flag = 1
	}
	if flag == 1 {
		return errors.New("invalid account type")
	}
	//else if (strings.Compare(accountType, "current") != 0) || (strings.Compare(accountType, "savings") != 0) {
	//	return errors.New("invalid account type")
	//}
	userDetails.accountType = accountType
	return nil
}

func (userDetails *UserDetails) AccountType() string {
	return userDetails.accountType
}

func (userDetails *UserDetails) SetContact(contact int) error {
	err := checkLimit(contact, 9, "contact number")
	if err != nil {
		return err
	}
	userDetails.contact = contact
	return nil
}

func (userDetails *UserDetails) Contact() int {
	return userDetails.contact
}

func (userDetails *UserDetails) SetBalance(balance float64) error {
	if balance < 0 {
		return errors.New("invalid contact number")
	}
	userDetails.balance = balance
	return nil
}

func (userDetails *UserDetails) Balance() float64 {
	return userDetails.balance
}

func (card *Card) SetNumber(number int) error {
	err := checkLimit(number, 5, "card number")
	if err != nil {
		return err
	}
	card.number = number
	return nil
}

func (card *Card) Number() int {
	return card.number
}

func (card *Card) SetPin(pin int) error {
	err := checkLimit(pin, 3, "pin")
	if err != nil {
		return err
	}
	card.pin = pin
	return nil
}

func (card *Card) Pin() int {
	return card.pin
}

func (card *Card) SetUserDetails(userDetails UserDetails) error {
	if (userDetails == UserDetails{}) {
		return errors.New("no user details to save")
	}
	card.userDetails = userDetails
	return nil
}

func (card *Card) UserDetails() UserDetails {
	return card.userDetails
}

func (userDetails *UserDetails) String() string {
	return fmt.Sprintf("Username: %s Account Type: %s Contact: %d Balance: %.2f", userDetails.Username(), userDetails.AccountType(), userDetails.Contact(), userDetails.Balance())
}

func (card *Card) String() string {
	userDetail := card.UserDetails()
	return fmt.Sprintf("Card Number: %v Pin: %v User Details: %v", card.Number(), card.Pin(), userDetail)
}
