package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListContainerLogsRequest Request Object
type ListContainerLogsRequest struct {

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 集群id
	ClusterId *string `json:"cluster_id,omitempty"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 产生日志的容器所属的命名空间
	Namespace *string `json:"namespace,omitempty"`

	// 产生日志的容器所属pod的名称
	PodName *string `json:"pod_name,omitempty"`

	// 产生日志的容器所属pod的id
	PodId *string `json:"pod_id,omitempty"`

	// 产生日志的容器所属pod的ip
	PodIp *string `json:"pod_ip,omitempty"`

	// 产生日志的容器所在主机的ip
	HostIp *string `json:"host_ip,omitempty"`

	// 容器id
	ContainerId *string `json:"container_id,omitempty"`

	// 产生日志的容器名称
	ContainerName *string `json:"container_name,omitempty"`

	// 日志内容
	Content *string `json:"content,omitempty"`

	// 查询日志范围的最小时间
	StartTime *int64 `json:"start_time,omitempty"`

	// 查询日志范围的最大时间
	EndTime *int64 `json:"end_time,omitempty"`

	// 每页显示个数，默认为10
	Limit int32 `json:"limit"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset int32 `json:"offset"`

	// 查询cce集群容器日志时需要传的分页行号
	LineNum *string `json:"line_num,omitempty"`
}

func (o ListContainerLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListContainerLogsRequest struct{}"
	}

	return strings.Join([]string{"ListContainerLogsRequest", string(data)}, " ")
}
