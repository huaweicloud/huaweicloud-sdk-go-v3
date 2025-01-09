package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmCondition 告警触发条件
type AlarmCondition struct {

	// 指标周期，单位是秒； 0是默认值，例如事件类告警该字段就用0即可； 1代表指标的原始周期，比如RDS监控指标原始周期是60s，表示该RDS指标按60s周期为一个数据点参与告警计算；
	Period *int32 `json:"period,omitempty"`

	// 聚合方式, 支持的值为(average|min|max|sum)
	Filter *string `json:"filter,omitempty"`

	// 阈值符号,支持的值为(>|<|>=|<=|=|!=|cycle_decrease|cycle_increase|cycle_wave);cycle_decrease为环比下降,cycle_increase为环比上升,cycle_wave为环比波动
	ComparisonOperator *string `json:"comparison_operator,omitempty"`

	// 告警阈值
	Value *float64 `json:"value,omitempty"`

	// 数据的单位，最大长度为32位
	Unit *string `json:"unit,omitempty"`

	// 次数
	Count *int32 `json:"count,omitempty"`

	// 告警抑制时间，单位为秒，对应页面上创建告警规则时告警策略最后一个字段，该字段主要为解决告警频繁的问题，0代表不抑制，满足条件即告警；300代表满足告警触发条件后每5分钟告警一次；
	SuppressDuration *int32 `json:"suppress_duration,omitempty"`
}

func (o AlarmCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmCondition struct{}"
	}

	return strings.Join([]string{"AlarmCondition", string(data)}, " ")
}
