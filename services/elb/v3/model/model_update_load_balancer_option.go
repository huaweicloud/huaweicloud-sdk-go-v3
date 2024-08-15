package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// UpdateLoadBalancerOption 更新负载均衡器参数。
type UpdateLoadBalancerOption struct {

	// 参数解释：负载均衡器的名称。
	Name *string `json:"name,omitempty"`

	// 参数解释：负载均衡器的启用状态。  取值范围： - true ：启用。 - false：停用。  [不支持该字段，请勿使用。](tag:dt,dt_test)
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// 参数解释：负载均衡器的描述。
	Description *string `json:"description,omitempty"`

	// 参数解释：双栈类型负载均衡器所在子网的IPv6网络ID，也称为该负载均衡器实例的前端子网。可以通过GET https://{VPC_Endpoint}/v1/{project_id}/subnets 响应参数中的neutron_network_id得到。  通过更新ipv6_vip_virsubnet_id可以更新负载均衡器所在IPv6子网，且负载均衡器的内网IPv6地址将发生变化。  约束限制： - ipv6_vip_virsubnet_id 对应的子网必须属于当前负载均衡器所在VPC，且子网需开启支持IPv6。 - 只有guaranteed是true的负载均衡器才支持更新ipv6_vip_virsubnet_id。 - *传入为null表示解绑IPv6子网。* - 更新IPv6子网不会导致IPv4子网发生变化。  [不支持IPv6，请勿使用。](tag:dt,dt_test)
	Ipv6VipVirsubnetId *string `json:"ipv6_vip_virsubnet_id,omitempty"`

	// 参数解释：负载均衡器所在的IPv4子网ID，也称为该负载均衡器实例的前端子网。可以通过GET https://{VPC_Endpoint}/v1/{project_id}/subnets 响应参数中的neutron_subnet_id得到。  通过更新vip_subnet_cidr_id可以更新负载均衡器所在IPv4子网，并且负载均衡器的内网IPv4地址将发生变化。  约束限制： - 若同时设置了vip_address，则必须保证vip_address对应的IP在vip_subnet_cidr_id的子网网段中。 - vip_subnet_cidr_id对应的子网必须属于当前负载均衡器vpc_id对应的VPC。 - 只有guaranteed是true的负载均衡器才支持更新vip_subnet_cidr_id。 - *传入null表示解绑IPv4子网。* - 更新IPv子网不会导致IPv4子网发生变化。
	VipSubnetCidrId *string `json:"vip_subnet_cidr_id,omitempty"`

	// 参数解释：负载均衡器的IPv4虚拟IP。  约束限制：该地址必须包含在所在子网的IPv4网段内，且未被占用。  注：仅当guaranteed是true的场合，才支持更新。
	VipAddress *string `json:"vip_address,omitempty"`

	// 参数解释：网络型规格ID。  [约束限制： - 可以通过GET https://{ELB_Endpoint}/v3/{project_id}/elb/flavors?type=L4 响应参数中的id得到。 - 仅当guaranteed是true的场合，才支持更新。 - 可以支持规格改大改小，注意改小过程中可能会造成部分长连接中断，影响部分链接的新建， - autoscaling.enable=true时，修改无意义，不生效。 - 当传入的规格类型为L4，表示该实例为固定规格实例，按规格计费。 - 当传入的规格类型为L4_elastic_max，表示该实例为弹性实例，按LCU计费。 ](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,dt)  [只支持设置为l4_flavor.elb.shared。](tag:hcso_dt)  [hcso场景下所有LB实例共享带宽，该字段无效，请勿使用。](tag:hcso,hk_vdf,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b)
	L4FlavorId *string `json:"l4_flavor_id,omitempty"`

	// 参数解释：应用型ID。  [约束限制： - 可以通过GET https://{ELB_Endpoint}/v3/{project_id}/elb/flavors?type=L7 响应参数中的id得到。 - 仅当guaranteed是true的场合，才支持更新。 - 可以支持规格改大改小，注意改小过程中可能会造成部分长连接中断，影响部分链接的新建， - autoscaling.enable=true时，修改无意义，不生效。 - 当传入的规格类型为L7，表示该实例为固定规格实例，按规格计费。 - 当传入的规格类型为L7_elastic_max，表示该实例为弹性实例，按LCU计费。 ](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,dt) [网关型LB不支持l7_flavor_id。](tag:hws_eu) [只支持设置为l7_flavor.elb.shared。](tag:hcso_dt)   [所有LB实例共享带宽，该字段无效，请勿使用。](tag:hcso,hk_vdf,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b)
	L7FlavorId *string `json:"l7_flavor_id,omitempty"`

	Ipv6Bandwidth *BandwidthRef `json:"ipv6_bandwidth,omitempty"`

	// 参数解释：是否启用跨VPC后端转发。 开启跨VPC后端转发后，后端服务器组不仅支持添加云上VPC内的服务器，还支持添加其他VPC、其他公有云、云下数据中心的服务器。  约束限制： - 使用共享VPC的实例使用此特性时，需确保共享资源所有者已开通VPC对等连接，否则通信异常。 - 开启后不能关闭。 [- 仅独享型负载均衡器支持该特性。 ](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,fcs,dt)  取值范围： - true：开启。 - false：不开启。  [荷兰region不支持该字段，请勿使用。](tag:dt,dt_test)
	IpTargetEnable *bool `json:"ip_target_enable,omitempty"`

	// 参数解释：下联面子网的网络ID列表。 可以通过GET https://{VPC_Endpoint}/v1/{project_id}/subnets 响应参数中的neutron_network_id得到。  约束限制： - 已绑定的下联面子网也在传参elb_virsubnet_ids列表中，则绑定关系保留。 - 已绑定的下联面子网若不在传参elb_virsubnet_ids列表中， 则将移除LB与该下联面子网的关联关系。但不允许移除已被ELB使用的子网，否则将报错，不做任何修改。 - 在传参elb_virsubnet_ids列表中但不在已绑定的下联面子网列表中，则将新增LB与下联面的绑定关系。 - 所有elb_virsubnet_ids中的ID同属于该LB所在的VPC。 - 不支持边缘云子网。
	ElbVirsubnetIds *[]string `json:"elb_virsubnet_ids,omitempty"`

	// 参数解释：是否开启删除保护。  约束限制：退场时需要先关闭所有资源的删除保护开关。  取值范围：false不开启，true开启。  [不支持该字段，请勿使用。](tag:hws_eu,g42,hk_g42)  [荷兰region不支持该字段，请勿使用。](tag:dt,dt_test)
	DeletionProtectionEnable *bool `json:"deletion_protection_enable,omitempty"`

	PrepaidOptions *PrepaidUpdateOption `json:"prepaid_options,omitempty"`

	Autoscaling *UpdateLoadbalancerAutoscalingOption `json:"autoscaling,omitempty"`

	// 参数解释：计费模式。  取值范围： - flavor：按规格计费
	ChargeMode *string `json:"charge_mode,omitempty"`

	// 参数解释：WAF故障时的流量处理策略。  约束限制：只有绑定了waf的LB实例，该字段才会生效。  取值范围：discard:丢弃，forward: 转发到后端。  默认取值：forward  [不支持该字段，请勿使用。](tag:hws_hk,hws_eu,hws_test,hcs,hcs_sm,hcso,hk_vdf,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b,hcso_dt,dt,dt_test,ocb,ctc,cmcc,tm,sbc,g42,hws_ocb,hk_sbc,hk_tm,hk_g42)
	WafFailureAction *UpdateLoadBalancerOptionWafFailureAction `json:"waf_failure_action,omitempty"`

	// 参数解释：修改保护状态。  取值范围： - nonProtection: 不保护 - consoleProtection: 控制台修改保护
	ProtectionStatus *UpdateLoadBalancerOptionProtectionStatus `json:"protection_status,omitempty"`

	// 参数解释：设置保护的原因。  约束限制：仅当protection_status为consoleProtection时有效。
	ProtectionReason *string `json:"protection_reason,omitempty"`

	// 参数解释：双栈类型负载均衡器的IPv6地址。  [不支持IPv6，请勿使用。](tag:dt,dt_test)
	Ipv6VipAddress *string `json:"ipv6_vip_address,omitempty"`
}

func (o UpdateLoadBalancerOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "UpdateLoadBalancerOption struct{}"
	}

	return strings.Join([]string{"UpdateLoadBalancerOption", string(data)}, " ")
}

type UpdateLoadBalancerOptionWafFailureAction struct {
	value string
}

type UpdateLoadBalancerOptionWafFailureActionEnum struct {
	DISCARD UpdateLoadBalancerOptionWafFailureAction
	FORWARD UpdateLoadBalancerOptionWafFailureAction
}

func GetUpdateLoadBalancerOptionWafFailureActionEnum() UpdateLoadBalancerOptionWafFailureActionEnum {
	return UpdateLoadBalancerOptionWafFailureActionEnum{
		DISCARD: UpdateLoadBalancerOptionWafFailureAction{
			value: "discard",
		},
		FORWARD: UpdateLoadBalancerOptionWafFailureAction{
			value: "forward",
		},
	}
}

func (c UpdateLoadBalancerOptionWafFailureAction) Value() string {
	return c.value
}

func (c UpdateLoadBalancerOptionWafFailureAction) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateLoadBalancerOptionWafFailureAction) UnmarshalJSON(b []byte) error {
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

type UpdateLoadBalancerOptionProtectionStatus struct {
	value string
}

type UpdateLoadBalancerOptionProtectionStatusEnum struct {
	NON_PROTECTION     UpdateLoadBalancerOptionProtectionStatus
	CONSOLE_PROTECTION UpdateLoadBalancerOptionProtectionStatus
}

func GetUpdateLoadBalancerOptionProtectionStatusEnum() UpdateLoadBalancerOptionProtectionStatusEnum {
	return UpdateLoadBalancerOptionProtectionStatusEnum{
		NON_PROTECTION: UpdateLoadBalancerOptionProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: UpdateLoadBalancerOptionProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c UpdateLoadBalancerOptionProtectionStatus) Value() string {
	return c.value
}

func (c UpdateLoadBalancerOptionProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *UpdateLoadBalancerOptionProtectionStatus) UnmarshalJSON(b []byte) error {
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
