package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// LoadBalancer 负载均衡器的详细信息。
type LoadBalancer struct {

	// 参数解释：负载均衡器ID。
	Id string `json:"id"`

	// 参数解释：负载均衡器描述信息。
	Description string `json:"description"`

	// 参数解释：负载均衡器的配置状态。  取值范围： - ACTIVE：使用中。 - PENDING_DELETE：删除中。
	ProvisioningStatus string `json:"provisioning_status"`

	// 参数解释：负载均衡器的启用状态。  取值范围： - true ：启用。 - false：停用。  [不支持该字段，请勿使用。](tag:dt,dt_test)
	AdminStateUp bool `json:"admin_state_up"`

	// 参数解释：负载均衡器的生产者名称。固定为vlb。
	Provider string `json:"provider"`

	// 参数解释：负载均衡器直接关联的后端云服务器组的ID列表。
	Pools []PoolRef `json:"pools"`

	// 参数解释：负载均衡器关联的监听器的ID列表。
	Listeners []ListenerRef `json:"listeners"`

	// 参数解释：负载均衡器的操作状态。  取值范围： - ONLINE：在线。 - FROZEN：已冻结。
	OperatingStatus string `json:"operating_status"`

	// 参数解释：负载均衡器的名称。
	Name string `json:"name"`

	// 参数解释：负载均衡器所属的项目ID。
	ProjectId string `json:"project_id"`

	// 参数解释：负载均衡器所在子网的IPv4子网ID，也称为该负载均衡器实例的前端子网。
	VipSubnetCidrId string `json:"vip_subnet_cidr_id"`

	// 参数解释：负载均衡器的IPv4虚拟IP地址。
	VipAddress string `json:"vip_address"`

	// 参数解释：负载均衡器的IPv4对应的port ID。 [创建弹性负载均衡时，会自动为负载均衡创建一个port并关联一个默认的安全组，这个安全组对所有流量不生效。 ](tag:dt,dt_test,hcso_dt)
	VipPortId string `json:"vip_port_id"`

	// 参数解释：负载均衡的标签列表。
	Tags []Tag `json:"tags"`

	// 参数解释：负载均衡器的创建时间。  取值范围： 格式：yyyy-MM-dd'T'HH:mm:ss'Z'
	CreatedAt string `json:"created_at"`

	// 参数解释：负载均衡器的更新时间。  取值范围; 格式：yyyy-MM-dd'T'HH:mm:ss'Z'
	UpdatedAt string `json:"updated_at"`

	// 参数解释：是否独享型LB。  取值范围： - false：共享型。 - true：独享型。
	Guaranteed bool `json:"guaranteed"`

	// 参数解释：负载均衡器所在VPC ID。
	VpcId string `json:"vpc_id"`

	// 参数解释：负载均衡器绑定的EIP。  约束限制：只支持绑定一个EIP。  注：该字段与publicips一致。
	Eips []EipInfo `json:"eips"`

	// 参数解释：双栈类型负载均衡器的IPv6地址。  [不支持IPv6，请勿使用。](tag:dt,dt_test)
	Ipv6VipAddress string `json:"ipv6_vip_address"`

	// 参数解释：双栈类型负载均衡器所在子网的IPv6网络ID，也称为该负载均衡器实例的前端子网。  [不支持IPv6，请勿使用。](tag:dt,dt_test)
	Ipv6VipVirsubnetId string `json:"ipv6_vip_virsubnet_id"`

	// 参数解释：双栈类型负载均衡器的IPv6对应的port ID。  [不支持IPv6，请勿使用。](tag:dt,dt_test)
	Ipv6VipPortId string `json:"ipv6_vip_port_id"`

	// 参数解释：负载均衡器所在的可用区列表。
	AvailabilityZoneList []string `json:"availability_zone_list"`

	// 参数解释：企业项目ID。创建时不传则返回\"0\"，表示资源属于default企业项目。  注意：\"0\"并不是真实存在的企业项目ID，在创建、更新和查询时不能作为请求参数传入。  [不支持该字段，请勿使用](tag:dt,dt_test,hcso_dt)
	EnterpriseProjectId string `json:"enterprise_project_id"`

	// 参数解释：资源账单信息。  取值范围： - 空：按需计费。 [- 非空：包周期计费，  包周期计费billing_info字段的格式为：order_id:product_id:region_id:project_id，如：  CS2107161019CDJZZ:OFFI569702121789763584: az:057ef081eb00d2732fd1c01a9be75e6f  不支持该字段，请勿使用](tag:hws_eu,g42,hk_g42,dt,dt_test,hcso_dt,hcso,hk_vdf,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b)
	BillingInfo string `json:"billing_info"`

	// 参数解释：网络型规格ID。  对于弹性扩缩容实例，表示上限规格。  约束限制： - 当传入的规格类型为L4，表示该实例为固定规格实例，按规格计费； - 当传入的规格类型为L4_elastic_max，表示该实例为弹性实例，按LCU计费  [hsco场景下所有LB实例共享带宽，该字段无效，请勿使用。](tag:hcso,hk_vdf,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b)
	L4FlavorId string `json:"l4_flavor_id"`

	// 参数解释：四层弹性Flavor ID。  不支持该字段，请勿使用。
	L4ScaleFlavorId string `json:"l4_scale_flavor_id"`

	// 参数解释：应用型ID。 对于弹性扩缩容实例，表示上限规格ID。  约束限制： - 当传入的规格类型为L7，表示该实例为固定规格实例，按规格计费； - 当传入的规格类型为L7_elastic_max，表示该实例为弹性实例，按LCU计费  [hsco场景下所有LB实例共享带宽，该字段无效，请勿使用。](tag:hcso,hk_vdf,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b)
	L7FlavorId string `json:"l7_flavor_id"`

	// 参数解释：七层弹性Flavor ID。  不支持该字段，请勿使用。
	L7ScaleFlavorId string `json:"l7_scale_flavor_id"`

	// 参数解释：负载均衡器绑定的公网IP。只支持绑定一个公网IP。  注：该字段与eips一致。
	Publicips []PublicIpInfo `json:"publicips"`

	// 参数解释：负载均衡器绑定的global eip。  约束限制：只支持绑定一个globaleip。  [不支持该字段，请勿使用。](tag:hws_eu,g42,hk_g42,dt,dt_test,hcso_dt,hk_vdf,fcs,ctc,ocb,hws_ocb)
	GlobalEips []GlobalEipInfo `json:"global_eips"`

	// 参数解释：下联面子网的网络ID列表。
	ElbVirsubnetIds []string `json:"elb_virsubnet_ids"`

	// 参数解释：下联面子网类型。  取值范围： - ipv4：ipv4 - dualstack：双栈
	ElbVirsubnetType LoadBalancerElbVirsubnetType `json:"elb_virsubnet_type"`

	// 参数解释：是否启用跨VPC后端转发。 开启跨VPC后端转发后，后端服务器组不仅支持添加云上VPC内的服务器，还支持添加其他VPC、其他公有云、云下数据中心的服务器。  约束限制： - 开启后不能关闭。 - 使用共享VPC的实例使用此特性时，需确保共享资源所有者已开通VPC对等连接，否则通信异常。 [- 仅独享型负载均衡器支持该特性。 ](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,hk_vdf,fcs,dt)  取值范围： - true：开启。 - false：不开启。  [荷兰region不支持该字段，请勿使用。](tag:dt,dt_test)
	IpTargetEnable bool `json:"ip_target_enable"`

	// 参数解释：负载均衡器的冻结场景。 [若负载均衡器有多个冻结场景，用逗号分隔。  取值范围： - POLICE：公安冻结场景。 - ILLEGAL：违规冻结场景。 - VERIFY：客户未实名认证冻结场景。 - PARTNER：合作伙伴冻结（合作伙伴冻结子客户资源）。 - AREAR：欠费冻结场景。](tag:hws,hws_hk)  [不支持该字段，请勿使用。](tag:hws_eu,g42,hk_g42,dt,dt_test,hcso_dt,ocb,hws_ocb)
	FrozenScene string `json:"frozen_scene"`

	Ipv6Bandwidth *BandwidthRef `json:"ipv6_bandwidth"`

	// 参数解释：是否开启删除保护。  约束限制： - 仅当前局点启用删除保护特性后才会返回该字段。 - 退场时需要先关闭所有资源的删除保护开关。  取值范围： - false：不开启。 - true：开启。 >退场时需要先关闭所有资源的删除保护开关。  [不支持该字段，请勿使用。](tag:hws_eu,g42,hk_g42)  [荷兰region不支持该字段，请勿使用。](tag:dt,dt_test)
	DeletionProtectionEnable *bool `json:"deletion_protection_enable,omitempty"`

	Autoscaling *AutoscalingRef `json:"autoscaling,omitempty"`

	// 参数解释：LB所属AZ组。
	PublicBorderGroup *string `json:"public_border_group,omitempty"`

	// 参数解释：负载均衡器实例的计费模式。  取值范围： - flavor：按规格计费 - lcu：按使用量计费 - 空值：若是共享型表示免费实例。若是独享型则与flavor模式一致，都是按规格计费。
	ChargeMode *string `json:"charge_mode,omitempty"`

	// 参数解释：WAF故障时的流量处理策略。  约束限制：只有绑定了waf的LB实例，该字段才会生效。  取值范围：discard:丢弃，forward: 转发到后端。  默认取值：forward  [不支持该字段，请勿使用。](tag:hws_hk,hws_eu,hws_test,hcs,hcs_sm,hcso,hk_vdf,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b,hcso_dt,dt,dt_test,ocb,ctc,cmcc,tm,sbc,g42,hws_ocb,hk_sbc,hk_tm,hk_g42)
	WafFailureAction *string `json:"waf_failure_action,omitempty"`

	// 参数解释：修改保护状态。  取值范围： - nonProtection: 不保护。 - consoleProtection: 控制台修改保护。  默认取值：nonProtection
	ProtectionStatus *LoadBalancerProtectionStatus `json:"protection_status,omitempty"`

	// 参数解释：设置保护的原因   约束限制：仅当protection_status为consoleProtection时有效。
	ProtectionReason *string `json:"protection_reason,omitempty"`

	// 参数解释：LB所绑定的logtank对应的group id
	LogGroupId *string `json:"log_group_id,omitempty"`

	// 参数解释：LB所绑定的logtank对应的topic id
	LogTopicId *string `json:"log_topic_id,omitempty"`
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

func (c LoadBalancerElbVirsubnetType) Value() string {
	return c.value
}

func (c LoadBalancerElbVirsubnetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LoadBalancerElbVirsubnetType) UnmarshalJSON(b []byte) error {
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

type LoadBalancerProtectionStatus struct {
	value string
}

type LoadBalancerProtectionStatusEnum struct {
	NON_PROTECTION     LoadBalancerProtectionStatus
	CONSOLE_PROTECTION LoadBalancerProtectionStatus
}

func GetLoadBalancerProtectionStatusEnum() LoadBalancerProtectionStatusEnum {
	return LoadBalancerProtectionStatusEnum{
		NON_PROTECTION: LoadBalancerProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: LoadBalancerProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c LoadBalancerProtectionStatus) Value() string {
	return c.value
}

func (c LoadBalancerProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *LoadBalancerProtectionStatus) UnmarshalJSON(b []byte) error {
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
