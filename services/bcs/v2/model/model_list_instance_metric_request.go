package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListInstanceMetricRequest struct {
	// 区块链服务id。

	BlockchainId string `json:"blockchain_id"`

	Body *ListInstanceMetricRequestBody `json:"body,omitempty"`
}

func (o ListInstanceMetricRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListInstanceMetricRequest struct{}"
	}

	return strings.Join([]string{"ListInstanceMetricRequest", string(data)}, " ")
}
