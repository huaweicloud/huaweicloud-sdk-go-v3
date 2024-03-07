package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SimpleDesktopInfo struct {

	// 桌面ID。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 桌面名。
	ComputerName *string `json:"computer_name,omitempty"`

	// 创建时间。
	Created *string `json:"created,omitempty"`

	// 桌面ip地址。
	IpAddress *string `json:"ip_address,omitempty"`

	// 用户名。
	UserName *string `json:"user_name,omitempty"`

	// 桌面已分配的用户信息列表。
	AttachUserInfos *[]AttachInstancesUserInfo `json:"attach_user_infos,omitempty"`

	// 权限组。
	UserGroup *string `json:"user_group,omitempty"`

	// 桌面的SID信息。
	Sid *string `json:"sid,omitempty"`

	// ou名称。
	OuName *string `json:"ou_name,omitempty"`

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 标签列表。
	Tags *[]Tag `json:"tags,omitempty"`

	// 是否处于管理员维护模式
	InMaintenanceMode *bool `json:"in_maintenance_mode,omitempty"`

	// 桌面的子网ID。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 桌面计费资源ID。
	BillResourceId *string `json:"bill_resource_id,omitempty"`
}

func (o SimpleDesktopInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimpleDesktopInfo struct{}"
	}

	return strings.Join([]string{"SimpleDesktopInfo", string(data)}, " ")
}
