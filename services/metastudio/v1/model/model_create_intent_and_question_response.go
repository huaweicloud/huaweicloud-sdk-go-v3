package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateIntentAndQuestionResponse Response Object
type CreateIntentAndQuestionResponse struct {

	// 意图ID。
	IntentId string `json:"intent_id"`

	// 意图标识。
	Identify *string `json:"identify,omitempty"`

	// 问法列表
	QuestionList *[]KnowledgeQuestionInfo `json:"question_list,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateIntentAndQuestionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateIntentAndQuestionResponse struct{}"
	}

	return strings.Join([]string{"CreateIntentAndQuestionResponse", string(data)}, " ")
}
