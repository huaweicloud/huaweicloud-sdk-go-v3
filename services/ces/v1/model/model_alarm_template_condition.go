package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AlarmTemplateCondition 创建自定义告警模板的告警策略。
type AlarmTemplateCondition struct {

	// **参数解释**： 告警阈值的比较条件。 **约束限制**： 不涉及。 **取值范围**： 只能是>、=、<、>=、<=、!=。 **默认取值**： 不涉及。
	ComparisonOperator string `json:"comparison_operator"`

	// **参数解释**： 触发告警的连续发生次数。 **约束限制**： 不涉及。 **取值范围**： 取值范围[1, 5]。告警类型为事件告警时，取值范围为[1, 100]。 **默认取值**： 不涉及。
	Count int32 `json:"count"`

	Filter *Filter `json:"filter"`

	// **参数解释**： 指标周期，单位是秒。如想了解各个云服务的指标原始周期可以参考“[支持服务列表](ces_03_0059.xml)” **约束限制**： 不涉及。 **取值范围**： 枚举值。 - 0代表立即触发，仅限事件场景使用。 - 1代表指标的原始周期，比如RDS监控指标原始周期是60s，表示该RDS指标按60s周期为一个数据点参与告警计算。 - 300代表指标按5分钟聚合周期为一个数据点参与告警计算。 - 1200代表指标按20分钟聚合周期为一个数据点参与告警计算。 - 3600代表指标按1小时聚合周期为一个数据点参与告警计算。 - 14400代表指标按4小时聚合周期为一个数据点参与告警计算。 - 86400代表指标按1天聚合周期为一个数据点参与告警计算。 **默认取值**：  不涉及。
	Period AlarmTemplateConditionPeriod `json:"period"`

	// **参数解释**： 数据的单位。 **约束限制**： 不涉及。 **取值范围**： 长度为[0,32]个字符。 **默认取值**： 不涉及。
	Unit *string `json:"unit,omitempty"`

	// **参数解释**： 告警阈值。具体阈值取值请参见附录中各服务监控指标中取值范围，如[支持监控的服务列表](ces_03_0059.xml)中ECS的CPU使用率cpu_util取值范围可配置80。 **约束限制**： 单一阈值时value和alarm_level配对使用，当hierarchical_value和value同时使用时以hierarchical_value为准。 **取值范围**： 最小值为-1.7976931348623157e+108，最大值为1.7976931348623157e+108。           **默认取值**： 不涉及。
	Value float64 `json:"value"`

	// **参数解释**： 告警抑制时间，单位为秒，对应页面上创建告警规则时告警策略最后一个字段，该字段主要为解决告警频繁的问题。 **约束限制**： 不涉及。 **取值范围**： 枚举值，只能为0、300、600、900、1800、3600、10800、21600、43200、86400。 - 0：对于指标类告警，0代表告警一次。对于事件类告警，在立即触发场景中，0代表不抑制；在累计触发场景，0代表只告警一次。 - 300代表满足告警触发条件后每5分钟告警一次。 - 600代表满足告警触发条件后每10分钟告警一次。 - 900代表满足告警触发条件后每15分钟告警一次。 - 1800代表满足告警触发条件后每30分钟告警一次。 - 3600代表满足告警触发条件后每60分钟告警一次。 - 10800代表满足告警触发条件后每3小时告警一次。 - 21600代表满足告警触发条件后每6小时告警一次。 - 43200代表满足告警触发条件后每12小时告警一次。 - 86000代表满足告警触发条件后每一天告警一次。 **默认取值**： 不涉及。
	SuppressDuration *AlarmTemplateConditionSuppressDuration `json:"suppress_duration,omitempty"`
}

func (o AlarmTemplateCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmTemplateCondition struct{}"
	}

	return strings.Join([]string{"AlarmTemplateCondition", string(data)}, " ")
}

type AlarmTemplateConditionPeriod struct {
	value int32
}

type AlarmTemplateConditionPeriodEnum struct {
	E_0     AlarmTemplateConditionPeriod
	E_1     AlarmTemplateConditionPeriod
	E_300   AlarmTemplateConditionPeriod
	E_1200  AlarmTemplateConditionPeriod
	E_3600  AlarmTemplateConditionPeriod
	E_14400 AlarmTemplateConditionPeriod
	E_86400 AlarmTemplateConditionPeriod
}

func GetAlarmTemplateConditionPeriodEnum() AlarmTemplateConditionPeriodEnum {
	return AlarmTemplateConditionPeriodEnum{
		E_0: AlarmTemplateConditionPeriod{
			value: 0,
		}, E_1: AlarmTemplateConditionPeriod{
			value: 1,
		}, E_300: AlarmTemplateConditionPeriod{
			value: 300,
		}, E_1200: AlarmTemplateConditionPeriod{
			value: 1200,
		}, E_3600: AlarmTemplateConditionPeriod{
			value: 3600,
		}, E_14400: AlarmTemplateConditionPeriod{
			value: 14400,
		}, E_86400: AlarmTemplateConditionPeriod{
			value: 86400,
		},
	}
}

func (c AlarmTemplateConditionPeriod) Value() int32 {
	return c.value
}

func (c AlarmTemplateConditionPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlarmTemplateConditionPeriod) UnmarshalJSON(b []byte) error {
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

type AlarmTemplateConditionSuppressDuration struct {
	value int32
}

type AlarmTemplateConditionSuppressDurationEnum struct {
	E_0     AlarmTemplateConditionSuppressDuration
	E_300   AlarmTemplateConditionSuppressDuration
	E_600   AlarmTemplateConditionSuppressDuration
	E_900   AlarmTemplateConditionSuppressDuration
	E_1800  AlarmTemplateConditionSuppressDuration
	E_3600  AlarmTemplateConditionSuppressDuration
	E_10800 AlarmTemplateConditionSuppressDuration
	E_21600 AlarmTemplateConditionSuppressDuration
	E_43200 AlarmTemplateConditionSuppressDuration
	E_86400 AlarmTemplateConditionSuppressDuration
}

func GetAlarmTemplateConditionSuppressDurationEnum() AlarmTemplateConditionSuppressDurationEnum {
	return AlarmTemplateConditionSuppressDurationEnum{
		E_0: AlarmTemplateConditionSuppressDuration{
			value: 0,
		}, E_300: AlarmTemplateConditionSuppressDuration{
			value: 300,
		}, E_600: AlarmTemplateConditionSuppressDuration{
			value: 600,
		}, E_900: AlarmTemplateConditionSuppressDuration{
			value: 900,
		}, E_1800: AlarmTemplateConditionSuppressDuration{
			value: 1800,
		}, E_3600: AlarmTemplateConditionSuppressDuration{
			value: 3600,
		}, E_10800: AlarmTemplateConditionSuppressDuration{
			value: 10800,
		}, E_21600: AlarmTemplateConditionSuppressDuration{
			value: 21600,
		}, E_43200: AlarmTemplateConditionSuppressDuration{
			value: 43200,
		}, E_86400: AlarmTemplateConditionSuppressDuration{
			value: 86400,
		},
	}
}

func (c AlarmTemplateConditionSuppressDuration) Value() int32 {
	return c.value
}

func (c AlarmTemplateConditionSuppressDuration) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlarmTemplateConditionSuppressDuration) UnmarshalJSON(b []byte) error {
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
