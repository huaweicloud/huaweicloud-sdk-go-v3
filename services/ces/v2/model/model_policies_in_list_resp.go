package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PoliciesInListResp struct {

	// 告警策略ID。
	AlarmPolicyId string `json:"alarm_policy_id"`

	// 资源的监控指标名称，必须以字母开头，只能包含0-9/a-z/A-Z/_，字符长度最短为1，最大为64；如：弹性云服务器中的监控指标cpu_util，表示弹性服务器的CPU使用率；文档数据库中的指标mongo001_command_ps，表示command执行频率；各服务的指标名称可查看：“[服务指标名称](ces_03_0059.xml)”。
	MetricName string `json:"metric_name"`

	ExtraInfo *MetricExtraInfo `json:"extra_info,omitempty"`

	Period *Period `json:"period"`

	// 聚合方式, 支持的值为(average|min|max|sum)
	Filter string `json:"filter"`

	// 阈值符号, 支持的值为(>|<|>=|<=|=|!=|cycle_decrease|cycle_increase|cycle_wave);cycle_decrease为环比下降,cycle_increase为环比上升,cycle_wave为环比波动； 指标告警可以使用的阈值符号有>、>=、<、<=、=、!=、cycle_decrease、cycle_increase、cycle_wave； 事件告警可以使用的阈值符号为>、>=、<、<=、=、!=；
	ComparisonOperator string `json:"comparison_operator"`

	// 告警阈值。单一阈值时value和alarm_level配对使用，当hierarchical_value和value同时使用时以hierarchical_value为准。取值范围[0, Number.MAX_VALUE]，Number.MAX_VALUE值为1.7976931348623157e+108。具体阈值取值请参见附录中各服务监控指标中取值范围，如支持监控的服务列表中ECS的CPU使用率cpu_util取值范围可配置80。 [具体阈值取值请参见附录中各服务监控指标中取值范围，如[支持监控的服务列表](ces_03_0059.xml)中ECS的CPU使用率cpu_util取值范围可配置80。](tag: dt,g42,dt_test,hk_g42,hk_sbc,hws,hws_hk,ocb,sbc,tm)
	Value *float64 `json:"value,omitempty"`

	HierarchicalValue *HierarchicalValue `json:"hierarchical_value,omitempty"`

	// 数据的单位。
	Unit *string `json:"unit,omitempty"`

	// 告警连续触发次数，事件告警时参数值为1~180（包括1和180）；指标告警和站点告警时，次数采用枚举值，枚举值分别为：1、2、3、4、5、10、15、30、60、90、120、180
	Count int32 `json:"count"`

	// 告警策略类型。（暂时未使用）
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
