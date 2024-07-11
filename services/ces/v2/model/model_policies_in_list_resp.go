package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PoliciesInListResp struct {

	// 告警策略ID。
	AlarmPolicyId string `json:"alarm_policy_id"`

	// 资源的监控指标名称，必须以字母开头，只能包含0-9/a-z/A-Z/_，字符长度最短为1，最大为64；如：弹性云服务器中的监控指标cpu_util，表示弹性服务器的CPU使用率；文档数据库中的指标mongo001_command_ps，表示command执行频率；各服务的指标名称可查看：“[服务指标名称](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。
	MetricName string `json:"metric_name"`

	ExtraInfo *MetricExtraInfo `json:"extra_info,omitempty"`

	Period *Period `json:"period"`

	// 聚合方式, 支持的值为(average|min|max|sum)
	Filter string `json:"filter"`

	// 告警阈值的比较条件，支持的值为(>|<|>=|<=|=|!=|cycle_decrease|cycle_increase|cycle_wave)，cycle_decrease为环比下降，cycle_increase为环比上升，cycle_wave为环比波动
	ComparisonOperator string `json:"comparison_operator"`

	// 阈值
	Value float64 `json:"value"`

	// 单位
	Unit *string `json:"unit,omitempty"`

	// 次数
	Count int32 `json:"count"`

	// 告警策略类型。
	Type *string `json:"type,omitempty"`

	SuppressDuration *SuppressDuration `json:"suppress_duration,omitempty"`

	// 告警级别，1为紧急，2为重要，3为次要，4为提示
	AlarmLevel *int32 `json:"alarm_level,omitempty"`

	// 用户在页面中选择的指标单位， 用于后续指标数据回显和计算
	SelectedUnit *string `json:"selected_unit,omitempty"`
}

func (o PoliciesInListResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesInListResp struct{}"
	}

	return strings.Join([]string{"PoliciesInListResp", string(data)}, " ")
}
