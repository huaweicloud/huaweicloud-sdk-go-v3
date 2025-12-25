package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AlarmHistoryItemV2Condition **参数解释**： 告警触发条件。
type AlarmHistoryItemV2Condition struct {

	// **参数解释**： 指标周期，单位是秒。如想了解各个云服务的指标原始周期可以参考“[支持服务列表](ces_03_0059.xml)”。 **取值范围**： 0是默认值，例如事件类告警该字段就用0即可； 1代表指标的原始周期，比如RDS监控指标原始周期是60s，表示该RDS指标按60s周期为一个数据点参与告警计算； 300代表指标按5分钟聚合周期为一个数据点参与告警计算； 1200代表指标按20分钟聚合周期为一个数据点参与告警计算； 3600代表指标按60分钟聚合周期为一个数据点参与告警计算； 14400代表指标按4小时聚合周期为一个数据点参与告警计算； 86400代表指标按1天聚合周期为一个数据点参与告警计算。
	Period *AlarmHistoryItemV2ConditionPeriod `json:"period,omitempty"`

	// **参数解释**： 聚合方式。 **取值范围**： - average：平均值 - variance：方差 - min：最小值 - max：最大值 - sum：求和
	Filter *string `json:"filter,omitempty"`

	// **参数解释**： 阈值符号。 **取值范围**： 枚举值。支持的值为(>|<|>=|<=|=|!=|cycle_decrease|cycle_increase|cycle_wave);cycle_decrease为环比下降,cycle_increase为环比上升,cycle_wave为环比波动。
	ComparisonOperator *string `json:"comparison_operator,omitempty"`

	// **参数解释**： 告警阈值。 **取值范围**： 具体阈值取值请参见附录中各服务监控指标中取值范围，如[支持监控的服务列表](ces_03_0059.xml)中ECS的CPU使用率cpu_util取值范围可配置80。最小值为-1.7976931348623157e+108，最大值为1.7976931348623157e+108。
	Value *float64 `json:"value,omitempty"`

	// **参数解释**： 数据的单位。 **取值范围**： 字符串长度最大为 32。
	Unit *string `json:"unit,omitempty"`

	// **参数解释**： 告警连续触发次数。 **取值范围**： [1,180]
	Count *int32 `json:"count,omitempty"`

	// **参数解释**： 告警抑制时间（告警周期），单位为秒，对应页面上创建告警规则时告警策略最后一个字段，该字段主要为解决告警频繁的问题。 **取值范围**： 枚举值，只能为0、300、600、900、1800、3600、10800、21600、43200、86400。 - 0：对于指标类告警，0代表告警一次，对于事件类告警，在立即触发场景中，0代表不抑制；在累计触发场景，0代表只告警一次。 - 300代表满足告警触发条件后每5分钟告警一次。 - 600代表满足告警触发条件后每10分钟告警一次。 - 900代表满足告警触发条件后每15分钟告警一次。 - 1800代表满足告警触发条件后每30分钟告警一次。 - 3600代表满足告警触发条件后每60分钟告警一次。 - 10800代表满足告警触发条件后每3小时告警一次。 - 21600代表满足告警触发条件后每6小时告警一次。 - 43200代表满足告警触发条件后每12小时告警一次。 - 86400代表满足告警触发条件后每一天告警一次。
	SuppressDuration *AlarmHistoryItemV2ConditionSuppressDuration `json:"suppress_duration,omitempty"`
}

func (o AlarmHistoryItemV2Condition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmHistoryItemV2Condition struct{}"
	}

	return strings.Join([]string{"AlarmHistoryItemV2Condition", string(data)}, " ")
}

type AlarmHistoryItemV2ConditionPeriod struct {
	value int32
}

type AlarmHistoryItemV2ConditionPeriodEnum struct {
	E_0     AlarmHistoryItemV2ConditionPeriod
	E_1     AlarmHistoryItemV2ConditionPeriod
	E_300   AlarmHistoryItemV2ConditionPeriod
	E_1200  AlarmHistoryItemV2ConditionPeriod
	E_3600  AlarmHistoryItemV2ConditionPeriod
	E_14400 AlarmHistoryItemV2ConditionPeriod
	E_86400 AlarmHistoryItemV2ConditionPeriod
}

func GetAlarmHistoryItemV2ConditionPeriodEnum() AlarmHistoryItemV2ConditionPeriodEnum {
	return AlarmHistoryItemV2ConditionPeriodEnum{
		E_0: AlarmHistoryItemV2ConditionPeriod{
			value: 0,
		}, E_1: AlarmHistoryItemV2ConditionPeriod{
			value: 1,
		}, E_300: AlarmHistoryItemV2ConditionPeriod{
			value: 300,
		}, E_1200: AlarmHistoryItemV2ConditionPeriod{
			value: 1200,
		}, E_3600: AlarmHistoryItemV2ConditionPeriod{
			value: 3600,
		}, E_14400: AlarmHistoryItemV2ConditionPeriod{
			value: 14400,
		}, E_86400: AlarmHistoryItemV2ConditionPeriod{
			value: 86400,
		},
	}
}

func (c AlarmHistoryItemV2ConditionPeriod) Value() int32 {
	return c.value
}

func (c AlarmHistoryItemV2ConditionPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlarmHistoryItemV2ConditionPeriod) UnmarshalJSON(b []byte) error {
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

type AlarmHistoryItemV2ConditionSuppressDuration struct {
	value int32
}

type AlarmHistoryItemV2ConditionSuppressDurationEnum struct {
	E_0     AlarmHistoryItemV2ConditionSuppressDuration
	E_300   AlarmHistoryItemV2ConditionSuppressDuration
	E_600   AlarmHistoryItemV2ConditionSuppressDuration
	E_900   AlarmHistoryItemV2ConditionSuppressDuration
	E_1800  AlarmHistoryItemV2ConditionSuppressDuration
	E_3600  AlarmHistoryItemV2ConditionSuppressDuration
	E_10800 AlarmHistoryItemV2ConditionSuppressDuration
	E_21600 AlarmHistoryItemV2ConditionSuppressDuration
	E_43200 AlarmHistoryItemV2ConditionSuppressDuration
	E_86400 AlarmHistoryItemV2ConditionSuppressDuration
}

func GetAlarmHistoryItemV2ConditionSuppressDurationEnum() AlarmHistoryItemV2ConditionSuppressDurationEnum {
	return AlarmHistoryItemV2ConditionSuppressDurationEnum{
		E_0: AlarmHistoryItemV2ConditionSuppressDuration{
			value: 0,
		}, E_300: AlarmHistoryItemV2ConditionSuppressDuration{
			value: 300,
		}, E_600: AlarmHistoryItemV2ConditionSuppressDuration{
			value: 600,
		}, E_900: AlarmHistoryItemV2ConditionSuppressDuration{
			value: 900,
		}, E_1800: AlarmHistoryItemV2ConditionSuppressDuration{
			value: 1800,
		}, E_3600: AlarmHistoryItemV2ConditionSuppressDuration{
			value: 3600,
		}, E_10800: AlarmHistoryItemV2ConditionSuppressDuration{
			value: 10800,
		}, E_21600: AlarmHistoryItemV2ConditionSuppressDuration{
			value: 21600,
		}, E_43200: AlarmHistoryItemV2ConditionSuppressDuration{
			value: 43200,
		}, E_86400: AlarmHistoryItemV2ConditionSuppressDuration{
			value: 86400,
		},
	}
}

func (c AlarmHistoryItemV2ConditionSuppressDuration) Value() int32 {
	return c.value
}

func (c AlarmHistoryItemV2ConditionSuppressDuration) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlarmHistoryItemV2ConditionSuppressDuration) UnmarshalJSON(b []byte) error {
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
