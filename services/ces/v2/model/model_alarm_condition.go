package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AlarmCondition 告警触发条件
type AlarmCondition struct {

	// 指标周期，单位是秒； 0是默认值，例如事件类告警该字段就用0即可； 1代表指标的原始周期，比如RDS监控指标原始周期是60s，表示该RDS指标按60s周期为一个数据点参与告警计算；如想了解各个云服务的指标原始周期可以参考[服务命名空间](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)， 300代表指标按5分钟聚合周期为一个数据点参与告警计算。
	Period AlarmConditionPeriod `json:"period"`

	// 聚合方式, 支持的值为(average|min|max|sum)
	Filter string `json:"filter"`

	// 阈值符号
	ComparisonOperator string `json:"comparison_operator"`

	// 告警阈值，取值范围[0, Number.MAX_VALUE]，Number.MAX_VALUE值为1.7976931348623157e+108。具体阈值取值请参见附录中各服务监控指标中取值范围，如支持监控的服务列表中ECS的CPU使用率cpu_util取值范围可配置80。
	Value float64 `json:"value"`

	// 数据的单位，最大长度为32位。
	Unit *string `json:"unit,omitempty"`

	// 次数
	Count int32 `json:"count"`

	// 告警抑制时间，单位为秒，对应页面上创建告警规则时告警策略最后一个字段，该字段主要为解决告警频繁的问题，0代表不抑制，满足条件即告警；300代表满足告警触发条件后每5分钟告警一次；
	SuppressDuration *AlarmConditionSuppressDuration `json:"suppress_duration,omitempty"`
}

func (o AlarmCondition) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmCondition struct{}"
	}

	return strings.Join([]string{"AlarmCondition", string(data)}, " ")
}

type AlarmConditionPeriod struct {
	value int32
}

type AlarmConditionPeriodEnum struct {
	E_0     AlarmConditionPeriod
	E_1     AlarmConditionPeriod
	E_300   AlarmConditionPeriod
	E_1200  AlarmConditionPeriod
	E_3600  AlarmConditionPeriod
	E_14400 AlarmConditionPeriod
	E_86400 AlarmConditionPeriod
}

func GetAlarmConditionPeriodEnum() AlarmConditionPeriodEnum {
	return AlarmConditionPeriodEnum{
		E_0: AlarmConditionPeriod{
			value: 0,
		}, E_1: AlarmConditionPeriod{
			value: 1,
		}, E_300: AlarmConditionPeriod{
			value: 300,
		}, E_1200: AlarmConditionPeriod{
			value: 1200,
		}, E_3600: AlarmConditionPeriod{
			value: 3600,
		}, E_14400: AlarmConditionPeriod{
			value: 14400,
		}, E_86400: AlarmConditionPeriod{
			value: 86400,
		},
	}
}

func (c AlarmConditionPeriod) Value() int32 {
	return c.value
}

func (c AlarmConditionPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlarmConditionPeriod) UnmarshalJSON(b []byte) error {
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

type AlarmConditionSuppressDuration struct {
	value int32
}

type AlarmConditionSuppressDurationEnum struct {
	E_0     AlarmConditionSuppressDuration
	E_300   AlarmConditionSuppressDuration
	E_600   AlarmConditionSuppressDuration
	E_900   AlarmConditionSuppressDuration
	E_1800  AlarmConditionSuppressDuration
	E_3600  AlarmConditionSuppressDuration
	E_10800 AlarmConditionSuppressDuration
	E_21600 AlarmConditionSuppressDuration
	E_43200 AlarmConditionSuppressDuration
}

func GetAlarmConditionSuppressDurationEnum() AlarmConditionSuppressDurationEnum {
	return AlarmConditionSuppressDurationEnum{
		E_0: AlarmConditionSuppressDuration{
			value: 0,
		}, E_300: AlarmConditionSuppressDuration{
			value: 300,
		}, E_600: AlarmConditionSuppressDuration{
			value: 600,
		}, E_900: AlarmConditionSuppressDuration{
			value: 900,
		}, E_1800: AlarmConditionSuppressDuration{
			value: 1800,
		}, E_3600: AlarmConditionSuppressDuration{
			value: 3600,
		}, E_10800: AlarmConditionSuppressDuration{
			value: 10800,
		}, E_21600: AlarmConditionSuppressDuration{
			value: 21600,
		}, E_43200: AlarmConditionSuppressDuration{
			value: 43200,
		},
	}
}

func (c AlarmConditionSuppressDuration) Value() int32 {
	return c.value
}

func (c AlarmConditionSuppressDuration) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlarmConditionSuppressDuration) UnmarshalJSON(b []byte) error {
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
