package techpalace

import "strings"

const border = "*"
const newline = "\n"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return "Welcome to the Tech Palace, "+strings.ToUpper(customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	borderLine := ""
	for i := 0; i < numStarsPerLine; i++ {
		borderLine += border	
	}
	return borderLine + newline + welcomeMsg + newline + borderLine
//  stars := strings.Repeat("*", numStarsPerLine)
//	return stars + "\n" + welcomeMsg + "\n" + stars
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.TrimSpace(strings.Replace(oldMsg, border, "", -1))
}
