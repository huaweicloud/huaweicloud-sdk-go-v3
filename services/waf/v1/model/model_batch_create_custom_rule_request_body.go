package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// BatchCreateCustomRuleRequestBody 创建精准防护规则body
type BatchCreateCustomRuleRequestBody struct {

	// 精准防护规则生效时间:  - “false”：表示该规则立即生效。   - “true”：表示自定义生效时间。
	Time bool `json:"time"`

	// 精准防护规则生效的起始时间戳（秒）。当time=true，才需要填写该参数。
	Start *int64 `json:"start,omitempty"`

	// 精准防护规则生效的终止时间戳（秒）。当time=true，才需要填写该参数。
	Terminal *int64 `json:"terminal,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// 匹配条件列表
	Conditions []CustomConditions `json:"conditions"`

	Action *CustomAction `json:"action"`

	// 执行该规则的优先级，值越小，优先级越高，值相同时，规则创建时间早，优先级越高。取值范围：0到65535。
	Priority int32 `json:"priority"`

	// 规则名称
	Name string `json:"name"`

	// 添加规则的策略id列表。策略id从\"查询防护策略列表\"(ListPolicy)接口获取，多个策略之间用“,”隔开
	PolicyIds []string `json:"policy_ids"`
}

func (o BatchCreateCustomRuleRequestBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateCustomRuleRequestBody struct{}"
	}

	return strings.Join([]string{"BatchCreateCustomRuleRequestBody", string(data)}, " ")
}
