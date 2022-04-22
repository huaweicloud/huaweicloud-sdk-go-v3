package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowBlockchainNodesRequest struct {

	// blockchainID
	BlockchainId string `json:"blockchain_id"`
}

func (o ShowBlockchainNodesRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBlockchainNodesRequest struct{}"
	}

	return strings.Join([]string{"ShowBlockchainNodesRequest", string(data)}, " ")
}
