package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteKnowledgeQuestionResponse Response Object
type DeleteKnowledgeQuestionResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteKnowledgeQuestionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKnowledgeQuestionResponse struct{}"
	}

	return strings.Join([]string{"DeleteKnowledgeQuestionResponse", string(data)}, " ")
}
