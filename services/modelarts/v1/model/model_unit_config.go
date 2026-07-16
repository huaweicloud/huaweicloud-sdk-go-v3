package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// UnitConfig **参数解释：** 推理单元配置。 **约束限制：** 单机推理时，个数只会为1；如果是分布式推理时，根据不同框架，实例单元配置可灵活配置。
type UnitConfig struct {

	// **参数解释：** 实例单元ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 实例单元角色名称。 **约束限制：** 最大长度为16字符，且需满足正则：^\\[a-z0-9]([-a-z0-9]*[a-z0-9])?$ **默认取值：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 实例单元角色。 **约束限制：** 不涉及。 **取值范围：** - COMMON：表示其他角色。 **默认取值：** 不涉及。
	Role *string `json:"role,omitempty"`

	CustomSpec *CustomResourceSpec `json:"custom_spec,omitempty"`

	// **参数解释：** 资源规格，根据所提供版本选择适合业务的规格。当specification为custom为自定义规格。由custom_spec指定部署的规格配置。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Flavor *string `json:"flavor,omitempty"`

	// **参数解释：** 资源规格的显示名称。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	FlavorDisplayName *string `json:"flavor_display_name,omitempty"`

	Image *ImageInfo `json:"image"`

	// **参数解释：** 废弃参数，推荐使用files进行模型相关配置。 模型相关配置，用户可以在此处选择模型及权重文件配合镜像使用。 **约束限制：** 不涉及。
	Models *[]ModelResource `json:"models,omitempty"`

	// **参数解释：** 废弃参数，推荐使用files进行代码相关配置。 代码相关配置，用户可以在此处选择代码所在的obs路径等。 **约束限制：** 不涉及。
	Codes *[]Code `json:"codes,omitempty"`

	// **参数解释：** 模型和代码相关配置，用户可以在此处选择模型及权重文件配合镜像使用以及代码所在的obs路径等。 **约束限制：** 不涉及。
	Files *[]FileInfo `json:"files,omitempty"`

	// **参数解释：** 用户转储配置，用户可以在此处选择要转储的目的obs。 **约束限制：** 最多配置20组。
	Dumps *[]Dump `json:"dumps,omitempty"`

	// **参数解释：** 配置实例个数。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Count *int32 `json:"count,omitempty"`

	// **参数解释：** 启动命令。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Cmd *string `json:"cmd,omitempty"`

	TerminationGrace *TerminationGrace `json:"termination_grace,omitempty"`

	// **参数解释：** 环境变量。 **约束限制：** 变量键长度不大于64，由字母、数字、下划线、中划线、点组成，不能以数字开头。值的输入内容不能存在HTML标签，包括<^>。
	Envs map[string]string `json:"envs,omitempty"`

	ReadinessHealth *Health `json:"readiness_health,omitempty"`

	StartupHealth *Health `json:"startup_health,omitempty"`

	LivenessHealth *Health `json:"liveness_health,omitempty"`

	// **参数解释：** 端口。 **约束限制：** 不涉及。 **取值范围：** [1,65535]。 **默认取值：** 不涉及。
	Port *int32 `json:"port,omitempty"`

	// **参数解释：** 自动重建策略，开启后，由于部署配置变更或者故障等原因导致Pod重启时，平台将按策略自动执行重建。若不开启，平台将不会主动干预处理。 **约束限制：** 不涉及。 **取值范围：** - Instance：部署副本重建，故障时重新拉起整个部署。 - Role：单元重建，当部署单元内的Pod出现故障时，重启该单元内的所有Pod。 - Pod：Pod重建，故障时重新拉起故障pod。 **默认取值：** 不涉及。
	Recovery *string `json:"recovery,omitempty"`

	// **参数解释：** 是否开启恢复策略。 **约束限制：** 不涉及。 **取值范围：** - true：开启恢复策略。 - false：不开启恢复策略。 **默认取值：** 不涉及。
	NpuResetEnable *bool `json:"npu_reset_enable,omitempty"`

	// **参数解释：** 单元副本数，当部署类型deploy_type为SINGLE或工作负载类型workload_type为DEPLOYMENT时，该参数无效。 **约束限制：** 不涉及。 **取值范围：** [1, 100] 或者为空。 **默认取值：** 默认值为1。
	GroupCount *int32 `json:"group_count,omitempty"`

	Affinity *Affinity `json:"affinity,omitempty"`

	SecurityConfig *ServiceSecurityConfig `json:"security_config,omitempty"`

	// **参数解释：** 节点池资源规格。 **约束限制：** 只能包含字母、数字、点、中划线和下划线。 **取值范围：** 长度不超过128字符。 **默认取值：** 不涉及。
	PoolResourceFlavor *string `json:"pool_resource_flavor,omitempty"`
}

func (o UnitConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UnitConfig struct{}"
	}

	return strings.Join([]string{"UnitConfig", string(data)}, " ")
}
