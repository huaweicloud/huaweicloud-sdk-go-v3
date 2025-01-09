package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type SimpleDesktopInfoDetail struct {

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

	// 企业项目ID
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 标签列表。
	Tags *[]Tag `json:"tags,omitempty"`

	// 是否处于管理员维护模式
	InMaintenanceMode *bool `json:"in_maintenance_mode,omitempty"`

	// 桌面协同资源SKU码
	ShareResourceSku *string `json:"share_resource_sku,omitempty"`

	// 桌面类型
	DesktopType *string `json:"desktop_type,omitempty"`

	// 桌面的子网ID。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 桌面计费资源ID。
	BillResourceId *string `json:"bill_resource_id,omitempty"`

	// 运行状态
	Status *string `json:"status,omitempty"`

	// 桌面的任务状态。
	TaskStatus *string `json:"task_status,omitempty"`

	// 系統状态
	InstanceStatus *string `json:"instance_status,omitempty"`

	// 连接状态
	ConnectStatus *string `json:"connect_status,omitempty"`

	// 套餐名称
	ProductName *string `json:"product_name,omitempty"`

	// AccessAgent版本号
	AgentVersion *string `json:"agent_version,omitempty"`

	// 租户ID
	TenantId *string `json:"tenant_id,omitempty"`

	// 租户名称
	TenantName *string `json:"tenant_name,omitempty"`

	// 资源池ID
	ResourcePoolId *string `json:"resource_pool_id,omitempty"`

	// 操作系统类型：Linux、Windows或Others。
	OsType *string `json:"os_type,omitempty"`

	// 智能休眠策略数。
	HibernatePolicyNum *int32 `json:"hibernate_policy_num,omitempty"`

	// 是否处于智能休眠中
	IsAutoHibernate *int32 `json:"is_auto_hibernate,omitempty"`

	// 所属的可用区。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 专享主机ID。
	ExclusiveHostId *string `json:"exclusive_host_id,omitempty"`

	// 云办公主机ID。
	DehId *string `json:"deh_id,omitempty"`
}

func (o SimpleDesktopInfoDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "SimpleDesktopInfoDetail struct{}"
	}

	return strings.Join([]string{"SimpleDesktopInfoDetail", string(data)}, " ")
}
