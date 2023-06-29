package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListEntityMetricRequest Request Object
type ListEntityMetricRequest struct {

	// 区块链服务id [目前不支持IEF实例](tag:hasief)
	BlockchainId string `json:"blockchain_id"`

	Body *ListEntityMetricRequestBody `json:"body,omitempty"`
}

func (o ListEntityMetricRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListEntityMetricRequest struct{}"
	}

	return strings.Join([]string{"ListEntityMetricRequest", string(data)}, " ")
}
