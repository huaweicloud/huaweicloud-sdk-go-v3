package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowVolumeResponse struct {
	Volume         *Volume `json:"volume,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowVolumeResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowVolumeResponse struct{}"
	}

	return strings.Join([]string{"ShowVolumeResponse", string(data)}, " ")
}
