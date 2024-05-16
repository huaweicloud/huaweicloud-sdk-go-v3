package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateKnowledgeIntentReq 修改知识库意图请求。
type UpdateKnowledgeIntentReq struct {

	// 意图名称。
	Name *string `json:"name,omitempty"`

	// 问题答案。
	Answer *string `json:"answer,omitempty"`
}

func (o UpdateKnowledgeIntentReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKnowledgeIntentReq struct{}"
	}

	return strings.Join([]string{"UpdateKnowledgeIntentReq", string(data)}, " ")
}
