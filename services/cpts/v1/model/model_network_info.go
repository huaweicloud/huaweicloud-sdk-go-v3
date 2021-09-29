package model

import (
	"encoding/json"

	"strings"
)

type NetworkInfo struct {
	// network_type

	NetworkType *string `json:"network_type,omitempty"`
}

func (o NetworkInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "NetworkInfo struct{}"
	}

	return strings.Join([]string{"NetworkInfo", string(data)}, " ")
}
