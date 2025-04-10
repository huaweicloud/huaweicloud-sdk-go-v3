package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListPortStatisticsRequest Request Object
type ListPortStatisticsRequest struct {

	// 端口号，精确匹配
	Port *int32 `json:"port,omitempty"`

	// 端口字符串，用来进行模糊匹配
	PortString *string `json:"port_string,omitempty"`

	// 端口类型
	Type *string `json:"type,omitempty"`

	// 端口状态，包含如下： - danger：危险端口 - unknow: 无已知危险的端口
	Status *string `json:"status,omitempty"`

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 排序的key值，目前支持按照端口号port排序
	SortKey *string `json:"sort_key,omitempty"`

	// 升序还是降序，默认升序，asc
	SortDir *string `json:"sort_dir,omitempty"`

	// 每页显示数量
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量：指定返回记录的开始位置
	Offset *int32 `json:"offset,omitempty"`

	// 类别，默认为host，包含如下： - host：主机 - container：容器
	Category *string `json:"category,omitempty"`
}

func (o ListPortStatisticsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListPortStatisticsRequest struct{}"
	}

	return strings.Join([]string{"ListPortStatisticsRequest", string(data)}, " ")
}
