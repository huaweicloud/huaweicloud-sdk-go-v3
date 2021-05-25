package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListBcsMetricRequest struct {
	// 区块链服务id,当前不支持IEF实例

	BlockchainId string `json:"blockchain_id"`

	Body *ListBcsMetricRequestBody `json:"body,omitempty"`
}

func (o ListBcsMetricRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListBcsMetricRequest struct{}"
	}

	return strings.Join([]string{"ListBcsMetricRequest", string(data)}, " ")
}
