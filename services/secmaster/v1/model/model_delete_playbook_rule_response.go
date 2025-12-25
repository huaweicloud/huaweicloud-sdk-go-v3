package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeletePlaybookRuleResponse Response Object
type DeletePlaybookRuleResponse struct {

	// 错误码
	Code *string `json:"code,omitempty"`

	// 响应消息
	Message *string `json:"message,omitempty"`

	XRequestId *string `json:"X-request-id,omitempty"`

	XAPIDeprecationInfo *string `json:"X-API-Deprecation-Info,omitempty"`

	XAPIDeprecationDate *string `json:"X-API-Deprecation-Date,omitempty"`
	HttpStatusCode      int     `json:"-"`
}

func (o DeletePlaybookRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeletePlaybookRuleResponse struct{}"
	}

	return strings.Join([]string{"DeletePlaybookRuleResponse", string(data)}, " ")
}
