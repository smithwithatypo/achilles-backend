// handlers/transcription.go
package handlers

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "mime/multipart"
    "net/http"
    "os"
    "path/filepath"

    "github.com/smithwithatypo/achilles-backend/config"
)

// Response from OpenAI API
type TranscriptionResponse struct {
    Text string `json:"text"`
}

func TranscribeAudioHandler(w http.ResponseWriter, r *http.Request) {
    // Only allow POST requests
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    // Parse the multipart form data with a 10MB limit
    err := r.ParseMultipartForm(10 << 20)
    if err != nil {
        http.Error(w, "Could not parse form", http.StatusBadRequest)
        return
    }

    // Get the audio file from the form
    file, header, err := r.FormFile("audio")
    if err != nil {
        http.Error(w, "Error retrieving the file", http.StatusBadRequest)
        return
    }
    defer file.Close()

    // Create a temporary file to store the upload
    tempFile, err := os.CreateTemp("", "upload-*."+filepath.Ext(header.Filename))
    if err != nil {
        http.Error(w, "Error creating temporary file", http.StatusInternalServerError)
        return
    }
    defer os.Remove(tempFile.Name())
    defer tempFile.Close()

    // Copy the uploaded file to the temporary file
    _, err = io.Copy(tempFile, file)
    if err != nil {
        http.Error(w, "Error saving the file", http.StatusInternalServerError)
        return
    }

    // Send the file to OpenAI's Whisper API
    transcription, err := transcribeWithWhisper(tempFile.Name())
    if err != nil {
        http.Error(w, "Error transcribing audio: "+err.Error(), http.StatusInternalServerError)
        return
    }

    // Set the response headers
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    
    // Write the response
    json.NewEncoder(w).Encode(map[string]string{
        "transcription": transcription,
    })
}

func transcribeWithWhisper(filePath string) (string, error) {
    // Get the OpenAI API key from environment variables
    apiKey := config.GetEnv("OPENAI_API_KEY")
    if apiKey == "" {
        return "", fmt.Errorf("OPENAI_API_KEY environment variable not set")
    }

    // Prepare the request body as multipart/form-data
    var requestBody bytes.Buffer
    multipartWriter := multipart.NewWriter(&requestBody)

    // Add the model field
    err := multipartWriter.WriteField("model", "whisper-1")
    if err != nil {
        return "", err
    }

    // Add the file field
    fileWriter, err := multipartWriter.CreateFormFile("file", filepath.Base(filePath))
    if err != nil {
        return "", err
    }

    // Open the file
    file, err := os.Open(filePath)
    if err != nil {
        return "", err
    }
    defer file.Close()

    // Copy the file to the request body
    _, err = io.Copy(fileWriter, file)
    if err != nil {
        return "", err
    }

    // Close the multipart writer
    err = multipartWriter.Close()
    if err != nil {
        return "", err
    }

    // Create the HTTP request
    req, err := http.NewRequest("POST", "https://api.openai.com/v1/audio/transcriptions", &requestBody)
    if err != nil {
        return "", err
    }

    // Set the headers
    req.Header.Set("Content-Type", multipartWriter.FormDataContentType())
    req.Header.Set("Authorization", "Bearer "+apiKey)

    // Send the request
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    // Check the response status code
    if resp.StatusCode != http.StatusOK {
        responseBody, _ := io.ReadAll(resp.Body)
        return "", fmt.Errorf("API error: %s - %s", resp.Status, string(responseBody))
    }

    // Parse the response
    var transcriptionResponse TranscriptionResponse
    err = json.NewDecoder(resp.Body).Decode(&transcriptionResponse)
    if err != nil {
        return "", err
    }

    return transcriptionResponse.Text, nil
}