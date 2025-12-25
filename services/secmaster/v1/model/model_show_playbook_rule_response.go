package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowPlaybookRuleResponse Response Object
type ShowPlaybookRuleResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 错误信息
	Message *string `json:"message,omitempty"`

	Data *RuleInfo `json:"data,omitempty"`

	XRequestId *string `json:"X-request-id,omitempty"`

	XAPIDeprecationInfo *string `json:"X-API-Deprecation-Info,omitempty"`

	XAPIDeprecationDate *string `json:"X-API-Deprecation-Date,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o ShowPlaybookRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowPlaybookRuleResponse struct{}"
	}

	return strings.Join([]string{"ShowPlaybookRuleResponse", string(data)}, " ")
}
