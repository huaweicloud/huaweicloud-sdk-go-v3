package model

import (
	"encoding/json"

	"strings"
)

// 区块生成配置信息
type CreateRequestBodyBlockInfo struct {
	// 区块产生时间（单位：秒），默认2秒

	BatchTimeout *int32 `json:"batch_timeout,omitempty"`
	// 区块包含交易数量，默认500

	MaxMessageCount *int32 `json:"max_message_count,omitempty"`
	// 区块容量（单位：MB），默认2MB

	PreferredMaxbytes *int32 `json:"preferred_maxbytes,omitempty"`
}

func (o CreateRequestBodyBlockInfo) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "CreateRequestBodyBlockInfo struct{}"
	}

	return strings.Join([]string{"CreateRequestBodyBlockInfo", string(data)}, " ")
}
