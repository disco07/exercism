package techpalace

import (
	"fmt"
	"strings"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var stars []string
	for i := 0; i < numStarsPerLine; i++ {
		stars = append(stars, "*")
	}
	affich := strings.Join(stars, "") + "\n" + welcomeMsg + "\n" + strings.Join(stars, "")

	return affich
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	newMsg := strings.Replace(oldMsg, "*", "", -1)
	return strings.TrimSpace(newMsg)
}
