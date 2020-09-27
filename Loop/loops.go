package main

import "fmt"

func main() {
	sum := 0
	// Basic for, declared variable for control(init), condition evaluated at each pass
	// and, in this case, increment of the variable(post)
	// Note that it doesn't use parentheses on the three components of the for
	// Braces are always required
	// This can be declared without the init and post statements(don't know why you would do that but, you can)
	// Syntax is the same,' for ; i < number; { code } '
	fmt.Println("----Basic for----")
	for i := 0; i < 10; i++ {
		sum += i
		fmt.Println("Sum in the pass number,", i, "is:", sum)
	}

	// Go's while is declared as a for, no semicolon needed
	fmt.Println("----While----")
	for sum < 1000 {
		sum += sum
	}
	fmt.Println("Total value of the sum:", sum)

	// For's can loop forever, just omit the parameters and conditions
	fmt.Println("----This for would go forever, but there's a break condition inside it----")
	value := 0
	for {
		value = value + 1
		if value > 5 {
			fmt.Println("Reached the break point, value is: ", value)
			break
		}
	}

	// For using range, iterating through a array
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("----For using range to iterate over the items in a array of numbers----")
	for i := range numbers {
		fmt.Println("Position is:", i, "and number in this position of the array is:", numbers[i])
	}

	fmt.Println("----For using range and the keys to iterate----")
	capitalCountrys := map[string]string{"Alemanha": "Berlin", "Irlanda": "Dublin", "Italia": "Roma"}
	for country := range capitalCountrys {
		fmt.Println("Capital of", country, "is", capitalCountrys[country])
	}

	fmt.Println("----For using range and the key-value to iterate----")
	for country, capital := range capitalCountrys {
		fmt.Println("Capital of", country, "is", capital)
	}
}
