package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePolicy struct {

	// 资源的监控指标名称，必须以字母开头，只能包含0-9/a-z/A-Z/_，字符长度最短为1，最大为64；如：弹性云服务器中的监控指标cpu_util，表示弹性服务器的CPU使用率；文档数据库中的指标mongo001_command_ps，表示command执行频率；各服务的指标名称可查看：“[服务指标名称](ces_03_0059.xml)”。
	MetricName string `json:"metric_name"`

	ExtraInfo *ExtraInfo `json:"extra_info,omitempty"`

	// 指标周期，单位是秒； 0是默认值，例如事件类告警该字段就用0即可； 1代表指标的原始周期，比如RDS监控指标原始周期是60s，表示该RDS指标按60s周期为一个数据点参与告警计算；如想了解各个云服务的指标原始周期可以参考“[支持服务列表](ces_03_0059.xml)”，300代表指标按5分钟聚合周期为一个数据点参与告警计算。
	Period int32 `json:"period"`

	// 聚合方式, 支持的值为(average|min|max|sum)
	Filter string `json:"filter"`

	// 阈值符号, 支持的值为(>|<|>=|<=|=|!=|cycle_decrease|cycle_increase|cycle_wave);cycle_decrease为环比下降,cycle_increase为环比上升,cycle_wave为环比波动； 指标告警可以使用的阈值符号有>、>=、<、<=、=、!=、cycle_decrease、cycle_increase、cycle_wave； 事件告警可以使用的阈值符号为>、>=、<、<=、=、!=；
	ComparisonOperator string `json:"comparison_operator"`

	// 告警阈值。单一阈值时value和alarm_level配对使用，当hierarchical_value和value同时使用时以hierarchical_value为准。取值范围[0, Number.MAX_VALUE]，Number.MAX_VALUE值为1.7976931348623157e+108。具体阈值取值请参见附录中各服务监控指标中取值范围，如支持监控的服务列表中ECS的CPU使用率cpu_util取值范围可配置80。 [具体阈值取值请参见附录中各服务监控指标中取值范围，如[支持监控的服务列表](ces_03_0059.xml)中ECS的CPU使用率cpu_util取值范围可配置80。](tag: dt,g42,dt_test,hk_g42,hk_sbc,hws,hws_hk,ocb,sbc,tm)
	Value *float64 `json:"value,omitempty"`

	HierarchicalValue *HierarchicalValue `json:"hierarchical_value,omitempty"`

	// 数据的单位。
	Unit *string `json:"unit,omitempty"`

	// 告警策略类型，默认为空
	Type *string `json:"type,omitempty"`

	// 告警连续触发次数，事件告警时参数值为1~180（包括1和180）；指标告警和站点告警时，次数采用枚举值，枚举值分别为：1、2、3、4、5、10、15、30、60、90、120、180
	Count int32 `json:"count"`

	// 告警抑制时间，单位为秒，对应页面上创建告警规则时告警策略最后一个字段，该字段主要为解决告警频繁的问题，0代表不抑制，满足条件即告警；300代表满足告警触发条件后每5分钟告警一次；
	SuppressDuration *int32 `json:"suppress_duration,omitempty"`

	// 告警级别, 1为紧急，2为重要，3为次要，4为提示。默认值为2。
	Level *int32 `json:"level,omitempty"`

	// 服务的命名空间，查询各服务命名空间请参考“[服务命名空间](ces_03_0059.xml)”
	Namespace *string `json:"namespace,omitempty"`

	// 指标维度名称。各服务资源的指标维度名称可查看：“[指标维度名称](ces_03_0059.xml)”。 约束与限制：    资源层级为子维度的告警规则，修改告警策略时，指标维度名必须与告警规则监控资源的维度保持一致。    资源层级为云产品的告警规则，修改告警策略时，指标维度名中首层维度必须与告警规则监控资源的维度保持一致。 举例：    ECS服务的资源层级为子维度、维度为云服务器的告警规则，需要设置维度名为：instance_id；    ECS服务的资源层级为云产品的告警规则时，维度名可以为：instance_id,disk
	DimensionName *string `json:"dimension_name,omitempty"`
}

func (o UpdatePolicy) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicy struct{}"
	}

	return strings.Join([]string{"UpdatePolicy", string(data)}, " ")
}
