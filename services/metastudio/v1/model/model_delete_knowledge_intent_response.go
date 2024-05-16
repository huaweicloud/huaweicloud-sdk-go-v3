package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteKnowledgeIntentResponse Response Object
type DeleteKnowledgeIntentResponse struct {
	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o DeleteKnowledgeIntentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteKnowledgeIntentResponse struct{}"
	}

	return strings.Join([]string{"DeleteKnowledgeIntentResponse", string(data)}, " ")
}
