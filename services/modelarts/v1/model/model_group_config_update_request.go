package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// GroupConfigUpdateRequest **参数解释：** 服务实例组配置，当推理方式为BATCH/EDGE时仅支持配置一个模型；当推理方式为REAL_TIME时，可根据业务需要配置多个服务实例并分配权重。 **约束限制：** 不涉及
type GroupConfigUpdateRequest struct {

	// **参数解释：** 部署ID。 **约束限制：** 不填保留原有值。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 部署名称。 **约束限制：** 必填参数，不填不保留原有值。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Name string `json:"name"`

	// **参数解释：** 资源池ID，查询指定资源池下的服务，默认不过滤。可通过[查询资源池列表](ShowPool.xml)获取。 **约束限制：** 不填保留原有值。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	PoolId *string `json:"pool_id,omitempty"`

	// **参数解释：** 框架类型。 **约束限制：** 不填则为默认值。 **取值范围：** - COMMON：普通在线服务 - VLLM：VLLM框架 - MINDIE：MINDIE框架 **默认取值：** COMMON
	Framework *string `json:"framework,omitempty"`

	// **参数解释：** 部署场景下，服务实例数量。 **约束限制：** 不填则为默认值。 **取值范围：** [1, 128]。 **默认取值：** 1
	Count *int32 `json:"count,omitempty"`

	// **参数解释：** 部署类型。 **约束限制：** 不填保留原有值。 **取值范围：** - SINGLE：常规部署 - MULTI：分布式部署 **默认取值：** 不涉及
	DeployType *string `json:"deploy_type,omitempty"`

	// **参数解释：** 系统日志转储开关。 **约束限制：** 不填则为默认值。 **取值范围：** 不涉及 **默认取值：** false
	SystemLogDumpEnable *bool `json:"system_log_dump_enable,omitempty"`

	// **参数解释：** 实例单元配置。 **约束限制：** - 单机推理时，个数只会为1；如果是分布式推理时，根据不同框架，实例单元配置可灵活配置。 - 必填字段。 **取值范围：** 不涉及 **默认取值：** 不涉及
	UnitConfigs *[]UnitConfig `json:"unit_configs,omitempty"`

	// **参数解释：** 权重百分比，分配到此模型的流量权重，仅当模型部署为在线服务时需要配置。 **约束限制：** 不填保留原有值。 **取值范围：** [0, 100]。 **默认取值：** 不涉及
	Weight int32 `json:"weight"`

	// **参数解释：** 凭证类型相关配置，用户可以在此处选择使用的凭证类型（dew，agency） **约束限制：** 1.使用临时委托凭证类型约束限制:集群已安装CCE容器存储（Everest）v2.4.204及以上版本，且集群版本为v1.28及以上且确保局点已启用IAM5服务。 2.若插件版本不足或集群不支持临时委托凭证，则需通过DEW服务挂载。 3.不填保留原有值。 **取值范围：** - [dew：DEW密钥。](tag:hws,hws_hk,fcs) - agency：临时委托凭证。 **默认取值：** 不涉及。
	SecretType *string `json:"secret_type,omitempty"`

	// **参数解释**： 认证凭证名称，用户使用dew类型凭证时可以在此处选择使用的凭证。 **约束限制**： secret_type是dew时必填，不填保留原有值。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	SecretName *string `json:"secret_name,omitempty"`

	// **参数解释：** 服务优先级。 **约束限制：** - 如服务处于“运行中”，priority字段为必要参数，且value必须为原版的值； - 如服务处于“停止”，priority字段为非必要参数。 - 不填保留原有值。 **取值范围：** [1, 3]。 **默认取值：** 不涉及。
	Priority *int32 `json:"priority,omitempty"`

	// **参数解释：** 高可用开关，开启后不同实例的pod将尽量均匀分布到不同的节点上。 **约束限制：** 不填则为默认值 **取值范围：** 不涉及 **默认取值：** true
	HighAvailSwitch *bool `json:"high_avail_switch,omitempty"`

	// **参数解释：** 部署备注。 **约束限制：** 不填则将部署描述清空。 **取值范围：** 长度不可以超过512，不能包含大于号，小于号。 **默认取值：** 默认为空。
	Description *string `json:"description,omitempty"`

	AdvancedConfig *AdvancedConfig `json:"advanced_config"`

	Model *GroupModel `json:"model,omitempty"`

	// **参数解释：** 镜像流量开关。 **约束限制：** 不填保留原有值 **取值范围：** 不涉及 **默认取值：** 不涉及
	MirrorTrafficEnable *bool `json:"mirror_traffic_enable,omitempty"`

	// **参数解释：** 镜像流量权重。 **约束限制：** 不填保留原有值。 **取值范围：** 不涉及 **默认取值：** 不涉及
	MirrorTrafficWeight *int32 `json:"mirror_traffic_weight,omitempty"`

	// **参数解释：** 部署状态。 **约束限制：** 不填保留原有值。 **取值范围：** - DEPLOYING：部署中 - FAILED：失败 - STOPPED：停止 - RUNNING：运行中 - DELETING：删除中 - STOPPING：停止中 - CONCERNING：存在潜在问题 - DELETED：删除 - RESTARTING：重启中 - UPGRADING：更新中 - ERROR：错误 - INTERRUPTING：中断中 **默认取值：** 不涉及
	Status *string `json:"status,omitempty"`

	DeploymentTaskLimit *DeploymentTaskLimit `json:"deployment_task_limit,omitempty"`

	// **参数解释：** 调度策略。 **约束限制：** 不涉及。 **取值范围：** - HIGH_AVAILABILITY：高可用调度 - HIGH_UTILIZATION：紧凑调度 - HIGH_PERFORMANCE：高性能调度 **默认取值：** HIGH_AVAILABILITY。
	ScheduleStrategy *string `json:"schedule_strategy,omitempty"`
}

func (o GroupConfigUpdateRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupConfigUpdateRequest struct{}"
	}

	return strings.Join([]string{"GroupConfigUpdateRequest", string(data)}, " ")
}
