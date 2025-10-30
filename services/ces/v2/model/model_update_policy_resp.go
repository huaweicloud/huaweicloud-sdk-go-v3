package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type UpdatePolicyResp struct {

	// **参数解释**： 资源的监控指标名称，各服务的指标名称可查看：“[服务指标名称](ces_03_0059.xml)”。 **取值范围**： 必须以字母开头，只能包含0-9/a-z/A-Z/_/-。字符长度最短为1，最大为96。如：弹性云服务器中的监控指标cpu_util，表示弹性服务器的CPU使用率；文档数据库中的指标mongo001_command_ps，表示command执行频率。
	MetricName *string `json:"metric_name,omitempty"`

	ExtraInfo *MetricExtraInfoResp `json:"extra_info,omitempty"`

	// **参数解释**： 指标周期，单位是秒。如想了解各个云服务的指标原始周期可以参考“[支持服务列表](ces_03_0059.xml)”。 **取值范围**： 枚举值。 - 0代表立即触发，仅限事件类告警使用。 - 1代表指标的原始周期，比如RDS监控指标原始周期是60s，表示该RDS指标按60s周期为一个数据点参与告警计算。 - 300代表指标按5分钟聚合周期为一个数据点参与告警计算。 - 1200代表指标按20分钟聚合周期为一个数据点参与告警计算。 - 3600代表指标按1小时聚合周期为一个数据点参与告警计算。 - 14400代表指标按4小时聚合周期为一个数据点参与告警计算。 - 86400代表指标按1天聚合周期为一个数据点参与告警计算。
	Period *int32 `json:"period,omitempty"`

	// **参数解释**： 聚合方式。         **取值范围**： average： 平均值，variance：方差，min：最小值，max：最大值，sum：求和。
	Filter *string `json:"filter,omitempty"`

	// **参数解释**： 阈值符号。     **取值范围**： 支持的值为(>|<|>=|<=|=|!=|cycle_decrease|cycle_increase|cycle_wave);cycle_decrease为环比下降,cycle_increase为环比上升,cycle_wave为环比波动。
	ComparisonOperator *string `json:"comparison_operator,omitempty"`

	// **参数解释**： 告警阈值。具体阈值取值请参见附录中各服务监控指标中取值范围。 **取值范围**： 最小值为0，最大值为1.7976931348623157e+108。
	Value *float64 `json:"value,omitempty"`

	HierarchicalValue *HierarchicalValueResp `json:"hierarchical_value,omitempty"`

	// **参数解释**： 数据的单位。    **取值范围**： 长度为[0,32]个字符。
	Unit *string `json:"unit,omitempty"`

	// **参数解释**： 告警策略类型，已废弃，不推荐使用。 **取值范围**： 只能为auto。
	Type *string `json:"type,omitempty"`

	// **参数解释**： 告警连续触发次数。     **取值范围**： 事件告警时参数值为1~180（包括1和180）；指标告警和站点告警时，次数采用枚举值，枚举值分别为：1、2、3、4、5、10、15、30、60、90、120、180。
	Count *int32 `json:"count,omitempty"`

	// **参数解释**： 告警抑制时间，单位为秒，对应页面上创建告警规则时告警策略最后一个字段，该字段主要为解决告警频繁的问题。 **取值范围**： 枚举值，只能为0、300、600、900、1800、3600、10800、21600、43200、86400。 - 0：在立即触发场景中，0代表不抑制；在累计触发场景，0代表只告警一次。 - 300代表满足告警触发条件后每5分钟告警一次。 - 600代表满足告警触发条件后每10分钟告警一次。 - 900代表满足告警触发条件后每15分钟告警一次。 - 1800代表满足告警触发条件后每30分钟告警一次。 - 3600代表满足告警触发条件后每60分钟告警一次。 - 10800代表满足告警触发条件后每3小时告警一次。 - 21600代表满足告警触发条件后每6小时告警一次。 - 43200代表满足告警触发条件后每12小时告警一次。 - 86400代表满足告警触发条件后每一天告警一次。
	SuppressDuration *int32 `json:"suppress_duration,omitempty"`

	// **参数解释**： 告警级别。    **取值范围**： 只能为1、2、3、4。1为紧急，2为重要，3为次要，4为提示。
	Level *int32 `json:"level,omitempty"`

	// **参数解释**： 各服务命名空间，请参考“[服务命名空间](ces_03_0059.xml)”。产品层级规则增加namespace（服务命名空间）和dimension_name（服务维度名称），用于指明策略生效的资源归属。 **取值范围**： 格式为service.item；service和item必须是字符串，必须以字母开头，只能包含0-9/a-z/A-Z/_。长度为[0,32]个字符。        。
	Namespace *string `json:"namespace,omitempty"`

	// **参数解释**： 指标维度名称，各服务资源的指标维度名称可查看：“[服务维度名称](ces_03_0059.xml)”。 **取值范围**： 目前最大支持4个维度。举例：单维度场景：instance_id；多维度场景：instance_id,disk
	DimensionName *string `json:"dimension_name,omitempty"`
}

func (o UpdatePolicyResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdatePolicyResp struct{}"
	}

	return strings.Join([]string{"UpdatePolicyResp", string(data)}, " ")
}
