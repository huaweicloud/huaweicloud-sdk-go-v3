package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Condition **参数解释**： 告警规则设置的告警策略。 **约束限制**： 不涉及。
type Condition struct {

	// **参数解释**： 阈值符号。 **约束限制**： 指标告警可以使用的阈值符号有>、>=、<、<=、=、!=、cycle_decrease、cycle_increase、cycle_wave； 事件告警可以使用的阈值符号为>、>=、<、<=、=、!=。 **取值范围**： 支持的值为(>|<|>=|<=|=|!=|cycle_decrease|cycle_increase|cycle_wave);cycle_decrease为环比下降,cycle_increase为环比上升,cycle_wave为环比波动。
	ComparisonOperator string `json:"comparison_operator"`

	// **参数解释**： 触发告警的连续发生次数。 **约束限制**： 不涉及。 **取值范围**： 整数，取值范围[1, 5]。 **默认取值**： 不涉及。
	Count int32 `json:"count"`

	Filter *Filter `json:"filter"`

	// **参数解释**： 指标周期，单位是秒。如想了解各个云服务的指标原始周期可以参考[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)。  **约束限制**： 不涉及。 **取值范围**：  枚举值。 - 0代表立即触发，仅限事件场景使用。 - 1代表指标的原始周期，比如RDS监控指标原始周期是60s，表示该RDS指标按60s周期为一个数据点参与告警计算。 - 300代表指标按5分钟聚合周期为一个数据点参与告警计算。 - 1200代表指标按20分钟聚合周期为一个数据点参与告警计算。 - 3600代表指标按1小时聚合周期为一个数据点参与告警计算。 - 14400代表指标按4小时聚合周期为一个数据点参与告警计算。 - 86400代表指标按1天聚合周期为一个数据点参与告警计算。  **默认取值**：  不涉及。
	Period ConditionPeriod `json:"period"`

	// **参数解释**： 数据的单位。 **约束限制**： 不涉及。 **取值范围**： 长度为[0,32]个字符。 **默认取值**： 不涉及。
	Unit *string `json:"unit,omitempty"`

	// **参数解释**： 告警阈值。具体阈值取值请参见附录中各服务监控指标中取值范围，如[[支持监控的服务列表](https://support.huaweicloud.com/api-ces/ces_03_0059.html)](tag:hc)[[支持监控的服务列表](https://support.huaweicloud.com/intl/en-us/api-ces/ces_03_0059.html)](tag:hk)[[支持监控的服务列表](https://support.huaweicloud.com/eu/en-us/api-ces/ces_03_0059.html)](tag:hws_eu)[[支持监控的服务列表](ces_03_0059.xml)](tag:ax,cmcc,ctc,dt,dt_test,hcso_dt,fcs,fcs_vm,mix,g42,hk_g42,hk_sbc,hk_tm,hk_vdf,hws_ocb,ocb,sbc,srg)中ECS的CPU使用率cpu_util取值范围可配置80。 **约束限制**： 单一阈值时value和alarm_level配对使用，当hierarchical_value和value同时使用时以hierarchical_value为准。 **取值范围**： 最小值为-1.7976931348623157e+108，最大值为1.7976931348623157e+108。           **默认取值**： 不涉及。
	Value float64 `json:"value"`

	SuppressDuration *SuppressDuration `json:"suppress_duration,omitempty"`
}

func (o Condition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Condition struct{}"
	}

	return strings.Join([]string{"Condition", string(data)}, " ")
}

type ConditionPeriod struct {
	value int32
}

type ConditionPeriodEnum struct {
	E_0     ConditionPeriod
	E_1     ConditionPeriod
	E_300   ConditionPeriod
	E_1200  ConditionPeriod
	E_3600  ConditionPeriod
	E_14400 ConditionPeriod
	E_86400 ConditionPeriod
}

func GetConditionPeriodEnum() ConditionPeriodEnum {
	return ConditionPeriodEnum{
		E_0: ConditionPeriod{
			value: 0,
		}, E_1: ConditionPeriod{
			value: 1,
		}, E_300: ConditionPeriod{
			value: 300,
		}, E_1200: ConditionPeriod{
			value: 1200,
		}, E_3600: ConditionPeriod{
			value: 3600,
		}, E_14400: ConditionPeriod{
			value: 14400,
		}, E_86400: ConditionPeriod{
			value: 86400,
		},
	}
}

func (c ConditionPeriod) Value() int32 {
	return c.value
}

func (c ConditionPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ConditionPeriod) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}
