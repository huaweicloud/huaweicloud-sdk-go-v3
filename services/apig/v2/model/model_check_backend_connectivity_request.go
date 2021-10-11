package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type CheckBackendConnectivityRequest struct {
	// 实例编号

	InstanceId string `json:"instance_id"`

	Body *CheckBackendConnectivityReq `json:"body,omitempty"`
}

func (o CheckBackendConnectivityRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CheckBackendConnectivityRequest struct{}"
	}

	return strings.Join([]string{"CheckBackendConnectivityRequest", string(data)}, " ")
}
