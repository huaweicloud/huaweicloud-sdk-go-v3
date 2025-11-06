package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ConditionResp **参数解释**： 告警规则设置的告警策略。
type ConditionResp struct {

	// **参数解释**： 告警阈值的比较条件。 **取值范围**： 只能是>、=、<、>=、<=、!=。
	ComparisonOperator *string `json:"comparison_operator,omitempty"`

	// **参数解释**： 触发告警的连续发生次数。 **取值范围**： 取值范围[1, 5]。告警类型为事件告警时，取值范围为[1, 100]。
	Count *int32 `json:"count,omitempty"`

	Filter *FilterResp `json:"filter,omitempty"`

	Period *PeriodResp `json:"period,omitempty"`

	// **参数解释**： 数据的单位。 **取值范围**： 长度为[0,32]个字符。
	Unit *string `json:"unit,omitempty"`

	// **参数解释**： 告警阈值。具体阈值取值请参见附录中各服务监控指标中取值范围，如[支持监控的服务列表](ces_03_0059.xml)中ECS的CPU使用率cpu_util取值范围可配置80。 **取值范围**： 最小值为-1.7976931348623157e+108，最大值为1.7976931348623157e+108。
	Value *float64 `json:"value,omitempty"`

	SuppressDuration *SuppressDurationResp `json:"suppress_duration,omitempty"`
}

func (o ConditionResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ConditionResp struct{}"
	}

	return strings.Join([]string{"ConditionResp", string(data)}, " ")
}
