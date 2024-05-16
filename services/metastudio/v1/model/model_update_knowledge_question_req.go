package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateKnowledgeQuestionReq 修改知识库问法请求。
type UpdateKnowledgeQuestionReq struct {

	// 问法。
	Question *string `json:"question,omitempty"`
}

func (o UpdateKnowledgeQuestionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKnowledgeQuestionReq struct{}"
	}

	return strings.Join([]string{"UpdateKnowledgeQuestionReq", string(data)}, " ")
}
