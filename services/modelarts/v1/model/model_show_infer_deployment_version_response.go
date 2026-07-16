package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ShowInferDeploymentVersionResponse Response Object
type ShowInferDeploymentVersionResponse struct {

	// 参数解释： 部署ID，在[添加部署](CreateInferDeployment.xml)时即可在返回体中获取，也可通过[查询服务部署列表](ListInferDeployments.xml)获取当前用户拥有的部署，其中deployment_id字段即为部署ID。 约束限制： 不涉及。 取值范围： 部署ID。 默认取值： 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释：** 服务部署名字
	Name *string `json:"name,omitempty"`

	// 参数解释： 部署id（废弃字段）。 取值范围： 不涉及。
	InferName *string `json:"infer_name,omitempty"`

	// **参数解释：** 创建时间，根据创建时的当前时间自动生成。 **取值范围：** 毫秒级时间戳，13位数字
	CreateAt *sdktime.SdkTime `json:"create_at,omitempty"`

	// 参数解释： 描述 取值范围： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释：** 专属资源池ID。 **取值范围：** 不涉及。
	PoolId *string `json:"pool_id,omitempty"`

	// **参数解释：** 算法框架。 **取值范围：** - COMMON： 普通在线服务。
	Framework *string `json:"framework,omitempty"`

	// **参数解释：** 服务优先级。 **约束限制：** 不涉及。 **取值范围：** [1, 3]。 **默认取值：** 不涉及。
	Priority *string `json:"priority,omitempty"`

	// **参数解释：** 凭证类型相关配置，用户可以在此处选择使用的凭证类型（dew，agency） **约束限制：** 不涉及。 **取值范围：** - [dew：DEW密钥。](tag:hws,hws_hk,fcs) - agency：临时委托凭证。 **默认取值：** 不涉及。
	SecretType *string `json:"secret_type,omitempty"`

	// **参数解释：** 服务部署状态
	Status *string `json:"status,omitempty"`

	// **参数解释：** 服务实例数
	Count *int32 `json:"count,omitempty"`

	// **参数解释：** 高可用开关，开启后不同实例的pod将尽量均匀分布到不同的节点上。 **约束限制：** 不涉及 **取值范围：** 不涉及 **默认取值：** true
	HighAvailSwitch *string `json:"high_avail_switch,omitempty"`

	// **参数解释：** 系统日志转储开关。 **约束限制：** 只有NPU资源池有，且逻辑池是没有的 **取值范围：** 不涉及 **默认取值：** false
	SystemLogDumpEnable *string `json:"system_log_dump_enable,omitempty"`

	// **参数解释：** 实例单元配置。 **约束限制：** 单机推理时，个数只会为1；如果是分布式推理时，根据不同框架，实例单元配置可灵活配置。
	UnitConfigs *[]UnitConfig `json:"unit_configs,omitempty"`

	// **参数解释：** 更新时间，根据更新时的当前时间自动生成。 **取值范围：** 毫秒级时间戳，13位数字。
	UpdateAt *sdktime.SdkTime `json:"update_at,omitempty"`

	// **参数解释：** 当前服务版本信息。
	Version *string `json:"version,omitempty"`

	// **参数解释：** 服务版本数量。 **取值范围：** 不涉及。
	VersionCount *int32 `json:"version_count,omitempty"`

	// **参数解释：** 权重百分比，分配到此模型的流量权重，仅当模型部署为在线服务时需要配置。 **约束限制：** 不涉及。 **取值范围：** [0, 100]。 **默认取值：** 不涉及。
	Weight *int32 `json:"weight,omitempty"`

	AdvancedConfig *AdvancedConfig `json:"advanced_config,omitempty"`

	// **参数解释：** 巫山工作流ID。 **取值范围：** 不涉及。
	JobId *string `json:"job_id,omitempty"`

	// 参数解释： 服务部署名字。
	DeploymentName *string `json:"deployment_name,omitempty"`

	// **参数解释：** 当服务或者部署被冻结时返回的冻结类型信息。
	FrozenInfos    *[]FrozenInfo `json:"frozen_infos,omitempty"`
	HttpStatusCode int           `json:"-"`
}

func (o ShowInferDeploymentVersionResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInferDeploymentVersionResponse struct{}"
	}

	return strings.Join([]string{"ShowInferDeploymentVersionResponse", string(data)}, " ")
}
