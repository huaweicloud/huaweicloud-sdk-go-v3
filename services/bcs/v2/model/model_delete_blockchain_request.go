package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteBlockchainRequest struct {
	BlockchainId string `json:"blockchain_id"`

	IsDeleteStorage *bool `json:"is_delete_storage,omitempty"`

	IsDeleteObs *bool `json:"is_delete_obs,omitempty"`

	IsDeleteResource *bool `json:"is_delete_resource,omitempty"`
}

func (o DeleteBlockchainRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteBlockchainRequest struct{}"
	}

	return strings.Join([]string{"DeleteBlockchainRequest", string(data)}, " ")
}
