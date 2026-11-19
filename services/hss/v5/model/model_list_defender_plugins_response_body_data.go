package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListDefenderPluginsResponseBodyData 智能体卫士插件列表数据项
type ListDefenderPluginsResponseBodyData struct {

	// **参数解释**： 实例ID **取值范围**： 字符长度为1-128位
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**： 实例名称，即智能体应用名称 **取值范围**： 字符长度为1-128位
	InstanceName *string `json:"instance_name,omitempty"`

	// **参数解释**： 实例版本号，即智能体应用版本号 **取值范围**： 字符长度为1-64位
	InstanceVersion *string `json:"instance_version,omitempty"`

	// **参数解释**： 实例端口，即智能体应用监听端口 **取值范围**： 最小值0，最大值65535
	InstancePort *int32 `json:"instance_port,omitempty"`

	// **参数解释**： 智能体应用类型 **取值范围**： - openclaw：OpenClaw
	AgentAppType *string `json:"agent_app_type,omitempty"`

	// **参数解释**： 智能体运行环境 **取值范围**： - ECS - CCE
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

	// **参数解释**： 安装或卸载失败原因，仅defender_status为install_failed时返回 **取值范围**： 字符长度为0-1024位
	FailedReason *string `json:"failed_reason,omitempty"`

	// **参数解释**： 容器实例ID **取值范围**： 字符长度0-128位
	ContainerId *string `json:"container_id,omitempty"`
}

func (o ListDefenderPluginsResponseBodyData) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListDefenderPluginsResponseBodyData struct{}"
	}

	return strings.Join([]string{"ListDefenderPluginsResponseBodyData", string(data)}, " ")
}
