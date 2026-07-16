package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// GroupConfig **参数解释：** 服务实例组配置，当推理方式为BATCH/EDGE时仅支持配置一个模型；当推理方式为REAL_TIME时，可根据业务需要配置多个服务实例并分配权重。 **约束限制：** 不涉及
type GroupConfig struct {

	// **参数解释：** 部署ID。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 部署名称。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	Name string `json:"name"`

	// **参数解释：** 资源池ID，查询指定资源池下的服务，默认不过滤。可通过[查询资源池列表](ShowPool.xml)获取。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	PoolId *string `json:"pool_id,omitempty"`

	// **参数解释：** 部署场景下，服务实例数量。 **约束限制：** 不涉及。 **取值范围：** [1, 128]。 **默认取值：** 不涉及。
	Count *int32 `json:"count,omitempty"`

	// **参数解释：** 系统日志转储开关。 **约束限制：** 只有NPU资源池有，且逻辑池是没有的。 **取值范围：** 不涉及。 **默认取值：** false。
	SystemLogDumpEnable *bool `json:"system_log_dump_enable,omitempty"`

	// **参数解释：** 推理单元配置。 **约束限制：** 单机推理时，个数只会为1；如果是分布式推理时，根据不同框架，实例单元配置可灵活配置。
	UnitConfigs *[]UnitConfig `json:"unit_configs,omitempty"`

	// **参数解释：** 权重百分比，分配到此模型的流量权重，仅当模型部署为在线服务时需要配置。 **约束限制：** 不涉及。 **取值范围：** [0, 100]。 **默认取值：** 不涉及。
	Weight int32 `json:"weight"`

	// **参数解释：** 凭证类型相关配置，用户可以在此处选择使用的凭证类型（dew，agency） **约束限制：** 1.使用临时委托凭证类型约束限制:集群已安装CCE容器存储（Everest）v2.4.204及以上版本，且集群版本为v1.28及以上且确保局点已启用IAM5服务。 2.若插件版本不足或集群不支持临时委托凭证，则需通过DEW服务挂载。 **取值范围：** - [dew：DEW密钥。](tag:hws,hws_hk,fcs) - agency：临时委托凭证。 **默认取值：** 不涉及。
	SecretType *string `json:"secret_type,omitempty"`

	// **参数解释**： 凭证名称，用户使用dew类型凭证时可以在此处选择使用的凭证。 **约束限制**： 不涉及。
	SecretName *string `json:"secret_name,omitempty"`

	// **参数解释：** 服务优先级。 **约束限制：** - 如服务处于\"运行中\"，priority字段为必要参数，且value必须与原服务的priority值相同； - 如服务处于\"停止\"，priority字段为非必要参数。 **取值范围：** [1, 3]。 **默认取值：** 不涉及。
	Priority *int32 `json:"priority,omitempty"`

	// **参数解释：** 高可用开关，开启后不同实例的pod将尽量均匀分布到不同的节点上。（准备下线，请使用schedule_strategy字段） **约束限制：** 不涉及。 **取值范围：** - true：高可用开启。 - false：高可用关闭。 **默认取值：** true。
	HighAvailSwitch *bool `json:"high_avail_switch,omitempty"`

	// **参数解释：** 调度策略。 **约束限制：** 不涉及。 **取值范围：** - HIGH_AVAILABILITY：高可用调度 - HIGH_UTILIZATION：紧凑调度 - HIGH_PERFORMANCE：高性能调度 **默认取值：** HIGH_AVAILABILITY。
	ScheduleStrategy *string `json:"schedule_strategy,omitempty"`

	// **参数解释：** 服务版本号，以数字及点号组成，形如1.0.1。 **约束限制：** 不涉及。 **取值范围：** 1.0.0 ~ 99.99.99，长度不超过8位。 **默认取值：** 不涉及。
	Version *string `json:"version,omitempty"`

	// **参数解释：** 版本id，可通过查询version列表查询。 **约束限制：** 不涉及。 **取值范围：** 不涉及。 **默认取值：** 不涉及。
	VersionId *string `json:"version_id,omitempty"`

	// **参数解释：** 服务备注。 **约束限制：** 不涉及。 **取值范围：** 长度不可以超过512，不能包含大于号，小于号。 **默认取值：** 默认为空。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 服务框架。 **约束限制：** 仅支持以下枚举值：COMMON | VLLM | MINDIE。 **取值范围：** 仅支持以下枚举值：COMMON | VLLM | MINDIE。 **默认取值：** COMMON。
	Framework *GroupConfigFramework `json:"framework,omitempty"`

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

	// **参数解释：** 工作负载类型。 **取值范围：** - DEPLOYMENT：DEPLOYMENT类型 - LWS：LWS类型
	WorkloadType *string `json:"workload_type,omitempty"`

	// **参数解释：** 更新时间，根据更新时的当前时间自动生成。 **取值范围：** 毫秒级时间戳，13位数字。
	UpdateAt *int64 `json:"update_at,omitempty"`

	Model *ModelResource `json:"model,omitempty"`

	AdvancedConfig *AdvancedConfig `json:"advanced_config,omitempty"`
}

func (o GroupConfig) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "GroupConfig struct{}"
	}

	return strings.Join([]string{"GroupConfig", string(data)}, " ")
}

type GroupConfigFramework struct {
	value string
}

type GroupConfigFrameworkEnum struct {
	COMMON GroupConfigFramework
	VLLM   GroupConfigFramework
	MINDIE GroupConfigFramework
}

func GetGroupConfigFrameworkEnum() GroupConfigFrameworkEnum {
	return GroupConfigFrameworkEnum{
		COMMON: GroupConfigFramework{
			value: "COMMON",
		},
		VLLM: GroupConfigFramework{
			value: "VLLM",
		},
		MINDIE: GroupConfigFramework{
			value: "MINDIE",
		},
	}
}

func (c GroupConfigFramework) Value() string {
	return c.value
}

func (c GroupConfigFramework) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *GroupConfigFramework) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
