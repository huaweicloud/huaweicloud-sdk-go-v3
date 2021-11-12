package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListLoadBalancersRequest struct {
	// 负载均衡器的管理状态。只支持设定为true。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 可用区。 注： 可用AZ的查询方式可用通过调用nova接口查询 /v2/{project_id}/os-availability-zone

	AvailabilityZoneList *[]string `json:"availability_zone_list,omitempty"`
	// 预留资源账单信息，默认为空表示按需计费， 非空为包周期。admin权限才能更新此字段。

	BillingInfo *[]string `json:"billing_info,omitempty"`
	// 是否开启删除保护，false不开启，默认为空都查询

	DeletionProtectionEnable *bool `json:"deletion_protection_enable,omitempty"`
	// 负载均衡器的描述信息。

	Description *[]string `json:"description,omitempty"`
	// 公网ELB实例绑定EIP。 示例如下：  \"eips\": [             {                 \"eip_id\": \"a6ded276-c88a-4c58-95e0-5b6d1d2297b3\",                 \"eip_address\": \"2001:db8:a583:86:cf24:5cc5:8117:6eaa\",                 \"ip_version\": 6             }         ] 查询时指定：eips=eip_id=XXXX

	Eips *[]string `json:"eips,omitempty"`
	// 企业项目ID。

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
	// 共享型：false 性能保障型：true

	Guaranteed *bool `json:"guaranteed,omitempty"`
	// 负载均衡器ID。

	Id *[]string `json:"id,omitempty"`
	// IP版本信息。 取值范围：4和6 4：IPv4 6：IPv6

	IpVersion *[]int32 `json:"ip_version,omitempty"`
	// 双栈实例对应v6的ip地址。

	Ipv6VipAddress *[]string `json:"ipv6_vip_address,omitempty"`
	// 双栈实例对应v6的端口。

	Ipv6VipPortId *[]string `json:"ipv6_vip_port_id,omitempty"`
	// 双栈实例对应v6的网络id 。 说明：vpc_id , vip_subnet_cidr_id, ipv6_vip_virsubnet_id不能同时为空。

	Ipv6VipVirsubnetId *[]string `json:"ipv6_vip_virsubnet_id,omitempty"`
	// 四层Flavor， 按需计费不填， 包周期由用户设置。

	L4FlavorId *[]string `json:"l4_flavor_id,omitempty"`
	// 预留弹性flavor。

	L4ScaleFlavorId *[]string `json:"l4_scale_flavor_id,omitempty"`
	// 七层Flavor， 按需计费不填， 包周期由用户设置。

	L7FlavorId *[]string `json:"l7_flavor_id,omitempty"`
	// 预留弹性flavor。

	L7ScaleFlavorId *[]string `json:"l7_scale_flavor_id,omitempty"`
	// 每页返回的个数。

	Limit *int32 `json:"limit,omitempty"`
	// 上一页最后一条记录的ID。  使用说明：  - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。

	Marker *string `json:"marker,omitempty"`
	// 后端云服务器的IP地址。

	MemberAddress *[]string `json:"member_address,omitempty"`
	// 后端云服务器对应的弹性云服务器的ID。

	MemberDeviceId *[]string `json:"member_device_id,omitempty"`
	// 负载均衡器名称。

	Name *[]string `json:"name,omitempty"`
	// 负载均衡器的操作状态。 可以为：ONLINE、OFFLINE、DEGRADED、DISABLED或NO_MONITOR。 说明 该字段为预留字段，暂未启用。

	OperatingStatus *[]string `json:"operating_status,omitempty"`
	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。 使用说明：必须与limit一起使用。

	PageReverse *bool `json:"page_reverse,omitempty"`
	// 负载均衡器的配置状态。 可以为：ACTIVE、PENDING_CREATE 或者ERROR。 说明 该字段为预留字段，暂未启用。

	ProvisioningStatus *[]string `json:"provisioning_status,omitempty"`
	// 公网IP 示例如下：  \"publicips\": [             {                 \"publicip_id\": \"a6ded276-c88a-4c58-95e0-5b6d1d2297b3\",                 \"publicip_address\": \"2001:db8:a583:86:cf24:5cc5:8117:6eaa\",                 \"publicip_ip_version\": 6             }         ] 查询时指定：publicips=publicip_id=XXXX,YYYY

	Publicips *[]string `json:"publicips,omitempty"`
	// 负载均衡器的虚拟IP。

	VipAddress *[]string `json:"vip_address,omitempty"`
	// 负载均衡器虚拟IP对应的端口ID。

	VipPortId *[]string `json:"vip_port_id,omitempty"`
	// 负载均衡器所在的子网ID，仅支持内网类型。 说明：vpc_id , vip_subnet_cidr_id, ipv6_vip_virsubnet_id不能同时为空。

	VipSubnetCidrId *[]string `json:"vip_subnet_cidr_id,omitempty"`
	// 实例对应的vpc属性。 若无，则从vip_subnet_cidr_id获取。  说明：vpc_id , vip_subnet_cidr_id, ipv6_vip_virsubnet_id不能同时为空。

	VpcId *[]string `json:"vpc_id,omitempty"`
}

func (o ListLoadBalancersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLoadBalancersRequest struct{}"
	}

	return strings.Join([]string{"ListLoadBalancersRequest", string(data)}, " ")
}
