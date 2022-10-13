package main

import "fmt"

func main() {
	var username string = "CHitranjan"
	fmt.Println(username)
	fmt.Printf("Username is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("isLoggedIn is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255 //we cannot assign smallVal more than 255, it will show error because 256 will be out of its capacity.
	// the set of all unsigned 8-bit integers can be of (0 to 255)
	fmt.Println(smallVal)
	fmt.Printf("smallVal is of type: %T \n", smallVal)

	// default values and some aliases
	var anothervariable int
	fmt.Println(anothervariable)
	fmt.Printf("anotherVariable is of type: %T \n", anothervariable)

	var kuchBhi string
	fmt.Println(kuchBhi)
	fmt.Printf("KuchBhi is of type: %T \n", kuchBhi)

	// implicit way of declearing variables

	var website = "facebook.com"
	fmt.Println(website)

	// no var style
	// this style is only valid in any method only, outside of any method or if trying to make global it will throw an Error.
	numberOfUser := 300000
	fmt.Println(numberOfUser)

}
