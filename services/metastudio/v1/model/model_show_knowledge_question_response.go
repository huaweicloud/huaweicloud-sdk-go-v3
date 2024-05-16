package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowKnowledgeQuestionResponse Response Object
type ShowKnowledgeQuestionResponse struct {

	// 问法ID。
	QuestionId *string `json:"question_id,omitempty"`

	// 问法。
	Question *string `json:"question,omitempty"`

	// 意图ID。
	IntentId *string `json:"intent_id,omitempty"`

	// 创建时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	CreateTime *string `json:"create_time,omitempty"`

	// 更新时间，格式遵循：RFC 3339 如\"2021-01-10T08:43:17Z\"。
	UpdateTime *string `json:"update_time,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowKnowledgeQuestionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowKnowledgeQuestionResponse struct{}"
	}

	return strings.Join([]string{"ShowKnowledgeQuestionResponse", string(data)}, " ")
}
