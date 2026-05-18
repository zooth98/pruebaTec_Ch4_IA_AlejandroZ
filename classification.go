package models

type ClassificationRequest struct {
	Text string `json:"text"`
}

type ClassificationResponse struct {
	Category    string `json:"category"`
	Sensitivity string `json:"sensitivity"`
	Risk        string `json:"risk"`
	Reason      string `json:"reason"`
}

