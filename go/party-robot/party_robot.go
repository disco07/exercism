package partyrobot

import (
	"fmt"
)

// Welcome greets a person by name.
func Welcome(name string) string {
	return fmt.Sprintf("Welcome to my party, %s!", name)
}

// HappyBirthday wishes happy birthday to the birthday person and exclaims their age.
func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Happy birthday %s! You are now %v years old!", name, age)
}

// AssignTable assigns a table to each guest.
func AssignTable(name string, table int, neighbor, direction string, distance float64) string {
	t := intToStringMat(table)
	return Welcome(name) + "\n" + fmt.Sprintf("You have been assigned to table %v. Your table is %s, exactly %.1f meters from here.\nYou will be sitting next to %s.", t, direction, distance, neighbor)
}

func intToStringMat(number int) string {
	switch {
	case number < 10:
		return "00" + fmt.Sprintf("%v", number)
	case number >= 10 && number <= 100:
		return "0" + fmt.Sprintf("%v", number)
	default:
		return fmt.Sprintf("%v", number)
	}
}
