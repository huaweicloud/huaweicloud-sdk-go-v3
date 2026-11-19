package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ProcessResponseInfo 进程信息
type ProcessResponseInfo struct {

	// **参数解释**: 主机上安装的杀毒Agent的唯一标识ID，用于关联主机与杀毒服务 **取值范围**: 字符长度1-64位
	AgentId *string `json:"agent_id,omitempty"`

	// **参数解释**： 服务器（主机）的唯一标识ID **取值范围**： 字符长度1-64位
	HostId *string `json:"host_id,omitempty"`

	// **参数解释**: 服务器名称 **取值范围**: 字符长度1-256位
	HostName *string `json:"host_name,omitempty"`

	// **参数解释**: 服务器IP **取值范围**: 字符长度1-128位
	HostIp *string `json:"host_ip,omitempty"`

	// **参数解释**： 进程路径 **取值范围**： 字符长度1-256位
	ProcessPath *string `json:"process_path,omitempty"`

	// **参数解释** 进程启动参数 **取值范围** 字符长度1-256位
	LaunchParams *string `json:"launch_params,omitempty"`

	// **参数解释**: 进程启动时间 **取值范围**: 最小值0，最大值9223372036854775807
	LaunchTime *int64 `json:"launch_time,omitempty"`

	// **参数解释**: 用户名 **取值范围**: 字符长度1-64位
	UserName *string `json:"user_name,omitempty"`

	// **参数解释** 运行权限，对应mode **取值范围** 字符长度1-256位
	RunPermission *string `json:"run_permission,omitempty"`

	// **参数解释**： 进程ID **取值范围**： 最小值0，最大值2147483647
	ProcessPid *int32 `json:"process_pid,omitempty"`

	// **参数解释**: 父进程id **取值范围**: 最小值0，最大值2147483647
	ParentProcessPid *int32 `json:"parent_process_pid,omitempty"`

	// **参数解释** 父进程名称 **取值范围** 字符长度1-256位
	ParentProcessName *string `json:"parent_process_name,omitempty"`

	// **参数解释** 文件hash **取值范围** 字符长度1-256位
	Hash *string `json:"hash,omitempty"`

	// **参数解释**: 容器ID **取值范围**: 字符长度1-128位
	ContainerId *string `json:"container_id,omitempty"`

	// **参数解释**： 容器实例名称 **取值范围**： 字符长度1-256位
	ContainerName *string `json:"container_name,omitempty"`

	// **参数解释** Pod ID **取值范围** 字符长度1-128位
	PodId *string `json:"pod_id,omitempty"`

	// **参数解释** Pod名称 **取值范围** 字符长度1-256位
	PodName *string `json:"pod_name,omitempty"`
}

func (o ProcessResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ProcessResponseInfo struct{}"
	}

	return strings.Join([]string{"ProcessResponseInfo", string(data)}, " ")
}
