package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type Policies struct {

	// **参数解释**： 查询服务的命名空间，各服务命名空间请参考“[服务命名空间](ces_03_0059.xml)”。 **约束限制**： 不涉及。 **取值范围**： 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_。字符串的长度必须在 3 到 32个字符之间。 **默认取值**： 不涉及。
	Namespace string `json:"namespace"`

	// **参数解释**： 资源维度。     **约束限制**： 事件告警模板DimensionName为空 **取值范围**： 必须以字母开头，多维度用\",\"分隔，只能包含0-9/a-z/A-Z/_/-，每个维度的最大长度为32。目前最大支持4个维度。举例：单维度场景：instance_id；多维度场景：instance_id,disk        **默认取值**： 不涉及。
	DimensionName *string `json:"dimension_name,omitempty"`

	// **参数解释**： 资源的监控指标名称，各服务的指标名称可查看：“[服务指标名称](ces_03_0059.xml)”。 **约束限制**： 不涉及。 **取值范围**： 必须以字母开头，只能包含0-9/a-z/A-Z/_/-。字符长度最短为1，最大为96。如：弹性云服务器中的监控指标cpu_util，表示弹性服务器的CPU使用率；文档数据库中的指标mongo001_command_ps，表示command执行频率。         **默认取值**： 不涉及。
	MetricName string `json:"metric_name"`

	ExtraInfo *MetricExtraInfo `json:"extra_info,omitempty"`

	// **参数解释**： 告警条件判断周期,单位为秒。 **约束限制**： 不涉及。 **取值范围**： 枚举值。 - 0是默认值，事件类告警该字段用0即可。 - 1代表指标的原始周期，比如RDS监控指标原始周期是60s，表示该RDS指标按60s周期为一个数据点参与告警计算。 - 300代表指标按5分钟聚合周期为一个数据点参与告警计算。 - 1200代表指标按20分钟聚合周期为一个数据点参与告警计算。 - 3600代表指标按1小时聚合周期为一个数据点参与告警计算。 - 14400代表指标按4小时聚合周期为一个数据点参与告警计算。 - 86400代表指标按1天聚合周期为一个数据点参与告警计算。          **默认取值**： 不涉及。
	Period PoliciesPeriod `json:"period"`

	// **参数解释**： 聚合方式。         **约束限制**： 不涉及。 **取值范围**： average： 平均值，variance：方差，min：最小值，max：最大值，sum：求和。           **默认取值**： 不涉及。
	Filter string `json:"filter"`

	// **参数解释**： 阈值符号。     **约束限制**： 指标告警可以使用的阈值符号有>、>=、<、<=、=、!=、cycle_decrease、cycle_increase、cycle_wave； 事件告警可以使用的阈值符号为>、>=、<、<=、=、!=。 **取值范围**： 支持的值为(>|<|>=|<=|=|!=|cycle_decrease|cycle_increase|cycle_wave);cycle_decrease为环比下降,cycle_increase为环比上升,cycle_wave为环比波动。           **默认取值**： 不涉及。
	ComparisonOperator string `json:"comparison_operator"`

	// **参数解释**： 告警阈值。 **约束限制**： 单一阈值时value和alarm_level配对使用，当hierarchical_value和value同时使用时以hierarchical_value为准。 **取值范围**： 不涉及。           **默认取值**： 不涉及。
	Value float32 `json:"value,omitempty"`

	HierarchicalValue *HierarchicalValue `json:"hierarchical_value,omitempty"`

	// **参数解释**： 数据的单位字符串，长度不超过32。 **约束限制**： 不涉及。 **取值范围**： 长度为[0,32]个字符    **默认取值**： 不涉及。
	Unit *string `json:"unit,omitempty"`

	// **参数解释**： 用户在页面中选择的指标单位， 用于后续指标数据回显和计算。     **约束限制**： 不涉及。 **取值范围**： 长度为[0,64]个字符。        **默认取值**： 不涉及。
	SelectedUnit *string `json:"selected_unit,omitempty"`

	// **参数解释**： 告警连续触发次数。     **约束限制**： 不涉及。 **取值范围**： 事件告警时参数值为1~180（包括1和180）；指标告警和站点告警时，次数采用枚举值，枚举值分别为：1、2、3、4、5、10、15、30、60、90、120、180。          **默认取值**： 不涉及。
	Count int32 `json:"count"`

	// **参数解释**： 告警策略类型，已废弃，不推荐使用。 **约束限制**： 不涉及。 **取值范围**： 只能为auto。          **默认取值**： 不涉及。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 告警级别。     **约束限制**： 不涉及。 **取值范围**： 只能为1、2、3、4。 - 1为紧急 - 2为重要 - 3为次要 - 4为提示           **默认取值**： 不涉及。
	AlarmLevel *int32 `json:"alarm_level,omitempty"`

	// **参数解释**： 告警抑制周期，单位为秒。 **约束限制**： 不涉及。 **取值范围**： 枚举值，只能为0、300、600、900、1800、3600、10800、21600、43200、86400。 - 0：在立即触发场景中，0代表不抑制；在累计触发场景，0代表只告警一次。 - 300代表满足告警触发条件后每5分钟告警一次。 - 600代表满足告警触发条件后每10分钟告警一次。 - 900代表满足告警触发条件后每15分钟告警一次。 - 1800代表满足告警触发条件后每30分钟告警一次。 - 3600代表满足告警触发条件后每60分钟告警一次。 - 10800代表满足告警触发条件后每3小时告警一次。 - 21600代表满足告警触发条件后每6小时告警一次。 - 43200代表满足告警触发条件后每12小时告警一次。 - 86400代表满足告警触发条件后每一天告警一次。 **默认取值**： 不涉及。
	SuppressDuration *PoliciesSuppressDuration `json:"suppress_duration,omitempty"`
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
