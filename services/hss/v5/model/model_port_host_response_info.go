package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// PortHostResponseInfo 开放端口的机器统计信息
type PortHostResponseInfo struct {

	// **参数解释**: 容器ID **取值范围**: 字符长度1-128位
	ContainerId *string `json:"container_id,omitempty"`

	// **参数解释**： 服务器（主机）的唯一标识ID **取值范围**： 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器IP **取值范围**: 字符长度1-128位
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 端口的监听IP地址（0.0.0.0表示监听所有网卡） **取值范围**: 支持IPv4或IPv6格式，IPv4长度7-15位，IPv6长度15-39位
	Laddr *string `json:"laddr,omitempty"`

	// **参数解释**: 占用当前端口的进程对应的可执行文件绝对路径 **取值范围**: 字符长度0-512位
	Path *string `json:"path,omitempty"`

	// **参数解释**: 占用当前端口的进程ID **取值范围**: 非负整数，最小值0（无对应进程时为0）；单位：个
	Pid *int32 `json:"pid,omitempty"`

	// **参数解释**: 端口的具体编号 **取值范围**: 1-65535（TCP/UDP标准端口范围）
	Port *int32 `json:"port,omitempty"`

	// 状态
	Status *string `json:"status,omitempty"`

	// **参数解释**: 端口类型 **取值范围**: 目前包括TCP，UDP两种
	Type *string `json:"type,omitempty"`

	// **参数解释**： 容器实例名称，只有容器类型的告警有 **取值范围**： 字符长度1-256位
	ContainerName *string `json:"container_name,omitempty"`

	// **参数解释**: 主机上安装的杀毒Agent的唯一标识ID，用于关联主机与杀毒服务 **约束限制**: 不涉及 **取值范围**: 字符长度1-64位 **默认取值**: 不涉及
	AgentId *string `json:"agent_id,omitempty"`
}

func (o PortHostResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PortHostResponseInfo struct{}"
	}

	return strings.Join([]string{"PortHostResponseInfo", string(data)}, " ")
}
