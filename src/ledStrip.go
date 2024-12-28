package ledStrip

import (
    "fmt"
    "log"

    "your-project/src/client"
    "your-project/src/device"
)

type LedStrip struct {
    dev *device.Device
}

// NewLedStrip fetches a device from the database by ID and ensures it is a LedStrip.
func NewLedStrip(id int) (*LedStrip, error) {
    d, err := device.Get(id)
    if err != nil {
        return nil, fmt.Errorf("unable to retrieve device from DB: %w", err)
    }

    if d.Type != "LedStrip" {
        return nil, fmt.Errorf("device is not a LedStrip, got type: %s", d.Type)
    }

    if !d.CurrentIP.Valid || d.CurrentIP.String == "" {
        return nil, fmt.Errorf("device does not have a valid IP")
    }

    return &LedStrip{dev: d}, nil
}

// GetState uses the device's IP as the URL to fetch the current LED strip state.
func (ls *LedStrip) GetState() (*client.LedStripState, error) {
    url := "http://" + ls.dev.CurrentIP.String
    return client.GetLedStripState(url)
}

// SetState uses the device's IP as the URL to update the LED strip state.
func (ls *LedStrip) SetState(state client.LedStripState) error {
    url := "http://" + ls.dev.CurrentIP.String
    return client.SetLedStripState(url, state)
}

func intPtr(i int) *int {
    return &i
}
