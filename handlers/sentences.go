package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"net/http"
)

// OpenAIRequest defines the expected request structure from the frontend
type OpenAIRequest struct {
	Prompt string `json:"prompt"`
	Model  string `json:"model,omitempty"` // Optional, can set a default if not provided
}

// OpenAIAPIRequest defines the structure to send to OpenAI
type OpenAIAPIRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Temperature float64   `json:"temperature,omitempty"`
}

// Message represents a message in the OpenAI chat format
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// OpenAIResponse represents the response from OpenAI API
type OpenAIResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Choices []Choice `json:"choices"`
	// Other OpenAI response fields can be added as needed
}

// Choice represents a completion choice from OpenAI
type Choice struct {
	Message      Message `json:"message"`
	Index        int     `json:"index"`
	FinishReason string  `json:"finish_reason"`
}

// HandleOpenAIRequest processes requests to the OpenAI API
func HandleOpenAIRequest(w http.ResponseWriter, r *http.Request) {
	// Only allow POST method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the incoming request
	var request OpenAIRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Set default model if not provided
	model := request.Model
	if model == "" {
		model = "gpt-4o-mini" // Default model
	}

	// Create the OpenAI API request
	openAIReq := OpenAIAPIRequest{
		Model: model,
		Messages: []Message{
			{
				Role:    "user",
				Content: request.Prompt,
			},
		},
		Temperature: 1,
	}

	// Convert the request to JSON
	reqBody, err := json.Marshal(openAIReq)
	if err != nil {
		http.Error(w, "Failed to create API request", http.StatusInternalServerError)
		return
	}

	// Create the HTTP request to OpenAI
	openaiURL := "https://api.openai.com/v1/chat/completions"
	req, err := http.NewRequest("POST", openaiURL, bytes.NewBuffer(reqBody))
	if err != nil {
		http.Error(w, "Failed to create API request", http.StatusInternalServerError)
		return
	}

	// Add necessary headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+getOpenAIAPIKey())

	// Make the request to OpenAI
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to make API request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read API response", http.StatusInternalServerError)
		return
	}

	// Check if OpenAI returned an error
	if resp.StatusCode != http.StatusOK {
		http.Error(w, "OpenAI API error: "+string(body), resp.StatusCode)
		return
	}

	// Parse the OpenAI response
	var openAIResp OpenAIResponse
	if err := json.Unmarshal(body, &openAIResp); err != nil {
		http.Error(w, "Failed to parse API response", http.StatusInternalServerError)
		return
	}

	// Extract the response text
	responseText := ""
	if len(openAIResp.Choices) > 0 {
		responseText = openAIResp.Choices[0].Message.Content
	}

	// Return the response to the client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"text": responseText,
	})
}

// getOpenAIAPIKey retrieves the OpenAI API key from environment variables
func getOpenAIAPIKey() string {
	return os.Getenv("OPENAI_API_KEY")
}
