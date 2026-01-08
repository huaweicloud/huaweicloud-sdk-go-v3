package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// AlarmRuleConditionResp **参数解释**： 告警规则设置的告警策略。
type AlarmRuleConditionResp struct {

	// **参数解释**： 阈值符号。     **取值范围**： 支持的值为(>|<|>=|<=|=|!=|cycle_decrease|cycle_increase|cycle_wave);cycle_decrease为环比下降,cycle_increase为环比上升,cycle_wave为环比波动。
	ComparisonOperator *string `json:"comparison_operator,omitempty"`

	// **参数解释**： 触发告警的连续发生次数。 **取值范围**： 整数，取值范围[1, 5]。
	Count *int32 `json:"count,omitempty"`

	Filter *FilterResp `json:"filter,omitempty"`

	Period *PeriodResp `json:"period,omitempty"`

	// **参数解释**： 数据的原始单位。 **取值范围**： 长度为[0,32]个字符。
	Unit *string `json:"unit,omitempty"`

	// **参数解释**： 用户选择的数据单位。 **取值范围**： 长度为[0,32]个字符。
	SelectedUnit *string `json:"selected_unit,omitempty"`

	// **参数解释**： 告警阈值。具体阈值取值请参见附录中各服务监控指标中取值范围，如[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)中ECS的CPU使用率cpu_util取值范围可配置80。 **取值范围**： 最小值为-1.7976931348623157e+108，最大值为1.7976931348623157e+108。
	Value *float64 `json:"value,omitempty"`

	SuppressDuration *SuppressDurationResp `json:"suppress_duration,omitempty"`
}

func (o AlarmRuleConditionResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmRuleConditionResp struct{}"
	}

	return strings.Join([]string{"AlarmRuleConditionResp", string(data)}, " ")
}
