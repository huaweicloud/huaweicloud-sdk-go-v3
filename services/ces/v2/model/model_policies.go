package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Policies struct {

	// 查询服务的命名空间，各服务命名空间请参考“[服务命名空间](ces_03_0059.xml)”
	Namespace string `json:"namespace"`

	// 资源维度，必须以字母开头，多维度用\",\"分割，只能包含0-9/a-z/A-Z/_/-，每个维度的最大长度为32, 事件告警模板DimensionName为空
	DimensionName *string `json:"dimension_name,omitempty"`

	// 资源的监控指标名称，必须以字母开头，只能包含0-9/a-z/A-Z/_，字符长度最短为1，最大为64；如：弹性云服务器中的监控指标cpu_util，表示弹性服务器的CPU使用率；文档数据库中的指标mongo001_command_ps，表示command执行频率；各服务的指标名称可查看：“[服务指标名称](ces_03_0059.xml)”。
	MetricName string `json:"metric_name"`

	// 告警条件判断周期,单位为秒
	Period PoliciesPeriod `json:"period"`

	// 聚合方式。average： 平均值，variance：方差，min：最小值，max：最大值，sum：求和，tp99：99百分位数，tp95：95百分位数，tp90：90百分位数
	Filter string `json:"filter"`

	// 阈值符号, 支持的值为(>|<|>=|<=|=|!=|cycle_decrease|cycle_increase|cycle_wave);cycle_decrease为环比下降,cycle_increase为环比上升,cycle_wave为环比波动； 指标告警可以使用的阈值符号有>、>=、<、<=、=、!=、cycle_decrease、cycle_increase、cycle_wave； 事件告警可以使用的阈值符号为>、>=、<、<=、=、!=；
	ComparisonOperator string `json:"comparison_operator"`

	// 告警阈值。单一阈值时value和alarm_level配对使用，当hierarchical_value和value同时使用时以hierarchical_value为准。
	Value float32 `json:"value,omitempty"`

	HierarchicalValue *HierarchicalValue `json:"hierarchical_value,omitempty"`

	// 数据的单位字符串，长度不超过32
	Unit *string `json:"unit,omitempty"`

	// 告警连续触发次数，事件告警时参数值为1~180（包括1和180）；指标告警和站点告警时，次数采用枚举值，枚举值分别为：1、2、3、4、5、10、15、30、60、90、120、180
	Count int32 `json:"count"`

	// 告警级别，1为紧急，2为重要，3为次要，4为提示
	AlarmLevel *int32 `json:"alarm_level,omitempty"`

	// 告警抑制周期，单位为秒，当告警抑制周期为0时，仅发送一次告警
	SuppressDuration PoliciesSuppressDuration `json:"suppress_duration"`
}

func (o Policies) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Policies struct{}"
	}

	return strings.Join([]string{"Policies", string(data)}, " ")
}

type PoliciesPeriod struct {
	value int32
}

type PoliciesPeriodEnum struct {
	E_0     PoliciesPeriod
	E_1     PoliciesPeriod
	E_300   PoliciesPeriod
	E_1200  PoliciesPeriod
	E_3600  PoliciesPeriod
	E_14400 PoliciesPeriod
	E_86400 PoliciesPeriod
}

func GetPoliciesPeriodEnum() PoliciesPeriodEnum {
	return PoliciesPeriodEnum{
		E_0: PoliciesPeriod{
			value: 0,
		}, E_1: PoliciesPeriod{
			value: 1,
		}, E_300: PoliciesPeriod{
			value: 300,
		}, E_1200: PoliciesPeriod{
			value: 1200,
		}, E_3600: PoliciesPeriod{
			value: 3600,
		}, E_14400: PoliciesPeriod{
			value: 14400,
		}, E_86400: PoliciesPeriod{
			value: 86400,
		},
	}
}

func (c PoliciesPeriod) Value() int32 {
	return c.value
}

func (c PoliciesPeriod) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoliciesPeriod) UnmarshalJSON(b []byte) error {
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

type PoliciesSuppressDuration struct {
	value int32
}

type PoliciesSuppressDurationEnum struct {
	E_0     PoliciesSuppressDuration
	E_300   PoliciesSuppressDuration
	E_600   PoliciesSuppressDuration
	E_900   PoliciesSuppressDuration
	E_1800  PoliciesSuppressDuration
	E_3600  PoliciesSuppressDuration
	E_10800 PoliciesSuppressDuration
	E_21600 PoliciesSuppressDuration
	E_43200 PoliciesSuppressDuration
	E_86400 PoliciesSuppressDuration
}

func GetPoliciesSuppressDurationEnum() PoliciesSuppressDurationEnum {
	return PoliciesSuppressDurationEnum{
		E_0: PoliciesSuppressDuration{
			value: 0,
		}, E_300: PoliciesSuppressDuration{
			value: 300,
		}, E_600: PoliciesSuppressDuration{
			value: 600,
		}, E_900: PoliciesSuppressDuration{
			value: 900,
		}, E_1800: PoliciesSuppressDuration{
			value: 1800,
		}, E_3600: PoliciesSuppressDuration{
			value: 3600,
		}, E_10800: PoliciesSuppressDuration{
			value: 10800,
		}, E_21600: PoliciesSuppressDuration{
			value: 21600,
		}, E_43200: PoliciesSuppressDuration{
			value: 43200,
		}, E_86400: PoliciesSuppressDuration{
			value: 86400,
		},
	}
}

func (c PoliciesSuppressDuration) Value() int32 {
	return c.value
}

func (c PoliciesSuppressDuration) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoliciesSuppressDuration) UnmarshalJSON(b []byte) error {
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
