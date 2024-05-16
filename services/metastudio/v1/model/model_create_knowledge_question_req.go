package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateKnowledgeQuestionReq 创建知识库问法请求。
type CreateKnowledgeQuestionReq struct {

	// 问法。
	Question string `json:"question"`

	// 意图ID。
	IntentId *string `json:"intent_id,omitempty"`
}

func (o CreateKnowledgeQuestionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKnowledgeQuestionReq struct{}"
	}

	return strings.Join([]string{"CreateKnowledgeQuestionReq", string(data)}, " ")
}
