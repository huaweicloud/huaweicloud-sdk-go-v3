package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteWindowsBareMetalServerPasswordRequest struct {
	ServerId string `json:"server_id"`
}

func (o DeleteWindowsBareMetalServerPasswordRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteWindowsBareMetalServerPasswordRequest struct{}"
	}

	return strings.Join([]string{"DeleteWindowsBareMetalServerPasswordRequest", string(data)}, " ")
}
