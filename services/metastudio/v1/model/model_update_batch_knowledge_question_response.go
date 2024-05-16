package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateBatchKnowledgeQuestionResponse Response Object
type UpdateBatchKnowledgeQuestionResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateBatchKnowledgeQuestionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateBatchKnowledgeQuestionResponse struct{}"
	}

	return strings.Join([]string{"UpdateBatchKnowledgeQuestionResponse", string(data)}, " ")
}
