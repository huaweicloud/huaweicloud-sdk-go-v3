package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CreateAppRuleResponse Response Object
type CreateAppRuleResponse struct {

	// 规则ID。
	Id *string `json:"id,omitempty"`

	// 规则名称。
	Name *string `json:"name,omitempty"`

	// 规则描述。
	Description *string `json:"description,omitempty"`

	Rule *Rule `json:"rule,omitempty"`

	// 规则来源。
	RuleSource *string `json:"rule_source,omitempty"`

	// 创建时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间。
	UpdateTime     *sdktime.SdkTime `json:"update_time,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CreateAppRuleResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateAppRuleResponse struct{}"
	}

	return strings.Join([]string{"CreateAppRuleResponse", string(data)}, " ")
}
