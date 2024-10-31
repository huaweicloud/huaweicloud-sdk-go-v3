package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type ShowHttpCcRuleResponseBody struct {

	// 规则id
	Id *string `json:"id,omitempty"`

	// 规则名称
	Name *string `json:"name,omitempty"`

	// cc规则优先级，越大优先级越高，默认1
	Priority *int32 `json:"priority,omitempty"`

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

	// 规则类型（0：标准/1：高级）
	Mode *int32 `json:"mode,omitempty"`

	// 所有用户的周期内请求次数
	TotalNum *int32 `json:"total_num,omitempty"`

	// 单个用户的周期内请求次数
	LimitNum *int32 `json:"limit_num,omitempty"`

	// 限速周期
	LimitPeriod *int32 `json:"limit_period,omitempty"`

	// 锁定时长
	LockTime *int32 `json:"lock_time,omitempty"`

	// 防护模式
	TagType *string `json:"tag_type,omitempty"`

	// 防护模式标签
	TagIndex *string `json:"tag_index,omitempty"`

	TagCondition *HttpCcRuleCondition `json:"tag_condition,omitempty"`

	// 放行次数
	UnlockNum *int32 `json:"unlock_num,omitempty"`

	// 是否聚合域名
	DomainAggregation *bool `json:"domain_aggregation,omitempty"`

	// 条件列表参数较为复杂，存在级联关系，建议同时使用控制台上的添加误报屏蔽规则，单击F12键查看路径后缀为/cc-rule，方法为POST的请求参数，便于理解参数的填写
	Conditions *[]HttpCcRuleCondition `json:"conditions,omitempty"`

	Action *HttpRuleAction `json:"action,omitempty"`

	// 创建来源
	Producer *int32 `json:"producer,omitempty"`

	// 生效模式
	TimeMode *string `json:"time_mode,omitempty"`

	// customize生效时间区间开始
	Start *int64 `json:"start,omitempty"`

	// customize生效时间区间结束
	Terminal *int64 `json:"terminal,omitempty"`

	// period每日生效时间类型，目前只有day
	PeriodType *string `json:"period_type,omitempty"`

	// period每日生效时间区间
	TimeRange *[]TimeRangeItem `json:"time_range,omitempty"`
}

func (o ShowHttpCcRuleResponseBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowHttpCcRuleResponseBody struct{}"
	}

	return strings.Join([]string{"ShowHttpCcRuleResponseBody", string(data)}, " ")
}
