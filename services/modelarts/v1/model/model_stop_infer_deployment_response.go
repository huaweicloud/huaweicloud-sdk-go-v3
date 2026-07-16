package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// StopInferDeploymentResponse Response Object
type StopInferDeploymentResponse struct {

	// **参数解释：** 部署ID。 **取值范围：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 部署名称。 **取值范围：** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释：** 资源池ID，查询指定资源池下的服务，默认不过滤。可通过[查询资源池列表](ShowPool.xml)获取。 **取值范围：** 不涉及。
	PoolId *string `json:"pool_id,omitempty"`

	// **参数解释：** 部署场景下，服务实例数量。 **约束限制：** 不涉及。 **取值范围：** [1, 128]。 **默认取值：** 不涉及。
	Count *int32 `json:"count,omitempty"`

	// **参数解释：** 系统日志转储开关。 **约束限制：** 只有NPU资源池有，且逻辑池是没有的。 **取值范围：** 不涉及。 **默认取值：** false。
	SystemLogDumpEnable *bool `json:"system_log_dump_enable,omitempty"`

	// **参数解释：** 推理单元配置。
	UnitConfigs *[]UnitConfigResponse `json:"unit_configs,omitempty"`

	// **参数解释：** 权重百分比，分配到此模型的流量权重，仅当模型部署为在线服务时需要配置。 **取值范围：** [0, 100]。
	Weight *int32 `json:"weight,omitempty"`

	// **参数解释：** 流量比例，单个部署实例预期接收用户的流量与总流量比值，是由流量权重配置和部署状态计算所得的值。 **取值范围：** 0.00%~100.00%。
	TrafficRatio *string `json:"traffic_ratio,omitempty"`

	// **参数解释：** 凭证类型相关配置，用户可以在此处选择使用的凭证类型（dew，agency） **约束限制：** 不涉及。 **取值范围：** - [dew：DEW密钥。](tag:hws,hws_hk,fcs) - agency：临时委托凭证。 **默认取值：** 不涉及。
	SecretType *string `json:"secret_type,omitempty"`

	// **参数解释：** 凭证名称，用户使用dew类型凭证时可以在此处选择使用的凭证。 **约束限制：** 不涉及。
	SecretName *string `json:"secret_name,omitempty"`

	// **参数解释：** 服务优先级。 **约束限制：** 不涉及。 **取值范围：** [1, 3]。 **默认取值：** 不涉及。
	Priority *int32 `json:"priority,omitempty"`

	// **参数解释：** 高可用开关，开启后不同实例的pod将尽量均匀分布到不同的节点上。（准备下线，请使用schedule_strategy字段） **取值范围：** - true: 高可用开启 - false: 高可用关闭。
	HighAvailSwitch *bool `json:"high_avail_switch,omitempty"`

	// **参数解释：** 算法框架。 **取值范围：** - COMMON： 普通在线服务。
	Framework *string `json:"framework,omitempty"`

	// **参数解释：** 服务版本号，以数字及点号组成，形如1.0.1。 **约束限制：** 不涉及。 **取值范围：** 1.0.0 ~ 99.99.99，长度不超过8位。 **默认取值：** 不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释：** 版本id，可通过查询version列表查询。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	VersionId *string `json:"version_id,omitempty"`

	// **参数解释：** 服务当前状态。 **取值范围：** - DEPLOYING：部署中。 - FAILED：失败。 - STOPPED：停止。 - RUNNING：运行中。 - DELETING：删除中。 - STOPPING：停止中。 - CONCERNING：告警。 - DELETED: 已删除。 - RESTARTING: 重启中。 - UPGRADING：升级中。 - ERROR：异常。 - INTERRUPTING：中断中。
	Status *string `json:"status,omitempty"`

	// **参数解释：** 部署场景下，服务运行实例数量。 **约束限制：** 不涉及。 **取值范围：** [1, 128]。 **默认取值：** 不涉及。
	RunningCount *int32 `json:"running_count,omitempty"`

	// **参数解释：** 部署类型。 **取值范围：** - SINGLE：单机单卡。 - MULTI：多机多卡。
	DeployType *string `json:"deploy_type,omitempty"`

	// **参数解释：** 是否开启镜像流量。 **取值范围：** 不涉及。
	MirrorTrafficEnable *bool `json:"mirror_traffic_enable,omitempty"`

	// **参数解释：** 镜像流量权重。 **取值范围：** 50。
	MirrorTrafficWeight *string `json:"mirror_traffic_weight,omitempty"`

	// **参数解释：** 服务版本数量。 **取值范围：** 不涉及。
	VersionCount *int32 `json:"version_count,omitempty"`

	// **参数解释：** 工作负载类型。 **取值范围：** 不涉及。
	WorkloadType *string `json:"workload_type,omitempty"`

	// **参数解释：** 更新时间，根据更新时的当前时间自动生成。 **取值范围：** 毫秒级时间戳，13位数字。
	UpdateAt *int64 `json:"update_at,omitempty"`

	// **参数解释：** 部署对接lts状态。 **取值范围：** - ON：开启。 - OFF：关闭。
	LtsState *string `json:"lts_state,omitempty"`

	// **参数解释：** 部署ID。
	InferName *string `json:"infer_name,omitempty"`

	Model *InferModelResponse `json:"model,omitempty"`

	AdvancedConfig *AdvancedConfig `json:"advanced_config,omitempty"`

	// **参数解释：** 部署描述。
	Description *string `json:"description,omitempty"`

	// 参数解释： 创建时间，根据创建时的当前时间自动生成。 取值范围： 毫秒级时间戳，13位数字，如1609459200000。
	CreateAt *string `json:"create_at,omitempty"`

	// **参数解释：** 调度策略。 **取值范围：** - HIGH_AVAILABILITY：高可用调度 - HIGH_UTILIZATION：紧凑调度 - HIGH_PERFORMANCE：高性能调度
	ScheduleStrategy *string `json:"schedule_strategy,omitempty"`
	HttpStatusCode   int     `json:"-"`
}

func (o StopInferDeploymentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "StopInferDeploymentResponse struct{}"
	}

	return strings.Join([]string{"StopInferDeploymentResponse", string(data)}, " ")
}
