package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowHttpIgnoreRuleResponseBody struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 规则名称
	Name *string `json:"name,omitempty"`

	// 规则所在策略id
	PolicyId *string `json:"policy_id,omitempty"`

	// 规则所在策略名称
	PolicyName *string `json:"policy_name,omitempty"`

	// 创建规则时间戳
	Timestamp *int64 `json:"timestamp,omitempty"`

	// 规则描述
	Description *string `json:"description,omitempty"`

	// 规则开关状态
	Status *int32 `json:"status,omitempty"`

	// 误报路径
	Url *string `json:"url,omitempty"`

	// 规则编号
	Rule *string `json:"rule,omitempty"`

	// 误报屏蔽模式，默认为0即旧模式
	Mode *int32 `json:"mode,omitempty"`

	// 域名列表
	Domains *[]string `json:"domains,omitempty"`

	// 屏蔽规则url类型（前缀：prefix，等于：equal）
	UrlLogic *string `json:"url_logic,omitempty"`

	Advanced *HttpIgnoreRuleCondition `json:"advanced,omitempty"`

	// 命中条件
	Conditions *[]HttpIgnoreRuleCondition `json:"conditions,omitempty"`

	// 命中次数
	HitNum *int32 `json:"hit_num,omitempty"`

	// 最后更新时间戳
	UpdateTime *int64 `json:"update_time,omitempty"`

	// 上一次命中次数清零时间戳
	ClearTime *int64 `json:"clear_time,omitempty"`
}

func (o ShowHttpIgnoreRuleResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpIgnoreRuleResponseBody struct{}"
	}

	return strings.Join([]string{"ShowHttpIgnoreRuleResponseBody", string(data)}, " ")
}
