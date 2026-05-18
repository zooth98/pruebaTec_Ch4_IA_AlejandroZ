package main

type ClassificationResult struct {
	Category    string `json:"category"`
	Sensitivity string `json:"sensitivity"`
	Risk        string `json:"risk"`
	Reason      string `json:"reason"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type OllamaRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
	Stream   bool      `json:"stream"`
}

type OllamaResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}