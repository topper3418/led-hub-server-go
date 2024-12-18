package client

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
)

type Color struct {
    R int `json:"r"`
    G int `json:"g"`
    B int `json:"b"`
}

type LedStripState struct {
    Brightness *int   `json:"brightness,omitempty"`
    On         bool   `json:"on"`
    Color      *Color `json:"color,omitempty"`
}

// GetLedStripState sends a GET request to the LED strip’s URL and returns its current state.
func GetLedStripState(url string) (*LedStripState, error) {
    resp, err := http.Get(url)
    if err != nil {
        return nil, fmt.Errorf("failed to GET from %s: %w", url, err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        bodyBytes, _ := ioutil.ReadAll(resp.Body)
        return nil, fmt.Errorf("GET request failed with status %d: %s", resp.StatusCode, string(bodyBytes))
    }

    var state LedStripState
    if err := json.NewDecoder(resp.Body).Decode(&state); err != nil {
        return nil, fmt.Errorf("failed to decode response body: %w", err)
    }

    return &state, nil
}

// SetLedStripState sends a POST request to update the LED strip’s state.
func SetLedStripState(url string, state LedStripState) error {
    data, err := json.Marshal(state)
    if err != nil {
        return fmt.Errorf("failed to marshal state: %w", err)
    }

    resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
    if err != nil {
        return fmt.Errorf("failed to POST to %s: %w", url, err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
        bodyBytes, _ := ioutil.ReadAll(resp.Body)
        return fmt.Errorf("POST request failed with status %d: %s", resp.StatusCode, string(bodyBytes))
    }

    return nil
}
