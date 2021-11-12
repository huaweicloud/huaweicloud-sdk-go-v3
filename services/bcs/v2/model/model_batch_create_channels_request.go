package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type BatchCreateChannelsRequest struct {
	// bcs 服务id

	BlockchainId string `json:"blockchain_id"`

	Body *BatchCreateChannelsRequestBody `json:"body,omitempty"`
}

func (o BatchCreateChannelsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateChannelsRequest struct{}"
	}

	return strings.Join([]string{"BatchCreateChannelsRequest", string(data)}, " ")
}
