package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HwcEip 云弹性公网IP
type HwcEip struct {

	// 弹性公网IP唯一标识
	Id string `json:"id"`

	// 弹性公网IP名称
	Alias string `json:"alias"`

	// DDoss或CFW开启状态：OPEN | CLOSE
	ProtectedStatus string `json:"protected_status"`

	// 项目ID
	ProjectId string `json:"project_id"`

	// 企业项目ID。
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// IP版本信息 取值范围： 4：公网IP地址为public_ip_address地址 6：公网IP地址为public_ipv6_address地址
	IpVersion int32 `json:"ip_version"`

	// 弹性公网IP或者IPv6端口的地址
	PublicIpAddress *string `json:"public_ip_address,omitempty"`

	// IPv4时无此字段，IPv6时为申请到的弹性公网IP
	PublicIpv6Address *string `json:"public_ipv6_address,omitempty"`

	// 弹性公网IP的网络类型, 包括公共池类型，如5_bgp/5_sbgp...，和用户购买的专属池。
	PublicipPoolName *string `json:"publicip_pool_name,omitempty"`

	// 公网IP所属网络的ID，publicip_pool_name对应的网络ID
	PublicipPoolId *string `json:"publicip_pool_id,omitempty"`

	// 弹性公网IP的状态 取值范围： FREEZED：冻结 BIND_ERROR：绑定失败 BINDING：绑定中 PENDING_DELETE：释放中 PENDING_CREATE：创建中 NOTIFYING：创建中 NOTIFY_DELETE：释放中 PENDING_UPDATE：更新中 DOWN：未绑定 ACTIVE：绑定 ELB：绑定ELB VPN：绑定VPN ERROR：失败
	Status string `json:"status"`

	// 弹性公网IP描述信息
	Description *string `json:"description,omitempty"`

	// 功能说明：用户标签。（默认不显示）
	Tags *[]string `json:"tags,omitempty"`

	// 弹性公网IP类型 枚举值： EIP DUALSTACK DUALSTACK_SUBNET
	Type *string `json:"type,omitempty"`

	Vnic *HwcEipVnic `json:"vnic,omitempty"`

	Bandwidth *HwcEipBandwidth `json:"bandwidth,omitempty"`

	// 记录公网IP当前的冻结状态 约束：metadata类型，标识欠费冻结、公安冻结 取值范围： police locked
	LockStatus *string `json:"lock_status,omitempty"`

	// 公网IP绑定的实例类型 取值范围： PORT NATGW ELB ELBV1 VPN null
	AssociateInstanceType *string `json:"associate_instance_type,omitempty"`

	// 公网IP绑定的实例ID
	AssociateInstanceId *string `json:"associate_instance_id,omitempty"`

	// 表示此publicip可以加入的共享带宽类型列表，如果为空列表，则表示该 publicip不能加入任何共享带宽 约束：publicip只能加入到有该带宽类型的共享带宽中
	AllowShareBandwidthTypes *[]string `json:"allow_share_bandwidth_types,omitempty"`

	// 资产创建UTC时间 格式：yyyy-MM-ddTHH:mm:ssZ
	CreatedAt *string `json:"created_at,omitempty"`

	// 资产更新UTC时间 格式：yyyy-MM-ddTHH:mm:ssZ
	UpdatedAt *string `json:"updated_at,omitempty"`

	// 表示中心站点资产或者边缘站点资产 取值范围： center、边缘站点名称 约束：publicip只能绑定该字段相同的资产
	PublicBorderGroup *string `json:"public_border_group,omitempty"`
}

func (o HwcEip) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HwcEip struct{}"
	}

	return strings.Join([]string{"HwcEip", string(data)}, " ")
}
