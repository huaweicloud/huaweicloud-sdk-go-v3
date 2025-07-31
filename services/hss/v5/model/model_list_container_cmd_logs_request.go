package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListContainerCmdLogsRequest Request Object
type ListContainerCmdLogsRequest struct {

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 产生日志的容器所在主机的名称
	HostName *string `json:"host_name,omitempty"`

	// 产生日志的容器所在主机的id
	HostId *string `json:"host_id,omitempty"`

	// 产生日志的容器所在主机的ip
	HostIp *string `json:"host_ip,omitempty"`

	// 容器id
	ContainerId *string `json:"container_id,omitempty"`

	// 产生日志的容器名称
	ContainerName *string `json:"container_name,omitempty"`

	// 运行的命令行
	Cmd *string `json:"cmd,omitempty"`

	// 命令行对应的进程路径
	Path *string `json:"path,omitempty"`

	// 查询日志范围的最小时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 查询日志范围的最大时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 每页显示个数，默认为10
	Limit int32 `json:"limit"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset int32 `json:"offset"`
}

func (o ListContainerCmdLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListContainerCmdLogsRequest struct{}"
	}

	return strings.Join([]string{"ListContainerCmdLogsRequest", string(data)}, " ")
}
