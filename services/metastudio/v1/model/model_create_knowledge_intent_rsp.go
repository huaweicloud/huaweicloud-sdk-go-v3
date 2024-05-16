package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateKnowledgeIntentRsp 创建知识库意图响应。
type CreateKnowledgeIntentRsp struct {

	// 意图ID。
	IntentId string `json:"intent_id"`

	// 意图标识。
	Identify *string `json:"identify,omitempty"`
}

func (o CreateKnowledgeIntentRsp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKnowledgeIntentRsp struct{}"
	}

	return strings.Join([]string{"CreateKnowledgeIntentRsp", string(data)}, " ")
}
