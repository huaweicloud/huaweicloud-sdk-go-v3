package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// PoolDesktopsDetailInfo 池桌面详情。
type PoolDesktopsDetailInfo struct {

	// 桌面ID。
	DesktopId *string `json:"desktop_id,omitempty"`

	// 桌面名。
	ComputerName *string `json:"computer_name,omitempty"`

	// 系统计算机名。
	OsHostName *string `json:"os_host_name,omitempty"`

	// 桌面IP地址列表。
	Addresses map[string][]AddressInfo `json:"addresses,omitempty"`

	// IP地址列表。
	IpAddresses *[]string `json:"ip_addresses,omitempty"`

	// 系统计算机IPV4。
	Ipv4 *string `json:"ipv4,omitempty"`

	// 系统计算机IPV6。
	Ipv6 *string `json:"ipv6,omitempty"`

	// 用户列表。
	UserList *[]string `json:"user_list,omitempty"`

	// 用户组列表。
	UserGroupList *[]string `json:"user_group_list,omitempty"`

	// 桌面类型。  - DEDICATED：专属桌面。 - SHARED: 多用户桌面 - POOLED: 池桌面
	DesktopType *string `json:"desktop_type,omitempty"`

	// 桌面元数据。   - image_name 创建桌面的镜像名称。 - bill_resource_id 镜像计费资源ID。 - metering.image_id 镜像ID。 - metering.resourcespeccode 桌面资源编码。 - metering.resourcetype 桌面资源类型。 - os_bit 操作系统位数：32或64。 - os_type 操作系统类型：Linux、Windows或Others。 - desktop_os_version 操作系统版本。
	Metadata map[string]string `json:"metadata,omitempty"`

	Flavor *FlavorInfo `json:"flavor,omitempty"`

	// 桌面状态。
	Status *string `json:"status,omitempty"`

	// 任务状态。  - scheduling：创建中，正在进行调度。 - block_device_mapping：创建中，正在准备磁盘。 - networking：创建中，正在准备网络。 - spawning：创建中，正在内部创建。 - rebooting：重启中。 - reboot_pending：重启中，正在下发重启。 - reboot_started：重启中，开始内部重启。 - rebooting_hard：强制重启中。 - reboot_pending_hard：强制重启中，正在下发重启。 - reboot_started_hard：强制重启中，开始内部重启。 - rebuilding：重建中。 - rebuild_block_device_mapping：重建中，正在准备磁盘。 - rebuild_spawning：重建中，正在内部重建。 - migrating：热迁移中。 - resize_prep：调整规格中，正在准备阶段。 - resize_migrating：调整规格中，正在迁移阶段。 - resize_migrated：调整规格中，已经完成迁移。 - resize_finish：调整规格中，正在完成调整。 - resize_reverting：调整规格中，正在回退调整。 - powering-off：停止中。 - powering-on：启动中。 - deleting：删除中。 - deleteFailed：删除失败。 - updating: 更新中。 - desktopNetworkChanging: 切换网络中。
	TaskStatus *string `json:"task_status,omitempty"`

	// 是否处于维护模式,true表示维护模式，false表示不处于维护模式。
	InMaintenanceMode *bool `json:"in_maintenance_mode,omitempty"`

	// 桌面创建时间。
	Created *string `json:"created,omitempty"`

	// 桌面安全组。 仅适用于查询单个桌面详情接口。
	SecurityGroups *[]SecurityGroupInfo `json:"security_groups,omitempty"`

	// 桌面的登录状态。  - UNREGISTER：表示桌面未注册时的状态（桌面启动后，会自动注册）。关机后也会出现未注册的状态。 - REGISTERED：表示桌面注册以后，等待用户连接的状态。 - CONNECTED：表示用户已经成功登录，正在使用桌面。 - DISCONNECTED：表示桌面与客户端断开会话后显示的状态，可能为关闭客户端窗口，或客户端与桌面网络断开引起。
	LoginStatus *string `json:"login_status,omitempty"`

	// 桌面所属用户。
	UserName *string `json:"user_name,omitempty"`

	// 桌面已分配的用户信息列表。
	AttachUserInfos *[]AttachInstancesUserInfo `json:"attach_user_infos,omitempty"`

	// 产品ID。
	ProductId *string `json:"product_id,omitempty"`

	// 桌面协同资源SKU码。
	ShareResourceSku *string `json:"share_resource_sku,omitempty"`

	RootVolume *VolumeDetail `json:"root_volume,omitempty"`

	// 数据盘列表。
	DataVolumes *[]VolumeDetail `json:"data_volumes,omitempty"`

	// 桌面用户所属的用户组。  - sudo：Linux管理员组。 - default：Linux默认用户组。 - administrators：Windows管理员组。管理员拥有对该桌面的完全访问权，可以做任何需要的更改（禁用操作除外）。 - users：Windows标准用户组。标准用户可以使用大多数软件，并可以更改不影响其他用户的系统设置。
	UserGroup *string `json:"user_group,omitempty"`

	// 可用分区。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 站点类型。
	SiteType *string `json:"site_type,omitempty"`

	// 站点名字。
	SiteName *string `json:"site_name,omitempty"`

	Product *ProductInfo `json:"product,omitempty"`

	// 创建桌面时加入的OU名称。
	OuName *string `json:"ou_name,omitempty"`

	// 操作系统版本号。
	OsVersion *string `json:"os_version,omitempty"`

	// SID。
	Sid *string `json:"sid,omitempty"`

	// 包周期产品的订单ID。
	OrderId *string `json:"order_id,omitempty"`

	// 桌面标签列表。
	Tags *[]Tag `json:"tags,omitempty"`

	// 是否开通互联网，true：已开通，false：未开通。
	IsSupportInternet *bool `json:"is_support_internet,omitempty"`

	// 上网方式。 - NAT：表示NAT上网方式。 - EIP：表示EIP上网方式。 - BOTH：表示两种上网方式都支持。
	InternetMode *PoolDesktopsDetailInfoInternetMode `json:"internet_mode,omitempty"`

	// 桌面使用的上网方式列表。 - NAT：表示NAT上网方式。 - EIP：表示EIP上网方式。
	InternetModeList *[]string `json:"internet_mode_list,omitempty"`

	// 桌面是否正在绑定EIP。
	IsAttachingEip *bool `json:"is_attaching_eip,omitempty"`

	// 分配状态。 - ATTACHED：已分配。 - UNATTACH：未分配 表示未关联。 - DEATTACHED：已解分配。 - ATTACHING：分配中。 - DEATTACHING：解分配中。 - ATTACHFAIL：分配失败。 - DEATTACHFAIL：解分配失败。 - WAITING：等待被分配中,描述从批量分配（解分配）下发到转入分配（解分配）的中间状态 同时方便单个关联流程的状态独立性。 - ATTACH_FAIL_CAN_ATTACH_AGAIN：分配失败,还可以再关联。 - DEATTACH_FAIL_CAN_DEATTACH_AGAIN：解分配失败,还可以再解分配。
	AttachState *PoolDesktopsDetailInfoAttachState `json:"attach_state,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 桌面的子网ID。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 桌面计费资源ID。
	BillResourceId *string `json:"bill_resource_id,omitempty"`

	// 桌面任务进度， 取值范围0-100以及null，null表示该桌面无任务，0-100表明该任务进度的百分比。
	Process *int32 `json:"process,omitempty"`

	// 整机实例根资源ID。
	RootResourceId *string `json:"root_resource_id,omitempty"`

	HourPackageInfo *HourPackageInfo `json:"hour_package_info,omitempty"`

	// 桌面与桌面池不一致的规格类型: - PRODUCT: 产品ID不一致 - IMAGE: 镜像ID不一致
	InconsistentTypes *[]string `json:"inconsistent_types,omitempty"`
}

func (o PoolDesktopsDetailInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "PoolDesktopsDetailInfo struct{}"
	}

	return strings.Join([]string{"PoolDesktopsDetailInfo", string(data)}, " ")
}

type PoolDesktopsDetailInfoInternetMode struct {
	value string
}

type PoolDesktopsDetailInfoInternetModeEnum struct {
	NAT  PoolDesktopsDetailInfoInternetMode
	EIP  PoolDesktopsDetailInfoInternetMode
	BOTH PoolDesktopsDetailInfoInternetMode
}

func GetPoolDesktopsDetailInfoInternetModeEnum() PoolDesktopsDetailInfoInternetModeEnum {
	return PoolDesktopsDetailInfoInternetModeEnum{
		NAT: PoolDesktopsDetailInfoInternetMode{
			value: "NAT",
		},
		EIP: PoolDesktopsDetailInfoInternetMode{
			value: "EIP",
		},
		BOTH: PoolDesktopsDetailInfoInternetMode{
			value: "BOTH",
		},
	}
}

func (c PoolDesktopsDetailInfoInternetMode) Value() string {
	return c.value
}

func (c PoolDesktopsDetailInfoInternetMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoolDesktopsDetailInfoInternetMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}

type PoolDesktopsDetailInfoAttachState struct {
	value string
}

type PoolDesktopsDetailInfoAttachStateEnum struct {
	ATTACHED                         PoolDesktopsDetailInfoAttachState
	UNATTACH                         PoolDesktopsDetailInfoAttachState
	DEATTACHED                       PoolDesktopsDetailInfoAttachState
	ATTACHING                        PoolDesktopsDetailInfoAttachState
	DEATTACHING                      PoolDesktopsDetailInfoAttachState
	ATTACHFAIL                       PoolDesktopsDetailInfoAttachState
	DEATTACHFAIL                     PoolDesktopsDetailInfoAttachState
	WAITING                          PoolDesktopsDetailInfoAttachState
	ATTACH_FAIL_CAN_ATTACH_AGAIN     PoolDesktopsDetailInfoAttachState
	DEATTACH_FAIL_CAN_DEATTACH_AGAIN PoolDesktopsDetailInfoAttachState
}

func GetPoolDesktopsDetailInfoAttachStateEnum() PoolDesktopsDetailInfoAttachStateEnum {
	return PoolDesktopsDetailInfoAttachStateEnum{
		ATTACHED: PoolDesktopsDetailInfoAttachState{
			value: "ATTACHED",
		},
		UNATTACH: PoolDesktopsDetailInfoAttachState{
			value: "UNATTACH",
		},
		DEATTACHED: PoolDesktopsDetailInfoAttachState{
			value: "DEATTACHED",
		},
		ATTACHING: PoolDesktopsDetailInfoAttachState{
			value: "ATTACHING",
		},
		DEATTACHING: PoolDesktopsDetailInfoAttachState{
			value: "DEATTACHING",
		},
		ATTACHFAIL: PoolDesktopsDetailInfoAttachState{
			value: "ATTACHFAIL",
		},
		DEATTACHFAIL: PoolDesktopsDetailInfoAttachState{
			value: "DEATTACHFAIL",
		},
		WAITING: PoolDesktopsDetailInfoAttachState{
			value: "WAITING",
		},
		ATTACH_FAIL_CAN_ATTACH_AGAIN: PoolDesktopsDetailInfoAttachState{
			value: "ATTACH_FAIL_CAN_ATTACH_AGAIN",
		},
		DEATTACH_FAIL_CAN_DEATTACH_AGAIN: PoolDesktopsDetailInfoAttachState{
			value: "DEATTACH_FAIL_CAN_DEATTACH_AGAIN",
		},
	}
}

func (c PoolDesktopsDetailInfoAttachState) Value() string {
	return c.value
}

func (c PoolDesktopsDetailInfoAttachState) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *PoolDesktopsDetailInfoAttachState) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
