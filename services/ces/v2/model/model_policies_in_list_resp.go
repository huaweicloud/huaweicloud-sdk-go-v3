package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type PoliciesInListResp struct {

	// 告警策略ID。
	AlarmPolicyId string `json:"alarm_policy_id"`

	// **参数解释**： 资源的监控指标名称，各服务的指标名称可查看：“[服务指标名称](ces_03_0059.xml)”。 **约束限制**： 不涉及。 **取值范围**： 必须以字母开头，只能包含0-9/a-z/A-Z/_/-。字符长度最短为1，最大为96。如：弹性云服务器中的监控指标cpu_util，表示弹性服务器的CPU使用率；文档数据库中的指标mongo001_command_ps，表示command执行频率。         **默认取值**： 不涉及。
	MetricName string `json:"metric_name"`

	ExtraInfo *MetricExtraInfo `json:"extra_info,omitempty"`

	Period *Period `json:"period"`

	// **参数解释**： 聚合方式。         **约束限制**： 不涉及。 **取值范围**： average： 平均值，variance：方差，min：最小值，max：最大值，sum：求和。           **默认取值**： 不涉及。
	Filter string `json:"filter"`

	// **参数解释**： 阈值符号。     **约束限制**： 指标告警可以使用的阈值符号有>、>=、<、<=、=、!=、cycle_decrease、cycle_increase、cycle_wave； 事件告警可以使用的阈值符号为>、>=、<、<=、=、!=。 **取值范围**： 支持的值为(>|<|>=|<=|=|!=|cycle_decrease|cycle_increase|cycle_wave);cycle_decrease为环比下降,cycle_increase为环比上升,cycle_wave为环比波动。           **默认取值**： 不涉及。
	ComparisonOperator string `json:"comparison_operator"`

	// **参数解释**： 告警阈值。具体阈值取值请参见附录中各服务监控指标中取值范围。    **约束限制**： 单一阈值时value和alarm_level配对使用，当hierarchical_value和value同时使用时以hierarchical_value为准。 **取值范围**： 最小值为-1.7976931348623157e+108，最大值为1.7976931348623157e+108。           **默认取值**： 不涉及。
	Value *float64 `json:"value,omitempty"`

	HierarchicalValue *HierarchicalValue `json:"hierarchical_value,omitempty"`

	// **参数解释**： 数据的单位。    **约束限制**： 不涉及。 **取值范围**： 长度为[0,32]个字符。         **默认取值**： 不涉及。
	Unit *string `json:"unit,omitempty"`

	// **参数解释**： 告警连续触发次数。     **约束限制**： 不涉及。 **取值范围**： 事件告警时参数值为1~180（包括1和180）；指标告警和站点告警时，次数采用枚举值，枚举值分别为：1、2、3、4、5、10、15、30、60、90、120、180。          **默认取值**： 不涉及。
	Count int32 `json:"count"`

	// 告警策略类型，已废弃，不推荐使用。
	Type *string `json:"type,omitempty"`

	SuppressDuration *SuppressDuration `json:"suppress_duration,omitempty"`

	// **参数解释**： 告警级别。     **约束限制**： 不涉及。 **取值范围**： 只能为1、2、3、4。 - 1为紧急 - 2为重要 - 3为次要 - 4为提示           **默认取值**： 不涉及。
	AlarmLevel *int32 `json:"alarm_level,omitempty"`

	// **参数解释**： 用户在页面中选择的指标单位， 用于后续指标数据回显和计算。     **约束限制**： 不涉及。 **取值范围**： 长度为[0,64]个字符。        **默认取值**： 不涉及。
	SelectedUnit *string `json:"selected_unit,omitempty"`
}

func (o PoliciesInListResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoliciesInListResp struct{}"
	}

	return strings.Join([]string{"PoliciesInListResp", string(data)}, " ")
}
