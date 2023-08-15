package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowBlockchainDetailRequest Request Object
type ShowBlockchainDetailRequest struct {

	// blockchainID
	BlockchainId string `json:"blockchain_id"`
}

func (o ShowBlockchainDetailRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowBlockchainDetailRequest struct{}"
	}

	return strings.Join([]string{"ShowBlockchainDetailRequest", string(data)}, " ")
}
