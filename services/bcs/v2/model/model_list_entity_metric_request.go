package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type ListEntityMetricRequest struct {
	// 区块链服务id，目前不支持IEF实例

	BlockchainId string `json:"blockchain_id"`

	Body *ListEntityMetricRequestBody `json:"body,omitempty"`
}

func (o ListEntityMetricRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "ListEntityMetricRequest struct{}"
	}

	return strings.Join([]string{"ListEntityMetricRequest", string(data)}, " ")
}
