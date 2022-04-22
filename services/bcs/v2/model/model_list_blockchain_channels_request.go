package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListBlockchainChannelsRequest struct {

	// blockchainID
	BlockchainId string `json:"blockchain_id"`
}

func (o ListBlockchainChannelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBlockchainChannelsRequest struct{}"
	}

	return strings.Join([]string{"ListBlockchainChannelsRequest", string(data)}, " ")
}
