package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// CancelInferDeploymentResponse Response Object
type CancelInferDeploymentResponse struct {

	// 参数解释： 部署ID，在[添加部署](CreateInferDeployment.xml)时即可在返回体中获取，也可通过[查询服务部署列表](ListInferDeployments.xml)获取当前用户拥有的部署，其中deployment_id字段即为部署ID。 约束限制： 不涉及。 取值范围： 部署ID。 默认取值： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 服务部署名字
	Name *string `json:"name,omitempty"`

	// 参数解释： 服务部署名字。
	DeploymentName *string `json:"deployment_name,omitempty"`

	// 参数解释： 部署id（废弃字段）。 取值范围： 不涉及。
	InferName *string `json:"infer_name,omitempty"`

	// 参数解释： 描述 取值范围： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 专属资源池ID。 **取值范围：** 不涉及。
	PoolId *string `json:"pool_id,omitempty"`

	// **参数解释：** 算法框架。 **取值范围：** - COMMON： 普通在线服务。
	Framework *string `json:"framework,omitempty"`

	// **参数解释：** 当前部署版本信息。
	Version *string `json:"version,omitempty"`

	// **参数解释：** 版本id，可通过查询version列表查询。 **取值范围：** 不涉及。
	VersionId *string `json:"version_id,omitempty"`

	// **参数解释：** 服务版本数量。 **取值范围：** 不涉及。
	VersionCount *int32 `json:"version_count,omitempty"`

	// **参数解释：** 服务部署实例数。
	Count *int32 `json:"count,omitempty"`

	// **参数解释：** 权重百分比，分配到此模型的流量权重，仅当模型部署为在线服务时需要配置。 **约束限制：** 不涉及。 **取值范围：** [0, 100]。 **默认取值：** 不涉及。
	Weight *int32 `json:"weight,omitempty"`

	// **参数解释：** 服务优先级。 **约束限制：** 不涉及。 **取值范围：** [1, 3]。 **默认取值：** 不涉及。
	Priority *int32 `json:"priority,omitempty"`

	// **参数解释：** 凭证类型相关配置，用户可以在此处选择使用的凭证类型（dew，agency） **约束限制：** 不涉及。 **取值范围：** - [dew：DEW密钥。](tag:hws,hws_hk,fcs) - agency：临时委托凭证。 **默认取值：** 不涉及。
	SecretType *string `json:"secret_type,omitempty"`

	// **参数解释：** 凭证名称，用户使用dew类型凭证时可以在此处选择使用的凭证。 **取值范围：** 1-64个字符，只能包含字母、数字、连字符和下划线。
	SecretName *string `json:"secret_name,omitempty"`

	// **参数解释：** 服务部署状态
	Status *string `json:"status,omitempty"`

	// **参数解释：** 部署类型。 **取值范围：** - SINGLE：单机单卡。 - MULTI：多机多卡。
	DeployType *string `json:"deploy_type,omitempty"`

	// **参数解释：** 工作负载类型。 **取值范围：** - LWS - DEPLOYMENT **默认取值：** DEPLOYMENT
	WorkloadType *string `json:"workload_type,omitempty"`

	// **参数解释：** 高可用开关，开启后不同实例的pod将尽量均匀分布到不同的节点上。 **约束限制：** 不涉及 **取值范围：** - true: 高可用开启。 - false: 高可用关闭。
	HighAvailSwitch *bool `json:"high_avail_switch,omitempty"`

	// **参数解释：** 系统日志转储开关。 **约束限制：** 只有NPU资源池有，且逻辑池是没有的 **取值范围：** 不涉及 **默认取值：** false
	SystemLogDumpEnable *bool `json:"system_log_dump_enable,omitempty"`

	// **参数解释：** 调度策略。 **取值范围：** 不涉及。
	ScheduleStrategy *string `json:"schedule_strategy,omitempty"`

	// **参数解释：** 实例单元配置。 **约束限制：** 单机推理时，个数只会为1；如果是分布式推理时，根据不同框架，实例单元配置可灵活配置。
	UnitConfigs *[]UnitConfig `json:"unit_configs,omitempty"`

	Model *GroupModel `json:"model,omitempty"`

	AdvancedConfig *AdvancedConfig `json:"advanced_config,omitempty"`

	// **参数解释：** 是否开启镜像流量。 **取值范围：** 不涉及。
	MirrorTrafficEnable *bool `json:"mirror_traffic_enable,omitempty"`

	// **参数解释：** 镜像流量权重。 **取值范围：** 不涉及。
	MirrorTrafficWeight *int32 `json:"mirror_traffic_weight,omitempty"`

	// **参数解释：** 用户名。 **取值范围：** 不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释：** 委托名称。 **取值范围：** 不涉及。
	AgencyName *string `json:"agency_name,omitempty"`

	// **参数解释：** 最后使用时间。 **取值范围：** 不涉及。
	LastUsedTime *sdktime.SdkTime `json:"last_used_time,omitempty"`

	DeploymentTaskLimit *DeploymentTaskLimit `json:"deployment_task_limit,omitempty"`

	// **参数解释：** 当服务或者部署被冻结时返回的冻结类型信息。
	FrozenInfos *[]FrozenInfo `json:"frozen_infos,omitempty"`

	// **参数解释：** 创建时间，根据创建时的当前时间自动生成。 **取值范围：** 毫秒级时间戳，13位数字
	CreateAt *sdktime.SdkTime `json:"create_at,omitempty"`

	// **参数解释：** 更新时间，根据更新时的当前时间自动生成。 **取值范围：** 毫秒级时间戳，13位数字。
	UpdateAt       *sdktime.SdkTime `json:"update_at,omitempty"`
	HttpStatusCode int              `json:"-"`
}

func (o CancelInferDeploymentResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CancelInferDeploymentResponse struct{}"
	}

	return strings.Join([]string{"CancelInferDeploymentResponse", string(data)}, " ")
}
