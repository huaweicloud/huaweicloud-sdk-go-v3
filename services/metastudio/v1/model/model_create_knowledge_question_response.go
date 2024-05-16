package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateKnowledgeQuestionResponse Response Object
type CreateKnowledgeQuestionResponse struct {

	// 问法ID。
	QuestionId *string `json:"question_id,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateKnowledgeQuestionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKnowledgeQuestionResponse struct{}"
	}

	return strings.Join([]string{"CreateKnowledgeQuestionResponse", string(data)}, " ")
}
