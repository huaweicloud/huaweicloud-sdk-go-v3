package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ChangeBaremetalServerOsRequest struct {
	// 裸金属服务器ID

	ServerId string `json:"server_id"`

	Body *OsChangeReq `json:"body,omitempty"`
}

func (o ChangeBaremetalServerOsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ChangeBaremetalServerOsRequest struct{}"
	}

	return strings.Join([]string{"ChangeBaremetalServerOsRequest", string(data)}, " ")
}
