package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreatePlaybookRuleResponse Response Object
type CreatePlaybookRuleResponse struct {

	// Error code
	Code *string `json:"code,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	Data *RuleInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o CreatePlaybookRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreatePlaybookRuleResponse struct{}"
	}

	return strings.Join([]string{"CreatePlaybookRuleResponse", string(data)}, " ")
}
