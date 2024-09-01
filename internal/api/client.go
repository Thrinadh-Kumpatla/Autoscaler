package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	BaseURL string
}

type AppStatus struct {
	CPU struct {
		HighPriority float64 `json:"highPriority"`
	} `json:"cpu"`
	Replicas int `json:"replicas"`
}

type ReplicaUpdate struct {
	Replicas int `json:"replicas"`
}

func NewClient(baseURL string) *Client {
	return &Client{BaseURL: baseURL}
}

func (c *Client) GetStatus() (*AppStatus, error) {
	fullURL := c.BaseURL + "/app/status"
	fmt.Printf("Accessing URL: %s\n", fullURL)

	// Create a new request
	req, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	// Add the Accept header
	req.Header.Add("Accept", "application/json")

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	fmt.Printf("Raw response: %s\n", string(body))

	var status AppStatus
	err = json.Unmarshal(body, &status)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v. Raw response: %s", err, string(body))
	}

	fmt.Printf("Parsed status: CPU=%v, Replicas=%d\n", status.CPU.HighPriority, status.Replicas)

	return &status, nil
}

func (c *Client) UpdateReplicas(replicas int) error {
    update := ReplicaUpdate{Replicas: replicas}
    jsonData, err := json.Marshal(update)
    if err != nil {
        return fmt.Errorf("error marshaling JSON: %v", err)
    }

    fullURL := c.BaseURL + "/app/replicas"
    req, err := http.NewRequest("PUT", fullURL, bytes.NewBuffer(jsonData))
    if err != nil {
        return fmt.Errorf("error creating request: %v", err)
    }

    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return fmt.Errorf("error making PUT request: %v", err)
    }
    defer resp.Body.Close()

    // Check for both 200 OK and 204 No Content as valid responses
    if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
        body, _ := io.ReadAll(resp.Body)
        return fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
    }

    return nil
}
