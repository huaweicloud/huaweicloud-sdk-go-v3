package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateBatchKnowledgeQuestionResponse Response Object
type CreateBatchKnowledgeQuestionResponse struct {

	// 问法列表
	QuestionList *[]KnowledgeQuestionInfo `json:"question_list,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateBatchKnowledgeQuestionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateBatchKnowledgeQuestionResponse struct{}"
	}

	return strings.Join([]string{"CreateBatchKnowledgeQuestionResponse", string(data)}, " ")
}
