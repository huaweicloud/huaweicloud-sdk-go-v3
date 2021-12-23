package model

import (
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/RandolphCYG/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 负载均衡器的详细信息。
type LoadBalancer struct {
	// 负载均衡器ID。

	Id string `json:"id"`
	// 负载均衡器描述信息。

	Description string `json:"description"`
	// 负载均衡器的配置状态。取值： - ACTIVE：使用中。 - PENDING_DELETE：删除中。

	ProvisioningStatus string `json:"provisioning_status"`
	// 负载均衡器的管理状态。固定为true。

	AdminStateUp bool `json:"admin_state_up"`
	// 负载均衡器的生产者名称。固定为vlb。

	Provider string `json:"provider"`
	// 负载均衡器直接关联的后端云服务器组的ID列表。

	Pools []PoolRef `json:"pools"`
	// 负载均衡器关联的监听器的ID列表。

	Listeners []ListenerRef `json:"listeners"`
	// 负载均衡器的操作状态。取值： - ONLINE：在线。

	OperatingStatus string `json:"operating_status"`
	// 负载均衡器的名称。

	Name string `json:"name"`
	// 负载均衡器所属的项目ID。

	ProjectId string `json:"project_id"`
	// 负载均衡器所在子网的IPv4子网ID。

	VipSubnetCidrId string `json:"vip_subnet_cidr_id"`
	// 负载均衡器的IPv4虚拟IP地址。

	VipAddress string `json:"vip_address"`
	// 负载均衡器的IPv4对应的port ID。

	VipPortId string `json:"vip_port_id"`
	// 负载均衡的标签列表。

	Tags []Tag `json:"tags"`
	// 负载均衡器的创建时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z'

	CreatedAt string `json:"created_at"`
	// 负载均衡器的更新时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z'

	UpdatedAt string `json:"updated_at"`
	// 是否独享型LB，取值： - false：共享型。 - true：独享型。

	Guaranteed bool `json:"guaranteed"`
	// 负载均衡器所在VPC ID。

	VpcId string `json:"vpc_id"`
	// 负载均衡器绑定的EIP。只支持绑定一个EIP。  注：该字段与publicips一致。

	Eips []EipInfo `json:"eips"`
	// 双栈类型负载均衡器的IPv6地址。  [不支持IPv6，请勿使用。](tag:otc,otc_test,dt,dt_test)

	Ipv6VipAddress string `json:"ipv6_vip_address"`
	// 双栈类型负载均衡器所在子网的IPv6网络ID。  [不支持IPv6，请勿使用。](tag:otc,otc_test,dt,dt_test)

	Ipv6VipVirsubnetId string `json:"ipv6_vip_virsubnet_id"`
	// 双栈类型负载均衡器的IPv6对应的port ID。  [不支持IPv6，请勿使用。](tag:otc,otc_test,dt,dt_test)

	Ipv6VipPortId string `json:"ipv6_vip_port_id"`
	// 负载均衡器所在的可用区列表。

	AvailabilityZoneList []string `json:"availability_zone_list"`
	// 企业项目ID。创建时不传则返回\"0\"，表示资源属于default企业项目。  注：\"0\"并不是真实存在的企业项目ID，在创建、更新和查询时不能作为请求参数传入。  [不支持该字段，请勿使用](tag:otcc,otc_test)

	EnterpriseProjectId string `json:"enterprise_project_id"`
	// 资源账单信息，取值： - 空：按需计费。 - 非空：包周期计费， 包周期计费billing_info字段的格式为：order_id&#58;product_id&#58;region_id&#58;project_id，如： CS2107161019CDJZZ&#58;OFFI569702121789763584&#58;eu-de&#58;057ef081eb00d2732fd1c01a9be75e6f 使用说明： - admin权限才能更新此字段。 [不支持该字段，请勿使用](tag:otc,otc_test,dt,dt_test)

	BillingInfo string `json:"billing_info"`
	// 四层Flavor ID。  [hsco场景下所有LB实例共享带宽，该字段无效，请勿使用。](tag:hc,hws,hcso)

	L4FlavorId string `json:"l4_flavor_id"`
	// 四层弹性Flavor ID。  不支持该字段，请勿使用。

	L4ScaleFlavorId string `json:"l4_scale_flavor_id"`
	// 七层Flavor ID。  [hsco场景下所有LB实例共享带宽，该字段无效，请勿使用。](tag:hc,hws,hcso)

	L7FlavorId string `json:"l7_flavor_id"`
	// 七层弹性Flavor ID。  不支持该字段，请勿使用。

	L7ScaleFlavorId string `json:"l7_scale_flavor_id"`
	// 负载均衡器绑定的公网IP。只支持绑定一个公网IP。  注：该字段与eips一致。

	Publicips []PublicIpInfo `json:"publicips"`
	// 下联面子网网络ID列表。可以通过GET https&#58;//{VPC_Endpoint}/v1/{project_id}/subnets 响应参数中的id得到。 [若不指定该字段，则会在当前负载均衡器所在的VPC中任意选一个子网，优选双栈网络。](tag:hc,hws,hk,ocb,tlf,ctc,hcso,sbc,g42,tm,cmcc,hk-g42) 若指定多个下联面子网，则按顺序优先使用第一个子网来为负载均衡器下联面端口分配ip地址。 下联面子网必须属于该LB所在的VPC。

	ElbVirsubnetIds []string `json:"elb_virsubnet_ids"`
	// 下联面子网类型 - ipv4：ipv4 - dualstack：双栈

	ElbVirsubnetType LoadBalancerElbVirsubnetType `json:"elb_virsubnet_type"`
	// 是否启用跨VPC后端转发。取值： - true：开启、 - false：不开启。  仅独享型负载均衡器支持该特性。  开启跨VPC后端转发后，后端服务器组不仅支持添加云上VPC内的服务器，还支持添加其他VPC、其他公有云、云下数据中心的服务器。  [不支持该字段，请勿使用。](tag:otc,otc_test,dt,dt_test)

	IpTargetEnable bool `json:"ip_target_enable"`
	// 负载均衡器的冻结场景。若负载均衡器有多个冻结场景，用逗号分隔。取值： - POLICE：公安冻结场景。 - ILLEGAL：违规冻结场景。 - VERIFY：客户未实名认证冻结场景。 - RTNER：合作伙伴冻结（合作伙伴冻结子客户资源）。 - REAR：欠费冻结场景。  [不支持该字段，请勿使用。](tag:otc,otc_test,dt,dt_test)

	FrozenScene string `json:"frozen_scene"`

	Ipv6Bandwidth *BandwidthRef `json:"ipv6_bandwidth"`
	// 是否开启删除保护，取值： - false：不开启。 - true：开启。 >退场时需要先关闭所有资源的删除保护开关。  仅当前局点启用删除保护特性后才会返回该字段。

	DeletionProtectionEnable *bool `json:"deletion_protection_enable,omitempty"`

	Autoscaling *AutoscalingRef `json:"autoscaling,omitempty"`
}

func (o LoadBalancer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "LoadBalancer struct{}"
	}

	return strings.Join([]string{"LoadBalancer", string(data)}, " ")
}

type LoadBalancerElbVirsubnetType struct {
	value string
}

type LoadBalancerElbVirsubnetTypeEnum struct {
	IPV4      LoadBalancerElbVirsubnetType
	DUALSTACK LoadBalancerElbVirsubnetType
}

func GetLoadBalancerElbVirsubnetTypeEnum() LoadBalancerElbVirsubnetTypeEnum {
	return LoadBalancerElbVirsubnetTypeEnum{
		IPV4: LoadBalancerElbVirsubnetType{
			value: "ipv4",
		},
		DUALSTACK: LoadBalancerElbVirsubnetType{
			value: "dualstack",
		},
	}
}

func (c LoadBalancerElbVirsubnetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LoadBalancerElbVirsubnetType) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter != nil {
		val, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
		if err == nil {
			c.value = val.(string)
			return nil
		}
		return err
	} else {
		return errors.New("convert enum data to string error")
	}
}
