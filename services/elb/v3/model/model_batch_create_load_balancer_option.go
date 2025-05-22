package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BatchCreateLoadBalancerOption 创建负载均衡器参数。
type BatchCreateLoadBalancerOption struct {

	// 负载均衡器ID（UUID）。不支持该字段，请勿使用。
	Id *string `json:"id,omitempty"`

	// 批创负载均衡器ID的列表（UUID）。不支持该字段，请勿使用。
	Ids *[]string `json:"ids,omitempty"`

	// 负载均衡器所在的项目ID。
	ProjectId *string `json:"project_id,omitempty"`

	// 批量创建负载均衡器的数量
	Number *int32 `json:"number,omitempty"`

	// 负载均衡器的名称。 使用说明： - 当批创负载均衡器数量为1时，负载均衡器的名字即为该传入的值。 - 当批创负载均衡器数量大于1时，负载均衡器的名字在传入的name后加上4位后缀。  例如：批量创建3个负载均衡器，传入name为\"elb-test\"，则三个负载均衡器名字分别为： \"elb-test-0001\"，\"elb-test-0002\"，\"elb-test-0003\"
	Name *string `json:"name,omitempty"`

	// 负载均衡器的描述。
	Description *string `json:"description,omitempty"`

	// 负载均衡器的IPv4虚拟IP。该地址必须包含在所在子网的IPv4网段内，且未被占用。  使用说明： - 传入vip_address时必须传入vip_subnet_cidr_id。 - 不传入vip_address，但传入vip_subnet_cidr_id，则自动分配IPv4虚拟IP。 - 不传入vip_address，且不传vip_subnet_cidr_id，则不分配虚拟IP，vip_address=null。 - 当批量创建的number数量大于1时不支持传入该字段，返回400。
	VipAddress *string `json:"vip_address,omitempty"`

	// 负载均衡器所在子网的IPv4子网ID。若需要创建带IPv4虚拟IP的LB，该字段必须传入。 可以通过GET https://{VPC_Endpoint}/v1/{project_id}/subnets 响应参数中的neutron_subnet_id得到。  使用说明： - vpc_id, vip_subnet_cidr_id, ipv6_vip_virsubnet_id不能同时为空，且需要在同一个vpc下。 - 若同时传入vpc_id和vip_subnet_cidr_id， 则vip_subnet_cidr_id对应的子网必须属于vpc_id对应的VPC。
	VipSubnetCidrId *string `json:"vip_subnet_cidr_id,omitempty"`

	// 双栈类型负载均衡器所在子网的IPv6网络ID。可以通过GET https://{VPC_Endpoint}/v1/{project_id}/subnets 响应参数中的neutron_network_id得到。  使用说明： - vpc_id，vip_subnet_cidr_id，ipv6_vip_virsubnet_id不能同时为空，且需要在同一个vpc下。 - 需要对应的子网开启IPv6。 - 批量创建共享型实例时，该字段不允许指定。  [不支持IPv6，请勿使用。](tag:dt)
	Ipv6VipVirsubnetId *string `json:"ipv6_vip_virsubnet_id,omitempty"`

	// 负载均衡器的生产者名称。固定为vlb。批量创建共享型实例时，该字段不允许指定。
	Provider *string `json:"provider,omitempty"`

	// 网络型规格ID。  [使用说明： - 可以通过GET https://{ELB_Endpoint}/v3/{project_id}/elb/flavors?type=L4 响应参数中的id得到。 - 当l4_flavor_id和l7_flavor_id都不传的时，会使用默认flavor （默认flavor根据不同局点有所不同，具体以实际值为准）。 - 当传入的规格类型为L4，表示该实例为固定规格实例，按规格计费。 - 当传入的规格类型为L4_elastic_max，表示该实例为弹性实例，按LCU计费。 - 批量创建共享型实例时，该字段不允许指定。  ](tag:hws,hws_hk,hws_eu,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb)   [网关型LB不支持指定l4_flavor_id。](tag:hws_eu) [只支持设置为l4_flavor.elb.shared。](tag:hcso_dt) [所有LB实例共享带宽，该字段无效，请勿使用。](tag:hcso,srg,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b)
	L4FlavorId *string `json:"l4_flavor_id,omitempty"`

	// 应用型规格ID。  [使用说明： - 可以通过GET https://{ELB_Endpoint}/v3/{project_id}/elb/flavors?type=L7 响应参数中的id得到。 - 当l4_flavor_id和l7_flavor_id都不传的时，会使用默认flavor （默认flavor根据不同局点有所不同，具体以实际值为准）。 - 当传入的规格类型为L7，表示该实例为固定规格实例，按规格计费。 - 当传入的规格类型为L7_elastic_max，表示该实例为弹性实例，按LCU计费。 - 批量创建共享型实例时，该字段不允许指定。 ](tag:hws,hws_hk,hws_eu,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb)   [网关型LB不支持指定l7_flavor_id。](tag:hws_eu) [只支持设置为l7_flavor.elb.shared。](tag:hcso_dt) [所有LB实例共享带宽，该字段无效，请勿使用。](tag:hcso,srg,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b)
	L7FlavorId *string `json:"l7_flavor_id,omitempty"`

	// 批量创建的是否是独享型负载均衡器。  取值： - true：独享型。 - false：共享型。  默认：true。
	Guaranteed *bool `json:"guaranteed,omitempty"`

	// 负载均衡器所在的VPC ID。可以通过GET https://{VPC_Endpoint}/v1/{project_id}/vpcs 响应参数中的id得到。  使用说明： - vpc_id，vip_subnet_cidr_id，ipv6_vip_virsubnet_id不能同时为空，且需要在同一个vpc下。 - 批量创建共享型实例时，该字段不允许指定。
	VpcId *string `json:"vpc_id,omitempty"`

	// 可用区列表。可通过GET https://{ELB_Endpoint}/v3/{project_id}/elb/availability-zones 接口来查询可用区集合列表。创建负载均衡器时，从查询结果选择某一个可用区集合，并从中选择一个或多个可用区。  使用说明： - 批量创建共享型实例时，该字段不允许指定。
	AvailabilityZoneList []string `json:"availability_zone_list"`

	// 负载均衡器所属的企业项目ID。创建时传入\"\"、\"0\"或不存在的企业项目ID，或者不传则资源属于default企业项目，默认返回\"0\"。  [不支持该字段，请勿使用。](tag:dt,hcso_dt)
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 负载均衡的标签列表。示例：\"tags\":[{\"key\":\"my_tag\",\"value\":\"my_tag_value\"}]
	Tags *[]Tag `json:"tags,omitempty"`

	// 参数解释：负载均衡器的启用状态。  取值范围： - true ：启用。 - false：停用。  默认取值：true。  [不支持该字段，请勿使用。](tag:dt)
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 资源账单信息。  取值范围： - 空：按需计费。 [- 非空：包周期计费，格式为：order_id:product_id:region_id:project_id。如：CS2107161019CDJZZ:OFFI569702121789763584:az1:057ef081eb00d2732fd1c01a9be75e6f](tag:hws)  [不支持该字段，请勿使用。](tag:hws_hk,hws_eu,hws_eu_wb,hws_test,srg,fcs,fcs_vm,dt,ctc,cmcc,tm,sbc,hk_sbc,hk_tm,hk_vdf)
	BillingInfo *string `json:"billing_info,omitempty"`

	Ipv6Bandwidth *BandwidthRef `json:"ipv6_bandwidth,omitempty"`

	Bandwidth *BandwidthRef `json:"bandwidth,omitempty"`

	// 负载均衡器绑定的公网IP ID。只支持绑定数组中的第一个EIP，其他将被忽略。 当批量创建的number数量大于1时不支持传入该字段，返回400。
	PublicipIds *[]string `json:"publicip_ids,omitempty"`

	Publicip *CreateLoadBalancerPublicIpOption `json:"publicip,omitempty"`

	// 下联面子网的网络ID列表。可以通过GET https://{VPC_Endpoint}/v1/{project_id}/subnets 响应参数中的neutron_network_id得到。  若不指定该字段，则按如下规则选择下联面网络： 1. 如果ELB实例开启ipv6，则选择ipv6_vip_virsubnet_id子网对应的网络ID。 2. 如果ELB实例没有开启ipv6，开启ipv4，则选择vip_subnet_cidr_id子网对应的网络ID。 3. 如果ELB实例没有选择私网，只开启公网，则会在当前负载均衡器所在的VPC中任意选一个子网，优选可用IP多的网络。 4. 批量创建共享型实例的场景，该字段不允许指定。  若指定多个下联面子网，则按顺序优先使用第一个子网来为负载均衡器下联面端口分配ip地址。  下联面子网必须属于该LB所在的VPC。
	ElbVirsubnetIds *[]string `json:"elb_virsubnet_ids,omitempty"`

	// 参数解释：是否启用跨VPC后端转发。 开启跨VPC后端转发后，后端服务器组不仅支持添加云上VPC内的服务器，还支持添加其他VPC、其他公有云、云下数据中心的服务器。  约束限制： - 开启后不能关闭。 - 使用共享VPC的实例使用此特性时，需确保共享资源所有者已开通VPC对等连接，否则通信异常。 [- 仅独享型负载均衡器支持该特性。](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,srg,fcs,dt) [- 网关型LB不支持该特性。](tag:hws_eu)  取值范围： - true：开启。 - false：不开启。  [荷兰region不支持该字段，请勿使用。](tag:dt)
	IpTargetEnable *bool `json:"ip_target_enable,omitempty"`

	// 是否开启删除保护。 批量创建共享型实例的场景，该字段不允许指定。  取值：false不开启，true开启。默认false不开启。  > 退场时需要先关闭所有资源的删除保护开关。  [不支持该字段，请勿使用。](tag:hws_eu,g42,hk_g42)  [荷兰region不支持该字段，请勿使用。](tag:dt)
	DeletionProtectionEnable *bool `json:"deletion_protection_enable,omitempty"`

	PrepaidOptions *PrepaidCreateOption `json:"prepaid_options,omitempty"`

	Autoscaling *CreateLoadbalancerAutoscalingOption `json:"autoscaling,omitempty"`

	// WAF故障时的流量处理策略。discard:丢弃，forward: 转发到后端（默认）  使用说明：只有绑定了waf的LB实例，该字段才会生效。  [不支持该字段，请勿使用。](tag:hws_eu,hws_test,hcs,hcs_sm,hcso,srg,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b,hcso_dt,dt,ocb,ctc,cmcc,tm,ct,sbc,g42,hws_ocb,hk_sbc,hk_tm,hk_g42)
	WafFailureAction *BatchCreateLoadBalancerOptionWafFailureAction `json:"waf_failure_action,omitempty"`

	// 修改保护状态, 取值： - nonProtection: 不保护，默认值为nonProtection - consoleProtection: 控制台修改保护
	ProtectionStatus *BatchCreateLoadBalancerOptionProtectionStatus `json:"protection_status,omitempty"`

	// 参数解释：设置保护的原因。作为protection_status的转态设置的原因。  约束限制：仅当protection_status为consoleProtection时有效。  取值范围：除<和>外通用Unicode字符集字符，最大255个字符。
	ProtectionReason *string `json:"protection_reason,omitempty"`

	// 创建实例采用的计费模式。flavor: 按实例规格收费，默认值。lcu: 按使用量收费。 批量创建共享型实例时，该字段不允许指定。
	ChargeMode *BatchCreateLoadBalancerOptionChargeMode `json:"charge_mode,omitempty"`

	// 双栈类型负载均衡器的IPv6地址。 当批量创建的number数量大于1时不支持传入该字段，返回400。 批量创建共享型实例的场景，该字段不允许指定。 [不支持IPv6，请勿使用。](tag:dt)
	Ipv6VipAddress *string `json:"ipv6_vip_address,omitempty"`
}

func (o BatchCreateLoadBalancerOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BatchCreateLoadBalancerOption struct{}"
	}

	return strings.Join([]string{"BatchCreateLoadBalancerOption", string(data)}, " ")
}

type BatchCreateLoadBalancerOptionWafFailureAction struct {
	value string
}

type BatchCreateLoadBalancerOptionWafFailureActionEnum struct {
	DISCARD BatchCreateLoadBalancerOptionWafFailureAction
	FORWARD BatchCreateLoadBalancerOptionWafFailureAction
}

func GetBatchCreateLoadBalancerOptionWafFailureActionEnum() BatchCreateLoadBalancerOptionWafFailureActionEnum {
	return BatchCreateLoadBalancerOptionWafFailureActionEnum{
		DISCARD: BatchCreateLoadBalancerOptionWafFailureAction{
			value: "discard",
		},
		FORWARD: BatchCreateLoadBalancerOptionWafFailureAction{
			value: "forward",
		},
	}
}

func (c BatchCreateLoadBalancerOptionWafFailureAction) Value() string {
	return c.value
}

func (c BatchCreateLoadBalancerOptionWafFailureAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateLoadBalancerOptionWafFailureAction) UnmarshalJSON(b []byte) error {
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

type BatchCreateLoadBalancerOptionProtectionStatus struct {
	value string
}

type BatchCreateLoadBalancerOptionProtectionStatusEnum struct {
	NON_PROTECTION     BatchCreateLoadBalancerOptionProtectionStatus
	CONSOLE_PROTECTION BatchCreateLoadBalancerOptionProtectionStatus
}

func GetBatchCreateLoadBalancerOptionProtectionStatusEnum() BatchCreateLoadBalancerOptionProtectionStatusEnum {
	return BatchCreateLoadBalancerOptionProtectionStatusEnum{
		NON_PROTECTION: BatchCreateLoadBalancerOptionProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: BatchCreateLoadBalancerOptionProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c BatchCreateLoadBalancerOptionProtectionStatus) Value() string {
	return c.value
}

func (c BatchCreateLoadBalancerOptionProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateLoadBalancerOptionProtectionStatus) UnmarshalJSON(b []byte) error {
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

type BatchCreateLoadBalancerOptionChargeMode struct {
	value string
}

type BatchCreateLoadBalancerOptionChargeModeEnum struct {
	FLAVOR BatchCreateLoadBalancerOptionChargeMode
	LCU    BatchCreateLoadBalancerOptionChargeMode
}

func GetBatchCreateLoadBalancerOptionChargeModeEnum() BatchCreateLoadBalancerOptionChargeModeEnum {
	return BatchCreateLoadBalancerOptionChargeModeEnum{
		FLAVOR: BatchCreateLoadBalancerOptionChargeMode{
			value: "flavor",
		},
		LCU: BatchCreateLoadBalancerOptionChargeMode{
			value: "lcu",
		},
	}
}

func (c BatchCreateLoadBalancerOptionChargeMode) Value() string {
	return c.value
}

func (c BatchCreateLoadBalancerOptionChargeMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BatchCreateLoadBalancerOptionChargeMode) UnmarshalJSON(b []byte) error {
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
