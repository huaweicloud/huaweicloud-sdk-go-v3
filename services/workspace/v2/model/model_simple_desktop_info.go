package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SimpleDesktopInfo struct {

	// domainId。
	DomainId *string `json:"domain_id,omitempty"`

	// 项目id。
	ProjectId *string `json:"project_id,omitempty"`

	// 桌面ID。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 桌面名。
	ComputerName *string `json:"computer_name,omitempty"`

	// 系统计算机名。
	OsHostName *string `json:"os_host_name,omitempty"`

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

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 标签列表。
	Tags *[]Tag `json:"tags,omitempty"`

	// 是否处于管理员维护模式。
	InMaintenanceMode *bool `json:"in_maintenance_mode,omitempty"`

	// 桌面协同资源SKU码。
	ShareResourceSku *string `json:"share_resource_sku,omitempty"`

	// 桌面类型。
	DesktopType *string `json:"desktop_type,omitempty"`

	// 桌面详细类型
	DesktopDetailType *string `json:"desktop_detail_type,omitempty"`

	// 桌面的子网ID。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 桌面计费资源ID。
	BillResourceId *string `json:"bill_resource_id,omitempty"`

	// 桌面的运行状态。
	Status *string `json:"status,omitempty"`

	// 桌面的任务状态。
	TaskStatus *string `json:"task_status,omitempty"`

	// 所属的可用区。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 桌面的连接状态
	ConnectStatus *string `json:"connect_status,omitempty"`

	// 桌面池id。
	PoolId *string `json:"pool_id,omitempty"`
}

func (o SimpleDesktopInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimpleDesktopInfo struct{}"
	}

	return strings.Join([]string{"SimpleDesktopInfo", string(data)}, " ")
}
