package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// ListRecycleBinLoadBalancersRequest Request Object
type ListRecycleBinLoadBalancersRequest struct {

	// **参数解释**：上一页最后一条记录的ID。  **约束限制**： - 必须与limit一起使用。 - 不指定时表示查询第一页。 - 该字段不允许为空或无效的ID。  **取值范围**：不涉及  **默认取值**：不涉及
	Marker *string `json:"marker,omitempty"`

	// **参数解释**：每页返回的个数。  **约束限制**：不涉及  **取值范围**：0-2000  **默认取值**：2000
	Limit *int32 `json:"limit,omitempty"`

	// **参数解释**：是否反向查询。  **约束限制**： - 必须与limit一起使用。 - 当page_reverse=true时，若要查询上一页，marker取值为当前页返回值的previous_marker。  **取值范围**： - true：查询上一页。 - false：查询下一页。  **默认取值**：false
	PageReverse *bool `json:"page_reverse,omitempty"`

	// 负载均衡器ID。  支持多值查询，查询条件格式：*id=xxx&id=xxx*。
	Id *[]string `json:"id,omitempty"`

	// 负载均衡器名称。  支持多值查询，查询条件格式：*name=xxx&name=xxx*。
	Name *[]string `json:"name,omitempty"`

	// 负载均衡器的描述信息。  支持多值查询，查询条件格式：*description=xxx&description=xxx*。
	Description *[]string `json:"description,omitempty"`

	// **参数解释**：负载均衡器的启用状态。  **取值范围**： - true ：启用。 - false：停用。  [不支持该字段，请勿使用。](tag:dt)
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 负载均衡器的操作状态。  取值： - ONLINE：正常运行。 - FROZEN：已冻结。  支持多值查询，查询条件格式：*operating_status=xxx&operating_status=xxx*。
	OperatingStatus *[]string `json:"operating_status,omitempty"`

	// 是否独享型LB。  取值： - false：共享型 - true：独享型  [仅支持独享型，固定为true。](tag:hws_eu,hcso_dt)
	Guaranteed *bool `json:"guaranteed,omitempty"`

	// 负载均衡器所在的VPC ID。  支持多值查询，查询条件格式：*vpc_id=xxx&vpc_id=xxx*。
	VpcId *[]string `json:"vpc_id,omitempty"`

	// 负载均衡器的IPv4对应的port ID。  支持多值查询，查询条件格式：*vip_port_id=xxx&vip_port_id=xxx*。
	VipPortId *[]string `json:"vip_port_id,omitempty"`

	// 负载均衡器的IPv4私网IP地址。  支持多值查询，查询条件格式：*vip_address=xxx&vip_address=xxx*。
	VipAddress *[]string `json:"vip_address,omitempty"`

	// 负载均衡器所在子网的IPv4子网ID，也称为该负载均衡器实例的前端子网。  支持多值查询，查询条件格式：*vip_subnet_cidr_id=xxx&vip_subnet_cidr_id=xxx*。
	VipSubnetCidrId *[]string `json:"vip_subnet_cidr_id,omitempty"`

	// 双栈类型负载均衡器的IPv6对应的port ID。  支持多值查询，查询条件格式：*ipv6_vip_port_id=xxx&ipv6_vip_port_id=xxx*。  [不支持IPv6，请勿使用。](tag:dt)
	Ipv6VipPortId *[]string `json:"ipv6_vip_port_id,omitempty"`

	// 双栈类型负载均衡器的IPv6地址。  支持多值查询，查询条件格式：*ipv6_vip_address=xxx&ipv6_vip_address=xxx*。  [不支持IPv6，请勿使用。](tag:dt)
	Ipv6VipAddress *[]string `json:"ipv6_vip_address,omitempty"`

	// 双栈类型负载均衡器所在的子网IPv6网络ID，也称为该负载均衡器实例的前端子网。  支持多值查询，查询条件格式：*ipv6_vip_virsubnet_id=xxx&ipv6_vip_virsubnet_id=xxx*。  [不支持IPv6，请勿使用。](tag:dt)
	Ipv6VipVirsubnetId *[]string `json:"ipv6_vip_virsubnet_id,omitempty"`

	// 负载均衡器绑定的EIP。例如要查询绑定以下EIP的LB： \"eips\": [     {         \"eip_id\": \"e9b72a9d-4275-455e-a724-853504e4d9c6\",         \"eip_address\": \"88.88.14.122\",         \"ip_version\": 4     } ] 可以通如下查询： eips=ip_version%3D4&eips=eip_address%3D88.88.14.122&eips=eip_id%3De9b72a9d-4275-455e-a724-853504e4d9c6  支持多值查询，查询条件格式： - eip_id作为查询条件：*eips=eip_id=xxx&eips=eip_id=xxx*。 - eip_address作为查询条件：*eips=eip_address=xxx&eips=eip_address=xxx*。 - ip_version作为查询条件：*eips=ip_version=xxx&eips=ip_version=xxx*。  注：该字段与publicips字段一致。
	Eips *[]string `json:"eips,omitempty"`

	// 负载均衡器绑定的公网IP。例如要查询绑定以下公网IP的LB： \"publicips=\": [     {         \"public_id\": \"e9b72a9d-4275-455e-a724-853504e4d9c6\",         \"public_address\": \"88.88.14.122\",         \"ip_version\": 4     } ] 可以通如下查询： publicips=ip_version%3D4&publicips=public_address%3D88.88.14.122&publicips=public_id%3De9b72a9d-4275-455e-a724-853504e4d9c6  支持多值查询，查询条件格式： - publicip_id作为查询条件： *publicips=publicip_id=xxx&publicips=publicip_id=xxx* - publicip_address作为查询条件： *publicips=publicip_address=xxx&publicips=publicip_address=xxx* - ip_version作为查询条件： *publicips=ip_version=xxx&publicips=ip_version=xxx*  注：该字段与eips字段一致。
	Publicips *[]string `json:"publicips,omitempty"`

	// 负载均衡器所在可用区列表。  支持多值查询，查询条件格式： *availability_zone_list=xxx&availability_zone_list=xxx*。
	AvailabilityZoneList *[]string `json:"availability_zone_list,omitempty"`

	// 网络型规格ID。  支持多值查询，查询条件格式：*l4_flavor_id=xxx&l4_flavor_id=xxx*。  [不支持该字段，请勿使用。](tag:hcso,hk_vdf,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b)
	L4FlavorId *[]string `json:"l4_flavor_id,omitempty"`

	// 应用型规格ID。  支持多值查询，查询条件格式：*l7_flavor_id=xxx&l7_flavor_id=xxx*。  [不支持该字段，请勿使用。](tag:hcso,hk_vdf,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b)
	L7FlavorId *[]string `json:"l7_flavor_id,omitempty"`

	// 资源账单信息。  支持多值查询，查询条件格式：*billing_info=xxx&billing_info=xxx*。  [不支持该字段，请勿使用。](tag:hws_hk,hws_eu,hws_test,hcs,hcs_sm,hcso,hk_vdf,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b,hcso_dt,dt,ocb,ctc,cmcc,tm,sbc,g42,hws_ocb,hk_sbc,hk_tm,hk_g42)
	BillingInfo *[]string `json:"billing_info,omitempty"`

	// 负载均衡器中的后端服务器对应的弹性云服务器的ID。仅用于查询条件，不作为响应参数字段。  支持多值查询，查询条件格式：*member_device_id=xxx&member_device_id=xxx*。
	MemberDeviceId *[]string `json:"member_device_id,omitempty"`

	// 负载均衡器中的后端服务器对应的弹性云服务器的IP地址。仅用于查询条件，不作为响应参数字段。  支持多值查询，查询条件格式：*member_address=xxx&member_address=xxx*。
	MemberAddress *[]string `json:"member_address,omitempty"`

	// **参数解释**：资源所属的企业项目ID。 支持多值查询，查询条件格式： *enterprise_project_id=xxx&enterprise_project_id=xxx*。  **约束限制**： 如果enterprise_project_id不传值，默认查询所有企业项目下的资源，鉴权按照细粒度权限鉴权，必须在用户组下分配elb:loadbalancers:list权限。 如果enterprise_project_id传值，鉴权按照企业项目权限鉴权，分为传入具体eps_id和all_granted_eps两种场景，前者查询指定eps_id的eps下的资源，后者查询的是所有有list权限的eps下的资源。  **取值范围**：不涉及  **默认取值**：不涉及  [不支持该字段，请勿使用。](tag:dt,hcso_dt)
	EnterpriseProjectId *[]string `json:"enterprise_project_id,omitempty"`

	// IP版本信息。  取值：4代表IPv4，6代表IPv6。  支持多值查询，查询条件格式：*ip_version=xxx&ip_version=xxx*。  [不支持IPv6，请勿设置为6。](tag:dt)
	IpVersion *[]int32 `json:"ip_version,omitempty"`

	// 是否开启删除保护，false 不开启，true 开启。[不支持该字段，请勿使用。](tag:hws_eu,g42,hk_g42)  [荷兰region不支持该字段，请勿使用。](tag:dt)
	DeletionProtectionEnable *bool `json:"deletion_protection_enable,omitempty"`

	// 下联面子网类型。  取值： - ipv4：ipv4。 - dualstack：双栈。  支持多值查询，查询条件格式： *elb_virsubnet_type=ipv4&elb_virsubnet_type=dualstack*。
	ElbVirsubnetType *[]string `json:"elb_virsubnet_type,omitempty"`

	// 是否开启弹性扩缩容。示例如下： \"autoscaling\": {             \"enable\": \"true\"         }  支持多值查询，查询条件格式：  *autoscaling=enable=true&autoscaling=enable=false*。  [不支持该字段，请勿使用。](tag:hws_eu,g42,hk_g42,hcso,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b)
	Autoscaling *[]string `json:"autoscaling,omitempty"`

	// 修改保护状态, 取值： - nonProtection: 不保护，默认值为nonProtection - consoleProtection: 控制台修改保护
	ProtectionStatus *[]string `json:"protection_status,omitempty"`

	// 负载均衡器绑定的公网IP。示例如下：  {     \"global_eips\": [         {             \"global_eip_id\": \"24000000-0000-0000-0000-100000000001\",             \"global_eip_address\": \"10.10.10.10\",             \"ip_version\": 4         }     ] }  支持多值查询，查询条件格式：  - global_eip_id作为查询条件：*global_eips=global_eip_id=xxx&global_eips=global_eip_id=xxx*。  - global_eip_address作为查询条件：*global_eips=global_eip_address=xxx&global_eips=global_eip_address=xxx*。  - ip_version作为查询条件：*global_eips=ip_version=xxx&global_eips=ip_version=xxx*。
	GlobalEips *[]string `json:"global_eips,omitempty"`

	// LB实例绑定的logtank的topic id信息，支持多值查询，查询条件格式：*log_topic_id=xxx&log_topic_id=xxx*。
	LogTopicId *string `json:"log_topic_id,omitempty"`

	// LB实例绑定的logtank的group id信息，支持多值查询，查询条件格式：*log_group_id=xxx&log_group_id=xxx*。
	LogGroupId *string `json:"log_group_id,omitempty"`
}

func (o ListRecycleBinLoadBalancersRequest) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ListRecycleBinLoadBalancersRequest struct{}"
	}

	return strings.Join([]string{"ListRecycleBinLoadBalancersRequest", string(data)}, " ")
}
