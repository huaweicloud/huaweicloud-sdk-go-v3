package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UpdateKnowledgeIntentResponse Response Object
type UpdateKnowledgeIntentResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o UpdateKnowledgeIntentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateKnowledgeIntentResponse struct{}"
	}

	return strings.Join([]string{"UpdateKnowledgeIntentResponse", string(data)}, " ")
}
