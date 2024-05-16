package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateKnowledgeQuestionResponse Response Object
type UpdateKnowledgeQuestionResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateKnowledgeQuestionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKnowledgeQuestionResponse struct{}"
	}

	return strings.Join([]string{"UpdateKnowledgeQuestionResponse", string(data)}, " ")
}
