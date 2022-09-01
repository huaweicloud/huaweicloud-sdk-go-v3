package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListBcsMetricRequest struct {

	// 区块链服务id,当前不支持IEF实例
	BlockchainId string `json:"blockchain_id" xml:"blockchain_id"`

	Body *ListBcsMetricRequestBody `json:"body,omitempty" xml:"body"`
}

func (o ListBcsMetricRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListBcsMetricRequest struct{}"
	}

	return strings.Join([]string{"ListBcsMetricRequest", string(data)}, " ")
}
