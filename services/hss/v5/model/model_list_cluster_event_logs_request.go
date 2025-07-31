package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterEventLogsRequest Request Object
type ListClusterEventLogsRequest struct {

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 集群id
	ClusterId string `json:"cluster_id"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 产生事件的命名空间
	Namespace *string `json:"namespace,omitempty"`

	// 事件名称
	EventName *string `json:"event_name,omitempty"`

	// 事件类型，包含以下几种： - Warning 警告事件 - Normal 普通事件
	EventType *string `json:"event_type,omitempty"`

	// 产生事件的资源类型
	ResourceType *string `json:"resource_type,omitempty"`

	// 产生事件的资源名称
	ResourceName *string `json:"resource_name,omitempty"`

	// 事件的触发原因
	Reason *string `json:"reason,omitempty"`

	// 查询日志范围的最小时间
	StartTime int64 `json:"start_time"`

	// 查询日志范围的最大时间
	EndTime int64 `json:"end_time"`

	// 每页显示个数，默认为10
	Limit int32 `json:"limit"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset int32 `json:"offset"`

	// 查询cce集群事件时需要传的分页行号
	LineNum *string `json:"line_num,omitempty"`
}

func (o ListClusterEventLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterEventLogsRequest struct{}"
	}

	return strings.Join([]string{"ListClusterEventLogsRequest", string(data)}, " ")
}
