package model

import (
	"encoding/json"

	"strings"
)

// Response Object
type ShowApiVersionResponse struct {
	Version        *VersionInfo `json:"version,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowApiVersionResponse) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowApiVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowApiVersionResponse", string(data)}, " ")
}
