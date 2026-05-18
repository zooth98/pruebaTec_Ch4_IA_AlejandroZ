package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	OllamaURL = "http://localhost:11434/v1/chat/completions"
	ModelName = "llama3"
)

func CallLLM(systemPrompt string, userText string) (string, error) {

	requestBody := OllamaRequest{
		Model: ModelName,
		Messages: []Message{
			{
				Role:    "system",
				Content: systemPrompt,
			},
			{
				Role:    "user",
				Content: userText,
			},
		},
		Stream: false,
	}

	jsonData, err := json.Marshal(requestBody)

	if err != nil {
		return "", err
	}

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	response, err := client.Post(
		OllamaURL,
		"application/json",
		bytes.NewBuffer(jsonData),
	)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	if response.StatusCode != http.StatusOK {
		return "", fmt.Errorf("ollama error: %s", string(body))
	}

	var ollamaResponse OllamaResponse

	err = json.Unmarshal(body, &ollamaResponse)

	if err != nil {
		return "", err
	}

	if len(ollamaResponse.Choices) == 0 {
		return "", fmt.Errorf("empty response from llm")
	}

	return ollamaResponse.Choices[0].Message.Content, nil
}