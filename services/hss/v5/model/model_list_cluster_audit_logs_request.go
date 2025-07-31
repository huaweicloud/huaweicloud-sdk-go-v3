package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListClusterAuditLogsRequest Request Object
type ListClusterAuditLogsRequest struct {

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 集群id
	ClusterId string `json:"cluster_id"`

	// 集群名称
	ClusterName *string `json:"cluster_name,omitempty"`

	// 主机名称
	HostName *string `json:"host_name,omitempty"`

	// 主机id
	HostId *string `json:"host_id,omitempty"`

	// 主机ip
	HostIp *string `json:"host_ip,omitempty"`

	// 审计日志对应的动作，包含以下几种： - create 创建资源 - delete 删除资源 - deletecollection 批量资源集合 - patch 修改资源 - update 更新资源 - get 获取资源 - list 获取资源列表 - watch 监控资源
	Verb *string `json:"verb,omitempty"`

	// 查询日志范围的最小时间
	StartTime int64 `json:"start_time"`

	// 查询日志范围的最大时间
	EndTime int64 `json:"end_time"`

	// 每页显示个数，默认为10
	Limit int32 `json:"limit"`

	// 偏移量：指定返回记录的开始位置，必须为数字，取值范围为大于或等于0，默认0
	Offset int32 `json:"offset"`

	// 查询cce集群日志时需要传的分页行号
	LineNum *string `json:"line_num,omitempty"`
}

func (o ListClusterAuditLogsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListClusterAuditLogsRequest struct{}"
	}

	return strings.Join([]string{"ListClusterAuditLogsRequest", string(data)}, " ")
}
