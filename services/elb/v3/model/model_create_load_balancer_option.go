package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// CreateLoadBalancerOption 参数解释：创建负载均衡器实例的参数对象。
type CreateLoadBalancerOption struct {

	// 参数解释：负载均衡器实例所在的项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 参数解释：负载均衡器的名称。  约束限制： 可以为空, 最大长度为255个字符。  取值范围：支持中文字符、英文字符等unicode字符，且长度为[0-255]个字符。
	Name *string `json:"name,omitempty"`

	// 参数解释：负载均衡器的描述。  约束限制： 可以为空, 最大长度为255个字符。  取值范围：支持中文字符、英文字符等unicode字符，且长度为[0-255]个字符。
	Description *string `json:"description,omitempty"`

	// 参数解释：负载均衡器的IPv4私网IP。该地址必须包含在所在子网的IPv4网段内，且未被占用。  约束限制： - 传入vip_address时，必须传入vip_subnet_cidr_id。 - 不传入vip_address，但传入vip_subnet_cidr_id，则自动分配IPv4私网IP。 - 不传入vip_address，且不传vip_subnet_cidr_id，则不分配IPv4私网IP，vip_address=null。 [- 网关型LB不支持传入vip_address。](tag:hws_eu)  取值范围：满足IPv4的地址格式，[0-255].[0-255].[0-255].[0-255]. 如192.168.1.1。
	VipAddress *string `json:"vip_address,omitempty"`

	// 参数解释：负载均衡器所在子网的IPv4子网ID，也称为该负载均衡器实例的前端子网。  约束限制： - 若创建带有IPv4私网IP的负载均衡实例，则字段必须传入。可以通过调用VPC的API, GET https://{VPC_Endpoint}/v1/{project_id}/subnets 响应参数中的neutron_subnet_id得到。 - vpc_id, vip_subnet_cidr_id, ipv6_vip_virsubnet_id不能同时为空，且需要在同一个vpc下。 - 若同时传入vpc_id和vip_subnet_cidr_id，则vip_subnet_cidr_id对应的子网必须属于vpc_id对应的VPC。 [- 创建网关型LB，vip_subnet_cidr_id和ipv6_vip_virsubnet_id不能同时为空。若都传入则必须是同一个子网的IPv4子网ID和IPv6网络ID。](tag:hws_eu)  取值范围： 标准的UUID格式，长度为36个字符。
	VipSubnetCidrId *string `json:"vip_subnet_cidr_id,omitempty"`

	// 参数解释：双栈类型负载均衡器所在子网的IPv6网络ID，也称为该负载均衡器实例的前端子网。  约束限制： - 若创建带有IPv6 IP的负载均衡实例，则字段必须传入。可以通过GET https://{VPC_Endpoint}/v1/{project_id}/subnets 响应参数中的neutron_network_id得到。 - vpc_id，vip_subnet_cidr_id，ipv6_vip_virsubnet_id不能同时为空，且需要在同一个vpc下。 - 需要对应的子网开启IPv6。 [- 创建网关型LB，vip_subnet_cidr_id和ipv6_vip_virsubnet_id不能同时为空。若都传入则必须是同一个子网的IPv4子网ID和IPv6网络ID。](tag:hws_eu)  取值范围： 标准的UUID格式，长度为36个字符。  [不支持IPv6，请勿使用。](tag:dt,dt_test)
	Ipv6VipVirsubnetId *string `json:"ipv6_vip_virsubnet_id,omitempty"`

	// 参数解释：负载均衡器的生产者名称。固定为vlb，无需指定。
	Provider *string `json:"provider,omitempty"`

	// 参数解释：网络型规格ID。  约束限制： [- 可以通过GET https://{ELB_Endpoint}/v3/{project_id}/elb/flavors?type=L4 响应参数中的id得到。 - 当l4_flavor_id和l7_flavor_id都不传的时，会使用默认flavor（默认flavor根据不同局点有所不同，具体以实际值为准）。 - 当传入的规格类型为L4，表示该实例为固定规格实例，按规格计费。 - 当传入的规格类型为L4_elastic_max，表示该实例为弹性实例，按LCU计费。](tag:hws,hws_hk,hws_eu,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb)   [网关型LB不支持指定l4_flavor_id。](tag:hws_eu) [只支持设置为l4_flavor.elb.shared。](tag:hcso_dt) [所有LB实例共享带宽，该字段无效，请勿使用。](tag:hcso,hk_vdf,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b)  取值范围： 标准的UUID格式，长度为36个字符。
	L4FlavorId *string `json:"l4_flavor_id,omitempty"`

	// 参数解释：应用型规格ID。  约束限制： [- 可以通过GET https://{ELB_Endpoint}/v3/{project_id}/elb/flavors?type=L7 响应参数中的id得到。 - 当l4_flavor_id和l7_flavor_id都不传的时，会使用默认flavor （默认flavor根据不同局点有所不同，具体以实际值为准）。 - 当传入的规格类型为L7，表示该实例为固定规格实例，按规格计费。 - 当传入的规格类型为L7_elastic_max，表示该实例为弹性实例，按LCU计费。](tag:hws,hws_hk,hws_eu,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb)   [网关型LB不支持指定l7_flavor_id。](tag:hws_eu) [只支持设置为l7_flavor.elb.shared。](tag:hcso_dt) [所有LB实例共享带宽，该字段无效，请勿使用。](tag:hcso,hk_vdf,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b)  取值范围： 标准的UUID格式，长度为36个字符。
	L7FlavorId *string `json:"l7_flavor_id,omitempty"`

	// 参数解释：是否为独享型负载均衡器实例。  约束限制：当前只支持设置为true，设置为false会返回400 Bad Request。  取值范围：布尔类型。 - true：独享型。 - false：共享型。  默认取值：true。
	Guaranteed *bool `json:"guaranteed,omitempty"`

	// 参数解释：负载均衡器所在的VPC ID。  约束限制: - 参数获取，可以通过 GET https://{VPC_Endpoint}/v1/{project_id}/vpcs 响应参数中的id得到。 - vpc_id，vip_subnet_cidr_id，ipv6_vip_virsubnet_id不能同时为空，且需要在同一个vpc下。  取值范围： 标准的UUID格式，长度为36个字符。
	VpcId *string `json:"vpc_id,omitempty"`

	// 参数解释：负载均衡器实例所在的可用区列表。  约束限制： - 可通过GET https://{ELB_Endponit}/v3/{project_id}/elb/availability-zones 接口来查询可用区集合列表。创建负载均衡器时，从查询结果选择某一个可用区集合，并从中选择一个或多个可用区。  >为了支持可用区容灾，建议选取不少于2个可用区。
	AvailabilityZoneList []string `json:"availability_zone_list"`

	// 参数解释：负载均衡器所属的企业项目ID。  约束限制：不能传入\"\"、\"0\"或不存在的企业项目ID，创建时不传则资源属于default企业项目，默认返回\"0\"。  [不支持该字段，请勿使用。](tag:dt,dt_test,hcso_dt)
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 参数解释：负载均衡的标签列表。示例：\"tags\":[{\"key\":\"my_tag\",\"value\":\"my_tag_value\"}]
	Tags *[]Tag `json:"tags,omitempty"`

	// 参数解释：负载均衡器的启用状态。  取值范围： - true ：启用。 - false：停用。  默认取值：true。  [不支持该字段，请勿使用。](tag:dt,dt_test)
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 参数解释: 预付费实例的订单信息。  取值范围： - 空：按需计费。 [- 非空：包周期计费。 格式为： order_id:product_id:region_id:project_id，如： CS2107161019CDJZZ:OFFI569702121789763584:region-xxx:057ef081eb00d2732fd1c01a9be75e6f   不支持该字段，请勿使用。](tag:hws_eu,g42,hk_g42,dt,dt_test,hcso_dt,hcso,hk_vdf,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b)
	BillingInfo *string `json:"billing_info,omitempty"`

	Ipv6Bandwidth *BandwidthRef `json:"ipv6_bandwidth,omitempty"`

	// 参数解释：负载均衡器绑定的公网IP的ID的数组。  约束限制： - 只支持绑定数组中的第一个EIP，其他将被忽略。 [- 网关型LB不支持该字段。](tag:hws_eu)
	PublicipIds *[]string `json:"publicip_ids,omitempty"`

	Publicip *CreateLoadBalancerPublicIpOption `json:"publicip,omitempty"`

	// 参数解释：负载均衡器的后端子网的网络ID（OpenStack Neutron接口）列表。负载均衡器实例会预占用该子网中的部分IP， 用于负载均衡器网关与该实例后端服务器通信的源地址（典型场景，健康检查探测的源地址，FULLNAT场景的源地址等）。 使用负载均衡器所在VPC的ID查询可用子网 GET https://{VPC_Endpoint}/v1/{project_id}/subnets?vpc_id=xxxx 获取响应参数中的neutron_network_id字段。  约束限制： - 后端子网必须属于该负载均衡器实例所在的VPC。 - 通常需要用户指定一个特殊的子网，方便用户在后端服务器关联的安全组中，放通该子网的地址段。 - 若指定多个下联面子网，则按顺序优先使用第一个子网来为负载均衡器下联面端口分配ip地址。 - 若不指定该字段，则按如下规则选择下联面网络：   1. 如果ELB实例开启ipv6，则选择ipv6_vip_virsubnet_id子网对应的网络ID。   2. 如果ELB实例没有开启ipv6，开启ipv4，则选择vip_subnet_cidr_id子网对应的网络ID。   3. 如果ELB实例没有选择私网，只开启公网，则会在当前负载均衡器所在的VPC中任意选一个子网，优选可用IP多的网络。  - 后端服务器安全组放通：由于负载均衡器网关会使用该子网中的预占的地址，作为源IP与后端服务器通信（健康检查探测，FULLNAT通信），为避免后端服务器关联的安全组拦截，建议将对应的子网地址段进行安全组放通。 - 预占地址变化：负载均衡实例，弹性扩缩场景，可能涉及到预占地址的变化，建议安全组对子网段放通，而不是具体预占地址的放通。 - 建议后端子网，使用一个独占的地址充足的子网，方便运维管理。
	ElbVirsubnetIds *[]string `json:"elb_virsubnet_ids,omitempty"`

	// 参数解释：是否启用跨VPC后端转发。 开启跨VPC后端转发后，后端服务器组不仅支持添加云上VPC内的服务器，还支持添加其他VPC、其他公有云、云下数据中心的服务器。  约束限制： - 开启后不能关闭。 - 使用共享VPC的实例使用此特性时，需确保共享资源所有者已开通VPC对等连接，否则通信异常。 [- 仅独享型负载均衡器支持该特性。](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,fcs,dt) [- 网关型LB不支持该特性。](tag:hws_eu)  取值范围： - true：开启。 - false：不开启。  [荷兰region不支持该字段，请勿使用。](tag:dt)
	IpTargetEnable *bool `json:"ip_target_enable,omitempty"`

	// 参数解释：实例删除保护开关。  约束限制：实例删除前，需要先关闭该实例下所有资源的删除保护开关。  取值范围：  - false： 不开启。  - true： 开启。  默认取值：false。
	DeletionProtectionEnable *bool `json:"deletion_protection_enable,omitempty"`

	PrepaidOptions *PrepaidCreateOption `json:"prepaid_options,omitempty"`

	Autoscaling *CreateLoadbalancerAutoscalingOption `json:"autoscaling,omitempty"`

	// 参数解释：WAF故障时的流量处理策略。  约束限制：只有绑定了waf的LB实例，该字段才会生效。 [不支持该字段，请勿使用。](tag:hws_hk,hws_eu,hws_test,hcs,hcs_sm,hcso,hk_vdf,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b,hcso_dt,dt,dt_test,ocb,ctc,cmcc,tm,sbc,g42,hws_ocb,hk_sbc,hk_tm,hk_g42)  取值范围： - discard:丢弃。 - forward: 转发到后端。  默认取值：forward。
	WafFailureAction *CreateLoadBalancerOptionWafFailureAction `json:"waf_failure_action,omitempty"`

	// 参数解释：修改保护状态。用于console控制台防误修改。console感知该状态为consoleProtection时，不允许用户直接修改资源其他配置属性。  约束限制：不影响通过API修改资源属性。类似资源标记，用于提升控制台等用户易用性场景，如防误修改。  取值范围： - nonProtection: 不保护。 - consoleProtection: 控制台修改保护。  默认取值：nonProtection。
	ProtectionStatus *CreateLoadBalancerOptionProtectionStatus `json:"protection_status,omitempty"`

	// 参数解释：设置保护的原因。作为protection_status的转态设置的原因。  约束限制：仅当protection_status为consoleProtection时有效。  取值范围：通用Unicode字符集字符，最大255个字符。
	ProtectionReason *string `json:"protection_reason,omitempty"`

	// 参数解释：负载均衡器实例的计费模式。  约束限制：不建议用户传该字段。系统会基于用户传入的l4_flavor_id/l7_flavor_id的规格类型，自动识别计费模式。  取值范围： - flavor：固定规格计费。 - lcu：弹性规格计费（按用户实际使用的lcu个数计费）。  默认取值： - 若是创建共享型实例，不传默认创建固定规格计费实例。 - 若是创建独享型实例，则系统会忽略改字段，而是基于用户传入的l4_flavor_id/l7_flavor_id的规格类型，自动识别计费模式。
	ChargeMode *CreateLoadBalancerOptionChargeMode `json:"charge_mode,omitempty"`

	// 参数解释：负载均衡器实例的IPv6地址。  约束限制：  - 必须属于ipv6_vip_virsubnet_id子网中地址。  - elb_virsubnet_ids中的后端子网必须支持双栈。 [- 网关型LB不支持指定ipv6_vip_address。](tag:hws_eu)  [不支持IPv6，请勿使用。](tag:dt,dt_test)
	Ipv6VipAddress *string `json:"ipv6_vip_address,omitempty"`
}

func (o CreateLoadBalancerOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateLoadBalancerOption struct{}"
	}

	return strings.Join([]string{"CreateLoadBalancerOption", string(data)}, " ")
}

type CreateLoadBalancerOptionWafFailureAction struct {
	value string
}

type CreateLoadBalancerOptionWafFailureActionEnum struct {
	DISCARD CreateLoadBalancerOptionWafFailureAction
	FORWARD CreateLoadBalancerOptionWafFailureAction
}

func GetCreateLoadBalancerOptionWafFailureActionEnum() CreateLoadBalancerOptionWafFailureActionEnum {
	return CreateLoadBalancerOptionWafFailureActionEnum{
		DISCARD: CreateLoadBalancerOptionWafFailureAction{
			value: "discard",
		},
		FORWARD: CreateLoadBalancerOptionWafFailureAction{
			value: "forward",
		},
	}
}

func (c CreateLoadBalancerOptionWafFailureAction) Value() string {
	return c.value
}

func (c CreateLoadBalancerOptionWafFailureAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLoadBalancerOptionWafFailureAction) UnmarshalJSON(b []byte) error {
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

type CreateLoadBalancerOptionProtectionStatus struct {
	value string
}

type CreateLoadBalancerOptionProtectionStatusEnum struct {
	NON_PROTECTION     CreateLoadBalancerOptionProtectionStatus
	CONSOLE_PROTECTION CreateLoadBalancerOptionProtectionStatus
}

func GetCreateLoadBalancerOptionProtectionStatusEnum() CreateLoadBalancerOptionProtectionStatusEnum {
	return CreateLoadBalancerOptionProtectionStatusEnum{
		NON_PROTECTION: CreateLoadBalancerOptionProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: CreateLoadBalancerOptionProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c CreateLoadBalancerOptionProtectionStatus) Value() string {
	return c.value
}

func (c CreateLoadBalancerOptionProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLoadBalancerOptionProtectionStatus) UnmarshalJSON(b []byte) error {
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

type CreateLoadBalancerOptionChargeMode struct {
	value string
}

type CreateLoadBalancerOptionChargeModeEnum struct {
	FLAVOR CreateLoadBalancerOptionChargeMode
	LCU    CreateLoadBalancerOptionChargeMode
}

func GetCreateLoadBalancerOptionChargeModeEnum() CreateLoadBalancerOptionChargeModeEnum {
	return CreateLoadBalancerOptionChargeModeEnum{
		FLAVOR: CreateLoadBalancerOptionChargeMode{
			value: "flavor",
		},
		LCU: CreateLoadBalancerOptionChargeMode{
			value: "lcu",
		},
	}
}

func (c CreateLoadBalancerOptionChargeMode) Value() string {
	return c.value
}

func (c CreateLoadBalancerOptionChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateLoadBalancerOptionChargeMode) UnmarshalJSON(b []byte) error {
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
