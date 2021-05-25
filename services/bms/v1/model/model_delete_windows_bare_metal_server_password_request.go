package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteWindowsBareMetalServerPasswordRequest struct {
	// 裸金属服务器ID

	ServerId string `json:"server_id"`
}

func (o DeleteWindowsBareMetalServerPasswordRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteWindowsBareMetalServerPasswordRequest struct{}"
	}

	return strings.Join([]string{"DeleteWindowsBareMetalServerPasswordRequest", string(data)}, " ")
}
