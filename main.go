package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Println(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	opt, _ := getInput("Choose option (a - add item, s - save bill, t - add tip): ", reader)

	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)

		p, err := strconv.ParseFloat(price, 64)

		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}

		b.addItem(name, p)

		fmt.Println("Item added")
		promptOptions(b)
	case "s":
		b.save()
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)

		parsedTip, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}

		b.updateTip(parsedTip)
		fmt.Println("Tip added")
		promptOptions(b)
	default:
		fmt.Println("invalid option selected...")
		promptOptions(b)
	}
}

func main() {
	myBill := createBill()
	promptOptions(myBill)

	fmt.Println(myBill)
	// myBill := newBill("maej's bill")

	// myBill.addItem("fish", 4.99)
	// myBill.addItem("pepper", 3.40)
	// myBill.addItem("water", 5.50)
	// myBill.addItem("coffee", 4.10)

	// myBill.updateTip(2.99)

	// fmt.Println(myBill.format())
}
