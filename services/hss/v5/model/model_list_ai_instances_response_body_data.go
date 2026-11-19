package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAiInstancesResponseBodyData AI智能体实例响应数据项
type ListAiInstancesResponseBodyData struct {

	// **参数解释**： 智能体名字 **取值范围**： 字符长度为1-128位
	AgentName *string `json:"agent_name,omitempty"`

	// **参数解释**： 智能体ID **取值范围**： 字符长度为1-128位
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**： 智能体运行时名字 **取值范围**： 字符长度为1-128位
	RuntimeName *string `json:"runtime_name,omitempty"`

	// **参数解释**： 智能体运行时ID **取值范围**： 字符长度为1-128位
	RuntimeId *string `json:"runtime_id,omitempty"`

	// **参数解释**： SDK检测状态 **取值范围**： - open：开启 - close：关闭
	SdkStatus *string `json:"sdk_status,omitempty"`

	// **参数解释**： 智能体的来源 **取值范围**： - AgentArts - OfficeClaw - OpenClaw
	Source *string `json:"source,omitempty"`

	// **参数解释**： 智能体类型 **取值范围**： - agent：单智能体。 - chat：工作流。 - controller：多智能体。 - teamAgent：团队智能体。
	AgentType *string `json:"agent_type,omitempty"`

	// **参数解释**： 会话数量 **取值范围**： 0~2147483647
	SessionNumber *int32 `json:"session_number,omitempty"`

	// **参数解释**： 项目ID **取值范围**： 字符长度为1-128位
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**： 企业项目ID **取值范围**： 字符长度为1-128位
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 镜像url **取值范围**： 字符长度为1-512位
	ImageUrl *string `json:"image_url,omitempty"`

	// **参数解释**： 智能体实例的创建时间，毫秒时间戳 **取值范围**： 0~9223372036854775807
	InstanceCreateTime *int64 `json:"instance_create_time,omitempty"`

	// **参数解释**： 智能体实例的更新时间，毫秒时间戳 **取值范围**： 0~9223372036854775807
	InstanceUpdateTime *int64 `json:"instance_update_time,omitempty"`

	// **参数解释**： 行为轨迹告警数量 **取值范围**： 0~2147483647
	TraceAlarmNumber *int32 `json:"trace_alarm_number,omitempty"`

	// **参数解释**： 运行时告警数量 **取值范围**： 0~2147483647
	RuntimeAlarmNumber *int32 `json:"runtime_alarm_number,omitempty"`

	// **参数解释**： 策略组类型 **取值范围**： - 0：Agent运行时防护-OfficeAce - 1：Agent行为轨迹检测-OfficeAce - 2：Agent角色限定 - 3：Agent运行时防护-AgentArts - 4：Agent行为轨迹检测-AgentArts - 5：Agent运行时防护-OpenClaw - 6：Agent运行时防护-其他三方Claw
	PolicyGroupType *int32 `json:"policy_group_type,omitempty"`

	// **参数解释**: AI策略组列表 **取值范围**: 不涉及
	PolicyInfoList *[]AiAgentPolicy `json:"policy_info_list,omitempty"`

	// **参数解释**： 智能体运行环境 **取值范围**： - ECS - CCE - Others - AgentArts - OfficeClaw
	AgentRuntimeEnv *string `json:"agent_runtime_env,omitempty"`

	// **参数解释**： 集群ID，仅agent_runtime_env为CCE时返回 **取值范围**： 字符长度为1-128位
	ClusterId *string `json:"cluster_id,omitempty"`

	// **参数解释**： 集群类型 **取值范围**： - CCE：云容器引擎
	ClusterType *string `json:"cluster_type,omitempty"`

	// **参数解释**： 集群名称，仅agent_runtime_env为CCE时返回 **取值范围**： 字符长度为1-256位
	ClusterName *string `json:"cluster_name,omitempty"`

	// **参数解释**： 服务器名称 **取值范围**： 字符长度为1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**： 服务器公网IP地址 **取值范围**： 字符长度为1-128位
	PublicIp *string `json:"public_ip,omitempty"`

	// **参数解释**： 服务器私有IP地址 **取值范围**： 字符长度为1-128位
	PrivateIp *string `json:"private_ip,omitempty"`

	// **参数解释**： 服务器状态 **取值范围**： - ACTIVE：运行中 - SHUTOFF：已关机 - BUILD：创建中 - ERROR：异常 - DELETED：已删除 - REBOOT：重启中
	HostStatus *string `json:"host_status,omitempty"`

	// **参数解释**： 服务器ID **取值范围**： 字符长度为1-128位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**： 操作系统类型 **取值范围**： - Linux - Windows
	OsType *string `json:"os_type,omitempty"`

	// **参数解释**： 操作系统名称 **取值范围**： - EulerOS 2.9 - CentOS 7.9 - Ubuntu 20.04 - 其他操作系统名称
	OsName *string `json:"os_name,omitempty"`

	// **参数解释**： 操作系统版本号 **取值范围**： 字符长度为1-64位
	OsVersion *string `json:"os_version,omitempty"`

	// **参数解释**： 操作系统内核版本号 **取值范围**： 字符长度为1-128位
	KernelVersion *string `json:"kernel_version,omitempty"`

	// **参数解释**： Agent状态 **取值范围**： - online：在线 - offline：离线 - not_installed：未安装 - install_failed：安装失败 - installing：安装中
	AgentStatus *string `json:"agent_status,omitempty"`

	// **参数解释**： 智能体卫士插件状态 **取值范围**： - not_installed：未安装 - install_failed：安装失败 - online：在线 - offline：离线 - installing：安装中
	DefenderStatus *string `json:"defender_status,omitempty"`

	// **参数解释** 容器实例id **取值范围**: 字符长度0-128位
	ContainerId *string `json:"container_id,omitempty"`
}

func (o ListAiInstancesResponseBodyData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAiInstancesResponseBodyData struct{}"
	}

	return strings.Join([]string{"ListAiInstancesResponseBodyData", string(data)}, " ")
}
