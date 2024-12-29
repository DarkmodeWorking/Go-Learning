package main

import "fmt"

// jwtToken := "jhegdyghe7813"
var jwtToken = "13y3eh23ey1ud"

const LoginToken string = "dhcyqeugdvyqd" // Capital letter declares Public

func main() {
	var username string = "Anurag"
	fmt.Println(username)
	fmt.Printf("Variable Type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable Type: %T \n", isLoggedIn)

	var smallValue uint8 = 240 // unsigned integer 8 bits
	fmt.Println(smallValue)
	fmt.Printf("Variable Type: %T \n", smallValue)

	var smallFloat float32 = 267.443532234224
	fmt.Println(smallFloat)
	fmt.Printf("Variable Type: %T \n", smallFloat)

	var bigFloat float64 = 267.443532234224
	fmt.Println(bigFloat)
	fmt.Printf("Variable Type: %T \n", bigFloat)

	// default value for variables
	var undeclared int
	fmt.Println(undeclared)
	fmt.Printf("Variable Type: %T \n", undeclared)

	// implicit types
	var website = "github.com/DarkmodeWorking"
	fmt.Println(website)

	// no var style
	numberOfCommits := 212
	fmt.Println(numberOfCommits)
	fmt.Printf("Variable Type: %T \n", numberOfCommits)

	fmt.Println(LoginToken)
	fmt.Printf("Variable Type: %T \n", LoginToken)
}
