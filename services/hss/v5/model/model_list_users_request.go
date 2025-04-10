package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListUsersRequest Request Object
type ListUsersRequest struct {

	// 主机ID
	HostId *string `json:"host_id,omitempty"`

	// 账号名称
	UserName *string `json:"user_name,omitempty"`

	// 主机名称
	HostName *string `json:"host_name,omitempty"`

	// 服务器私有IP
	PrivateIp *string `json:"private_ip,omitempty"`

	// 是否允许登录
	LoginPermission *bool `json:"login_permission,omitempty"`

	// 是否有root权限
	RootPermission *bool `json:"root_permission,omitempty"`

	// 主机用户组
	UserGroup *string `json:"user_group,omitempty"`

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

func (o ListUsersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListUsersRequest struct{}"
	}

	return strings.Join([]string{"ListUsersRequest", string(data)}, " ")
}
