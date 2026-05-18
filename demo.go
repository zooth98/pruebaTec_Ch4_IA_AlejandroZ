package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	textSamples := []string{
		"My password is 123456",
		"My SSN is 555-11-9999",
		"Public weather report for tomorrow",
		"Credit card number 4111111111111111",
		"Internal company financial report",
	}

	fmt.Println("AI Classification Engine Demo")
	fmt.Println("====================================")

	for _, text := range textSamples {

		fmt.Println()
		fmt.Println("INPUT:")
		fmt.Println(text)

		result, err := ClassifyText(text)

		if err != nil {
			fmt.Println("ERROR:", err)
			continue
		}

		formattedJSON, _ := json.MarshalIndent(result, "", "  ")

		fmt.Println()
		fmt.Println("CLASSIFICATION:")
		fmt.Println(string(formattedJSON))

		fmt.Println("------------------------------------")
	}
}
