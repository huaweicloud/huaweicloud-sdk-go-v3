package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBatchKnowledgeQuestionReq 批量创建知识库问法请求。
type CreateBatchKnowledgeQuestionReq struct {

	// 意图ID。
	IntentId string `json:"intent_id"`

	// 问法列表
	QuestionList []KnowledgeQuestionCreateInfo `json:"question_list"`
}

func (o CreateBatchKnowledgeQuestionReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBatchKnowledgeQuestionReq struct{}"
	}

	return strings.Join([]string{"CreateBatchKnowledgeQuestionReq", string(data)}, " ")
}
