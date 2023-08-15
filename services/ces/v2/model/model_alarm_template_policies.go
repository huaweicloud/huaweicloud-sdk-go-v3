package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AlarmTemplatePolicies struct {

	// 查询服务的命名空间，各服务命名空间请参考“[服务命名空间](ces_03_0059.xml)”
	Namespace string `json:"namespace"`

	// 资源维度，必须以字母开头，多维度用\",\"分割，只能包含0-9/a-z/A-Z/_/-，每个维度的最大长度为32
	DimensionName string `json:"dimension_name"`

	// 资源的监控指标名称，必须以字母开头，只能包含0-9/a-z/A-Z/_，字符长度最短为1，最大为64；如：弹性云服务器中的监控指标cpu_util，表示弹性服务器的CPU使用率；文档数据库中的指标mongo001_command_ps，表示command执行频率；各服务的指标名称可查看：“[服务指标名称](https://support.huaweicloud.com/usermanual-ces/zh-cn_topic_0202622212.html)”。
	MetricName string `json:"metric_name"`

	// 告警条件判断周期,单位为秒
	Period AlarmTemplatePoliciesPeriod `json:"period"`

	// 数据聚合方式
	Filter string `json:"filter"`

	// 告警阈值的比较条件
	ComparisonOperator string `json:"comparison_operator"`

	// 告警阈值
	Value float32 `json:"value"`

	// 数据的单位字符串，长度不超过32
	Unit string `json:"unit"`

	// 告警连续触发次数，正整数[1, 5]
	Count int32 `json:"count"`

	// 告警级别，1为紧急，2为重要，3为次要，4为提示
	AlarmLevel int32 `json:"alarm_level"`

	// 告警抑制周期，单位为秒，当告警抑制周期为0时，仅发送一次告警
	SuppressDuration AlarmTemplatePoliciesSuppressDuration `json:"suppress_duration"`
}

func (o AlarmTemplatePolicies) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AlarmTemplatePolicies struct{}"
	}

	return strings.Join([]string{"AlarmTemplatePolicies", string(data)}, " ")
}

type AlarmTemplatePoliciesPeriod struct {
	value int32
}

type AlarmTemplatePoliciesPeriodEnum struct {
	E_1     AlarmTemplatePoliciesPeriod
	E_300   AlarmTemplatePoliciesPeriod
	E_1200  AlarmTemplatePoliciesPeriod
	E_3600  AlarmTemplatePoliciesPeriod
	E_14400 AlarmTemplatePoliciesPeriod
	E_86400 AlarmTemplatePoliciesPeriod
}

func GetAlarmTemplatePoliciesPeriodEnum() AlarmTemplatePoliciesPeriodEnum {
	return AlarmTemplatePoliciesPeriodEnum{
		E_1: AlarmTemplatePoliciesPeriod{
			value: 1,
		}, E_300: AlarmTemplatePoliciesPeriod{
			value: 300,
		}, E_1200: AlarmTemplatePoliciesPeriod{
			value: 1200,
		}, E_3600: AlarmTemplatePoliciesPeriod{
			value: 3600,
		}, E_14400: AlarmTemplatePoliciesPeriod{
			value: 14400,
		}, E_86400: AlarmTemplatePoliciesPeriod{
			value: 86400,
		},
	}
}

func (c AlarmTemplatePoliciesPeriod) Value() int32 {
	return c.value
}

func (c AlarmTemplatePoliciesPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlarmTemplatePoliciesPeriod) UnmarshalJSON(b []byte) error {
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

type AlarmTemplatePoliciesSuppressDuration struct {
	value int32
}

type AlarmTemplatePoliciesSuppressDurationEnum struct {
	E_0     AlarmTemplatePoliciesSuppressDuration
	E_300   AlarmTemplatePoliciesSuppressDuration
	E_600   AlarmTemplatePoliciesSuppressDuration
	E_900   AlarmTemplatePoliciesSuppressDuration
	E_1800  AlarmTemplatePoliciesSuppressDuration
	E_3600  AlarmTemplatePoliciesSuppressDuration
	E_10800 AlarmTemplatePoliciesSuppressDuration
	E_21600 AlarmTemplatePoliciesSuppressDuration
	E_43200 AlarmTemplatePoliciesSuppressDuration
	E_86400 AlarmTemplatePoliciesSuppressDuration
}

func GetAlarmTemplatePoliciesSuppressDurationEnum() AlarmTemplatePoliciesSuppressDurationEnum {
	return AlarmTemplatePoliciesSuppressDurationEnum{
		E_0: AlarmTemplatePoliciesSuppressDuration{
			value: 0,
		}, E_300: AlarmTemplatePoliciesSuppressDuration{
			value: 300,
		}, E_600: AlarmTemplatePoliciesSuppressDuration{
			value: 600,
		}, E_900: AlarmTemplatePoliciesSuppressDuration{
			value: 900,
		}, E_1800: AlarmTemplatePoliciesSuppressDuration{
			value: 1800,
		}, E_3600: AlarmTemplatePoliciesSuppressDuration{
			value: 3600,
		}, E_10800: AlarmTemplatePoliciesSuppressDuration{
			value: 10800,
		}, E_21600: AlarmTemplatePoliciesSuppressDuration{
			value: 21600,
		}, E_43200: AlarmTemplatePoliciesSuppressDuration{
			value: 43200,
		}, E_86400: AlarmTemplatePoliciesSuppressDuration{
			value: 86400,
		},
	}
}

func (c AlarmTemplatePoliciesSuppressDuration) Value() int32 {
	return c.value
}

func (c AlarmTemplatePoliciesSuppressDuration) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AlarmTemplatePoliciesSuppressDuration) UnmarshalJSON(b []byte) error {
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
