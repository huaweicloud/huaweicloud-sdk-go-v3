package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPasswordComplexityRequest Request Object
type ListPasswordComplexityRequest struct {

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 服务器名称
	HostName *string `json:"host_name,omitempty"`

	// 服务器IP地址
	HostIp *string `json:"host_ip,omitempty"`

	// 主机id，不赋值时，查租户所有主机
	HostId *string `json:"host_id,omitempty"`

	// 结果类型  - \"unhandled\"#未忽略的  - \"ignored\"#已忽略的
	ResultType *string `json:"result_type,omitempty"`

	// 每页显示数量
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量：指定返回记录的开始位置
	Offset *int32 `json:"offset,omitempty"`
}

func (o ListPasswordComplexityRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPasswordComplexityRequest struct{}"
	}

	return strings.Join([]string{"ListPasswordComplexityRequest", string(data)}, " ")
}
