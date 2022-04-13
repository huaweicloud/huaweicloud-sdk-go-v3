package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Request Object
type ListLoadBalancersRequest struct {
	// 上一页最后一条记录的ID。  使用说明： - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。

	Marker *string `json:"marker,omitempty"`
	// 每页返回的个数。

	Limit *int32 `json:"limit,omitempty"`
	// 分页的顺序，true表示从后往前分页，false表示从前往后分页，默认为false。  使用说明： - 必须与limit一起使用。

	PageReverse *bool `json:"page_reverse,omitempty"`
	// 负载均衡器ID。  支持多值查询，查询条件格式：*id=xxx&id=xxx*。

	Id *[]string `json:"id,omitempty"`
	// 负载均衡器名称。  支持多值查询，查询条件格式：*name=xxx&name=xxx*。

	Name *[]string `json:"name,omitempty"`
	// 负载均衡器的描述信息。  支持多值查询，查询条件格式：*description=xxx&description=xxx*。

	Description *[]string `json:"description,omitempty"`
	// 负载均衡器的管理状态。  不支持该字段，请勿使用。

	AdminStateUp *bool `json:"admin_state_up,omitempty"`
	// 负载均衡器的配置状态。取值： - ACTIVE：使用中。 - PENDING_DELETE：删除中。  支持多值查询，查询条件格式：*provisioning_status=xxx&provisioning_status=xxx*。

	ProvisioningStatus *[]string `json:"provisioning_status,omitempty"`
	// 负载均衡器的操作状态。取值： - ONLINE：正常运行。 - FROZEN：已冻结。  支持多值查询，查询条件格式：*operating_status=xxx&operating_status=xxx*。

	OperatingStatus *[]string `json:"operating_status,omitempty"`
	// 是否独享型LB，取值： - false：共享型 - true：独享型

	Guaranteed *bool `json:"guaranteed,omitempty"`
	// 负载均衡器所在的VPC ID。  支持多值查询，查询条件格式：*vpc_id=xxx&vpc_id=xxx*。

	VpcId *[]string `json:"vpc_id,omitempty"`
	// 负载均衡器的IPv4对应的port ID。  支持多值查询，查询条件格式：*vip_port_id=xxx&vip_port_id=xxx*。

	VipPortId *[]string `json:"vip_port_id,omitempty"`
	// 负载均衡器的IPv4虚拟IP地址。  支持多值查询，查询条件格式：*vip_address=xxx&vip_address=xxx*。

	VipAddress *[]string `json:"vip_address,omitempty"`
	// 负载均衡器所在子网的IPv4子网ID。  支持多值查询，查询条件格式：*vip_subnet_cidr_id=xxx&vip_subnet_cidr_id=xxx*。

	VipSubnetCidrId *[]string `json:"vip_subnet_cidr_id,omitempty"`
	// 双栈类型负载均衡器的IPv6对应的port ID。  支持多值查询，查询条件格式：*ipv6_vip_port_id=xxx&ipv6_vip_port_id=xxx*。   [不支持IPv6，请勿使用。](tag:dt,dt_test)

	Ipv6VipPortId *[]string `json:"ipv6_vip_port_id,omitempty"`
	// 双栈类型负载均衡器的IPv6地址。  支持多值查询，查询条件格式：*ipv6_vip_address=xxx&ipv6_vip_address=xxx*。   [不支持IPv6，请勿使用。](tag:dt,dt_test)

	Ipv6VipAddress *[]string `json:"ipv6_vip_address,omitempty"`
	// 双栈类型负载均衡器所在的子网IPv6网络ID。  支持多值查询，查询条件格式：*ipv6_vip_virsubnet_id=xxx&ipv6_vip_virsubnet_id=xxx*。   [不支持IPv6，请勿使用。](tag:dt,dt_test)

	Ipv6VipVirsubnetId *[]string `json:"ipv6_vip_virsubnet_id,omitempty"`
	// 负载均衡器绑定的EIP ID。示例如下： \"eips\": [             {                 \"eip_id\": \"e9b72a9d-4275-455e-a724-853504e4d9c6\",                 \"eip_address\": \"88.88.14.122\",                 \"ip_version\": 4             }         ]  支持多值查询，查询条件格式： - eip_id作为查询条件：*eips=eip_id=xxx&eips=eip_id=xxx*。 - eip_address作为查询条件：*eips=eip_address=xxx&eips=eip_address=xxx*。 - ip_version作为查询条件：*eips=ip_version=xxx&eips=ip_version=xxx*。  注：该字段与publicips字段一致。

	Eips *[]string `json:"eips,omitempty"`
	// 负载均衡器绑定的公网IP。示例如下：  \"publicips\": [                 {                     \"publicip_id\": \"e9b72a9d-4275-455e-a724-853504e4d9c6\",                     \"publicip_address\": \"88.88.14.122\",                     \"ip_version\": 4                 }             ]  支持多值查询，查询条件格式： - publicip_id作为查询条件：*publicips=publicip_id=xxx&publicips=publicip_id=xxx*。 - publicip_address作为查询条件：*publicips=publicip_address=xxx&publicips=publicip_address=xxx*。 - ip_version作为查询条件：*publicips=ip_version=xxx&publicips=ip_version=xxx*。  注：该字段与eips字段一致。

	Publicips *[]string `json:"publicips,omitempty"`
	// 负载均衡器所在可用区列表。  支持多值查询，查询条件格式：*availability_zone_list=xxx&availability_zone_list=xxx*。

	AvailabilityZoneList *[]string `json:"availability_zone_list,omitempty"`
	// 四层Flavor ID。  支持多值查询，查询条件格式：*l4_flavor_id=xxx&l4_flavor_id=xxx*。

	L4FlavorId *[]string `json:"l4_flavor_id,omitempty"`
	// 四层弹性Flavor ID。  支持多值查询，查询条件格式：*l4_scale_flavor_id=xxx&l4_scale_flavor_id=xxx*。  不支持该字段，请勿使用。

	L4ScaleFlavorId *[]string `json:"l4_scale_flavor_id,omitempty"`
	// 七层Flavor ID。  支持多值查询，查询条件格式：*l7_flavor_id=xxx&l7_flavor_id=xxx*。

	L7FlavorId *[]string `json:"l7_flavor_id,omitempty"`
	// 七层弹性Flavor ID。  支持多值查询，查询条件格式：*l7_scale_flavor_id=xxx&l7_scale_flavor_id=xxx*。  不支持该字段，请勿使用。

	L7ScaleFlavorId *[]string `json:"l7_scale_flavor_id,omitempty"`
	// 资源账单信息。admin权限才能更新此字段。  支持多值查询，查询条件格式：*billing_info=xxx&billing_info=xxx*。   [不支持该字段，请勿使用。](tag:dt,dt_test)

	BillingInfo *[]string `json:"billing_info,omitempty"`
	// 负载均衡器中的后端云服务器对应的弹性云服务器的ID。仅用于查询条件，不作为响应参数字段。  支持多值查询，查询条件格式：*member_device_id=xxx&member_device_id=xxx*。

	MemberDeviceId *[]string `json:"member_device_id,omitempty"`
	// 负载均衡器中的后端云服务器对应的弹性云服务器的IP地址。仅用于查询条件，不作为响应参数字段。  支持多值查询，查询条件格式：*member_address=xxx&member_address=xxx*。

	MemberAddress *[]string `json:"member_address,omitempty"`
	// 负载均衡器所属的企业项目ID。  支持多值查询，查询条件格式：*enterprise_project_id=xxx&enterprise_project_id=xxx*。  [不支持该字段，请勿使用。](tag:dt,dt_test,hcso_dt)

	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`
	// IP版本信息。 取值：4代表IPv4，6代表IPv6。  支持多值查询，查询条件格式：*ip_version=xxx&ip_version=xxx*。   [不支持IPv6，请勿设置为6。](tag:dt,dt_test)

	IpVersion *[]int32 `json:"ip_version,omitempty"`
	// 是否开启删除保护，false不开启，true开启。

	DeletionProtectionEnable *bool `json:"deletion_protection_enable,omitempty"`
	// 下联面子网类型。取值： - ipv4：ipv4。 - dualstack：双栈。  支持多值查询，查询条件格式： *elb_virsubnet_type=ipv4&elb_virsubnet_type=dualstack*。

	ElbVirsubnetType *[]string `json:"elb_virsubnet_type,omitempty"`
	// 是否开启弹性扩缩容。示例如下： \"autoscaling\": {             \"enable\": \"true\"         }  支持多值查询，查询条件格式：  *autoscaling=enable=true&autoscaling=enable=false*。

	Autoscaling *[]string `json:"autoscaling,omitempty"`
}

func (o ListLoadBalancersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListLoadBalancersRequest struct{}"
	}

	return strings.Join([]string{"ListLoadBalancersRequest", string(data)}, " ")
}
