package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	textSamples := []string{
	"Public press release for next quarter earnings",
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
