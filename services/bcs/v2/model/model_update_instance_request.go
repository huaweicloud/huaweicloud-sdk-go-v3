package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type UpdateInstanceRequest struct {
	BlockchainId string `json:"blockchain_id"`

	Body *OpIdRes `json:"body,omitempty"`
}

func (o UpdateInstanceRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "UpdateInstanceRequest struct{}"
	}

	return strings.Join([]string{"UpdateInstanceRequest", string(data)}, " ")
}
