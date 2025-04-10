package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListAppsRequest Request Object
type ListAppsRequest struct {

	// 主机id
	HostId *string `json:"host_id,omitempty"`

	// 主机名称
	HostName *string `json:"host_name,omitempty"`

	// 软件名称
	AppName *string `json:"app_name,omitempty"`

	// 主机ip
	HostIp *string `json:"host_ip,omitempty"`

	// 软件版本号
	Version *string `json:"version,omitempty"`

	// 安装目录
	InstallDir *string `json:"install_dir,omitempty"`

	// 主机所属的企业项目ID。 开通企业项目功能后才需要配置企业项目。 企业项目ID默认取值为“0”，表示默认企业项目。如果需要查询所有企业项目下的主机，请传参“all_granted_eps”。如果您只有某个企业项目的权限，则需要传递该企业项目ID，查询该企业项目下的主机，否则会因权限不足而报错。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 每页显示数量
	Limit *int32 `json:"limit,omitempty"`

	// 偏移量：指定返回记录的开始位置
	Offset *int32 `json:"offset,omitempty"`

	// 类别，默认为host，包含如下： - host：主机 - container：容器
	Category *string `json:"category,omitempty"`

	// 是否模糊匹配，默认false表示精确匹配
	PartMatch *bool `json:"part_match,omitempty"`
}

func (o ListAppsRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListAppsRequest struct{}"
	}

	return strings.Join([]string{"ListAppsRequest", string(data)}, " ")
}
