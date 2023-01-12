package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowPlaybookRuleResponse struct {

	// Error code
	Code *string `json:"code,omitempty"`

	// Error message
	Message *string `json:"message,omitempty"`

	Data *RuleInfo `json:"data,omitempty"`

	XRequestId     *string `json:"X-request-id,omitempty"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowPlaybookRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlaybookRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowPlaybookRuleResponse", string(data)}, " ")
}
