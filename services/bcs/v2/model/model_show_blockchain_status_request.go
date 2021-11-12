package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ShowBlockchainStatusRequest struct {
	// blockchainID

	BlockchainId string `json:"blockchain_id"`
}

func (o ShowBlockchainStatusRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBlockchainStatusRequest struct{}"
	}

	return strings.Join([]string{"ShowBlockchainStatusRequest", string(data)}, " ")
}
