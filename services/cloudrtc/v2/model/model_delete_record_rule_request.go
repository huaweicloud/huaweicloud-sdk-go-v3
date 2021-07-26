package model

import (
	"encoding/json"

	"strings"
)

// Request Object
type DeleteRecordRuleRequest struct {
	// 使用AK/SK方式认证时必选，携带的鉴权信息。

	Authorization *string `json:"Authorization,omitempty"`
	// 使用AK/SK方式认证时必选，请求的发生时间。

	XSdkDate *string `json:"X-Sdk-Date,omitempty"`
	// 使用AK/SK方式认证时必选，携带项目ID信息。

	XProjectId *string `json:"X-Project-Id,omitempty"`
	// 应用id

	AppId string `json:"app_id"`
	// 规则id

	RuleId string `json:"rule_id"`
}

func (o DeleteRecordRuleRequest) String() string {
	data, err := json.Marshal(o)
	if err != nil {
		return "DeleteRecordRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteRecordRuleRequest", string(data)}, " ")
}
