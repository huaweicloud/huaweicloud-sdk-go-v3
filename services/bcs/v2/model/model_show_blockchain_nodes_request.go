package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowBlockchainNodesRequest struct {
	// blockchainID

	BlockchainId string `json:"blockchain_id"`
}

func (o ShowBlockchainNodesRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowBlockchainNodesRequest struct{}"
	}

	return strings.Join([]string{"ShowBlockchainNodesRequest", string(data)}, " ")
}
