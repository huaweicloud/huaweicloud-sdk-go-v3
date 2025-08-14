package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ContainerCmdLogResponseInfo 容器内运行命令的日志信息
type ContainerCmdLogResponseInfo struct {

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群类型，包含以下几种： - cce：cce集群 - ali：阿里云集群 - tencent：腾讯云集群 - azure：微软云集群 - aws：亚马逊集群 - self_built_hw：华为云自建集群 - self_built_idc：IDC自建集群
	ClusterType *string `json:"cluster_type,omitempty"`

	// 日志产生的时间
	Time *int64 `json:"time,omitempty"`

	// 主机ID
	HostId *string `json:"host_id,omitempty"`

	// 主机名称
	HostName *string `json:"host_name,omitempty"`

	// 容器所在主机的私网ip
	PrivateIp *string `json:"private_ip,omitempty"`

	// 主机ip
	PublicIp *string `json:"public_ip,omitempty"`

	// 产生日志的容器名称
	ContainerName *string `json:"container_name,omitempty"`

	// 产生日志的容器id
	ContainerId *string `json:"container_id,omitempty"`

	// 执行的命令
	Cmd *string `json:"cmd,omitempty"`

	// 命令行对应的进程路径
	Path *string `json:"path,omitempty"`

	// 命令行对应的进程pid
	ProcessPid *int32 `json:"process_pid,omitempty"`

	// 执行命令的用户
	UserName *string `json:"user_name,omitempty"`

	// 执行命令的用户所属用户组
	GroupName *string `json:"group_name,omitempty"`

	// 父进程pid
	ParentProcessPid *int32 `json:"parent_process_pid,omitempty"`

	// 父进程路径
	ParentPath *string `json:"parent_path,omitempty"`
}

func (o ContainerCmdLogResponseInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ContainerCmdLogResponseInfo struct{}"
	}

	return strings.Join([]string{"ContainerCmdLogResponseInfo", string(data)}, " ")
}
