package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ShowBlockchainDetailRequest struct {
	// blockchainID

	BlockchainId string `json:"blockchain_id"`
}

func (o ShowBlockchainDetailRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ShowBlockchainDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowBlockchainDetailRequest", string(data)}, " ")
}
