package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListBlockchainChannelsRequest struct {
	// blockchainID

	BlockchainId string `json:"blockchain_id"`
}

func (o ListBlockchainChannelsRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListBlockchainChannelsRequest struct{}"
	}

	return strings.Join([]string{"ListBlockchainChannelsRequest", string(data)}, " ")
}
