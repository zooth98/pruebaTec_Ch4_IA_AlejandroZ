package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func ClassifyText(text string) (*ClassificationResult, error) {

	systemPrompt, err := os.ReadFile("system_prompt.md")

	if err != nil {
		return nil, fmt.Errorf("error reading system prompt: %w", err)
	}

	response, err := CallLLM(string(systemPrompt), text)

	if err != nil {
		return nil, err
	}

	var result ClassificationResult

	err = json.Unmarshal([]byte(response), &result)

	if err != nil {
		return nil, fmt.Errorf("invalid json returned by llm: %w", err)
	}

	return &result, nil
}