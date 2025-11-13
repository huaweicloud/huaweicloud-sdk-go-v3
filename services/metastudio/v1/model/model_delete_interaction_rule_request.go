package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// DeleteInteractionRuleRequest Request Object
type DeleteInteractionRuleRequest struct {

	// 使用AK/SK方式认证时必选，携带的鉴权信息。
	Authorization *string `json:"Authorization,omitempty"`

	// 使用AK/SK方式认证时必选，请求的发生时间。  格式为(YYYYMMDD'T'HHMMSS'Z')。
	XSdkDate *string `json:"X-Sdk-Date,omitempty"`

	// 使用AK/SK方式认证时必选，携带项目ID信息。
	XProjectId *string `json:"X-Project-Id,omitempty"`

	// 第三方用户ID。不允许输入中文。
	XAppUserId *string `json:"X-App-UserId,omitempty"`

	// 互动规则库ID。
	GroupId string `json:"group_id"`

	// 互动规则ID。
	RuleId string `json:"rule_id"`
}

func (o DeleteInteractionRuleRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "DeleteInteractionRuleRequest struct{}"
	}

	return strings.Join([]string{"DeleteInteractionRuleRequest", string(data)}, " ")
}
