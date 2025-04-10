package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPortsRequest Request Object
type ListPortsRequest struct {

	// 主机id
	HostId string `json:"host_id"`

	// 主机名称
	HostName *string `json:"host_name,omitempty"`

	// 主机ip
	HostIp *string `json:"host_ip,omitempty"`

	// 端口号
	Port *int32 `json:"port,omitempty"`

	// 端口类型：目前包括TCP，UDP两种
	Type *string `json:"type,omitempty"`

	// 端口状态，包含如下： - danger：危险端口 - unknow: 无已知危险的端口
	Status *string `json:"status,omitempty"`

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 每页显示数量
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量：指定返回记录的开始位置
	Offset *int32 `json:"offset,omitempty"`

	// 类别，默认为host，包含如下： - host：主机 - container：容器
	Category *string `json:"category,omitempty"`
}

func (o ListPortsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPortsRequest struct{}"
	}

	return strings.Join([]string{"ListPortsRequest", string(data)}, " ")
}
