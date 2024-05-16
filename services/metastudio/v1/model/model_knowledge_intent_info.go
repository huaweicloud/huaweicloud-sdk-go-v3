package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// KnowledgeIntentInfo 知识库意图基本信息。
type KnowledgeIntentInfo struct {

	// 意图ID。
	IntentId *string `json:"intent_id,omitempty"`

	// 你是谁。
	Name *string `json:"name,omitempty"`

	// 意图标识。
	Identify *string `json:"identify,omitempty"`

	// 问题答案。
	Answer *string `json:"answer,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`
}

func (o KnowledgeIntentInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "KnowledgeIntentInfo struct{}"
	}

	return strings.Join([]string{"KnowledgeIntentInfo", string(data)}, " ")
}
