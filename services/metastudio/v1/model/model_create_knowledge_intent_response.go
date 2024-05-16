package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateKnowledgeIntentResponse Response Object
type CreateKnowledgeIntentResponse struct {

	// 意图ID。
	IntentId *string `json:"intent_id,omitempty"`

	// 意图标识。
	Identify *string `json:"identify,omitempty"`

	XRequestId     *string `json:"X-Request-Id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreateKnowledgeIntentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateKnowledgeIntentResponse struct{}"
	}

	return strings.Join([]string{"CreateKnowledgeIntentResponse", string(data)}, " ")
}
