package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type AlarmTemplatePolicies struct {

	// **参数解释**： 查询服务的命名空间，各服务命名空间请参考“[服务命名空间](ces_03_0059.xml)”。 **约束限制**： 不涉及。 **取值范围**： 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_。字符串的长度必须在 3 到 32个字符之间。 **默认取值**： 不涉及。
	Namespace string `json:"namespace"`

	// **参数解释**： 资源维度。     **约束限制**： 事件告警模板DimensionName为空 **取值范围**： 必须以字母开头，多维度用\",\"分隔，只能包含0-9/a-z/A-Z/_/-，每个维度的最大长度为32。目前最大支持4个维度。举例：单维度场景：instance_id；多维度场景：instance_id,disk        **默认取值**： 不涉及。
	DimensionName string `json:"dimension_name"`

	// **参数解释**： 资源的监控指标名称，各服务的指标名称可查看：“[服务指标名称](ces_03_0059.xml)”。 **约束限制**： 不涉及。 **取值范围**： 必须以字母开头，只能包含0-9/a-z/A-Z/_/-。字符长度最短为1，最大为96。如：弹性云服务器中的监控指标cpu_util，表示弹性服务器的CPU使用率；文档数据库中的指标mongo001_command_ps，表示command执行频率。         **默认取值**： 不涉及。
	MetricName string `json:"metric_name"`

	// 告警条件判断周期,单位为秒
	Period AlarmTemplatePoliciesPeriod `json:"period"`

	// 数据聚合方式
	Filter string `json:"filter"`

	// **参数解释**： 阈值符号。     **约束限制**： 指标告警可以使用的阈值符号有>、>=、<、<=、=、!=、cycle_decrease、cycle_increase、cycle_wave； 事件告警可以使用的阈值符号为>、>=、<、<=、=、!=。 **取值范围**： 支持的值为(>|<|>=|<=|=|!=|cycle_decrease|cycle_increase|cycle_wave);cycle_decrease为环比下降,cycle_increase为环比上升,cycle_wave为环比波动。           **默认取值**： 不涉及。
	ComparisonOperator string `json:"comparison_operator"`

	// 告警阈值。单一阈值时value和alarm_level配对使用，当hierarchical_value和value同时使用时以hierarchical_value为准
	Value float32 `json:"value,omitempty"`

	HierarchicalValue *HierarchicalValue `json:"hierarchical_value,omitempty"`

	// 数据的单位字符串，长度不超过32
	Unit string `json:"unit"`

	// **参数解释**： 告警连续触发次数。     **约束限制**： 不涉及。 **取值范围**： 事件告警时参数值为1~180（包括1和180）；指标告警和站点告警时，次数采用枚举值，枚举值分别为：1、2、3、4、5、10、15、30、60、90、120、180。          **默认取值**： 不涉及。
	Count int32 `json:"count"`

	// **参数解释**： 告警级别。     **约束限制**： 不涉及。 **取值范围**： 只能为1、2、3、4。 - 1为紧急 - 2为重要 - 3为次要 - 4为提示           **默认取值**： 不涉及。
	AlarmLevel int32 `json:"alarm_level"`

	// 告警抑制周期，单位为秒，当告警抑制周期为0时，仅发送一次告警
	SuppressDuration AlarmTemplatePoliciesSuppressDuration `json:"suppress_duration"`

	// **参数解释**： 用户在页面中选择的指标单位， 用于后续指标数据回显和计算。     **约束限制**： 不涉及。 **取值范围**： 长度为[0,64]个字符。        **默认取值**： 不涉及。
	SelectedUnit *string `json:"selected_unit,omitempty"`
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
	E_0     AlarmTemplatePoliciesPeriod
	E_1     AlarmTemplatePoliciesPeriod
	E_300   AlarmTemplatePoliciesPeriod
	E_1200  AlarmTemplatePoliciesPeriod
	E_3600  AlarmTemplatePoliciesPeriod
	E_14400 AlarmTemplatePoliciesPeriod
	E_86400 AlarmTemplatePoliciesPeriod
}

func GetAlarmTemplatePoliciesPeriodEnum() AlarmTemplatePoliciesPeriodEnum {
	return AlarmTemplatePoliciesPeriodEnum{
		E_0: AlarmTemplatePoliciesPeriod{
			value: 0,
		}, E_1: AlarmTemplatePoliciesPeriod{
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
