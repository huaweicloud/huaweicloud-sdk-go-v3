package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateInstanceRequest struct {
	// bcs 服务id

	BlockchainId string `json:"blockchain_id"`

	Body *UpdateInstanceRequestBody `json:"body,omitempty"`
}

func (o UpdateInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateInstanceRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceRequest", string(data)}, " ")
}
