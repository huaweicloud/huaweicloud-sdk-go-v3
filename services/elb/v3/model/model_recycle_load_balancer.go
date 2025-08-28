package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// RecycleLoadBalancer 负载均衡器的详细信息。
type RecycleLoadBalancer struct {

	// **参数解释**：负载均衡器ID。  **取值范围**：不涉及
	Id *string `json:"id,omitempty"`

	// **参数解释**：回收站elb的自动到期销毁时间。格式：yyyy-MM-dd'T'HH:mm:ss'Z'  **取值范围**：不涉及
	AutoTerminateTime *string `json:"auto_terminate_time,omitempty"`

	// **参数解释**：负载均衡器描述信息。  **取值范围**：不涉及
	Description *string `json:"description,omitempty"`

	// **参数解释**：负载均衡器的配置状态。  **取值范围**： - RECYCLING：处于回收站用中。
	ProvisioningStatus *string `json:"provisioning_status,omitempty"`

	// **参数解释**：负载均衡器的启用状态。  **取值范围**： - true ：启用。 - false：停用。  [不支持该字段，请勿使用。](tag:dt)
	AdminStateUp *bool `json:"admin_state_up,omitempty"`

	// **参数解释**：负载均衡器的生产者名称。固定为vlb。  **取值范围**：不涉及
	Provider *string `json:"provider,omitempty"`

	// **参数解释**：负载均衡器直接关联的后端服务器组的ID列表。
	Pools *[]PoolRef `json:"pools,omitempty"`

	// **参数解释**：负载均衡器关联的监听器的ID列表。
	Listeners *[]ListenerRef `json:"listeners,omitempty"`

	// **参数解释**：负载均衡器的操作状态。  **取值范围**： - ONLINE：在线。 - FROZEN：已冻结。
	OperatingStatus *string `json:"operating_status,omitempty"`

	// **参数解释**：负载均衡器的名称。  **取值范围**：不涉及
	Name *string `json:"name,omitempty"`

	// **参数解释**：负载均衡器所属的项目ID。  **取值范围**：不涉及
	ProjectId *string `json:"project_id,omitempty"`

	// **参数解释**：负载均衡器所在子网的IPv4子网ID，也称为该负载均衡器实例的前端子网。  **取值范围**：不涉及
	VipSubnetCidrId *string `json:"vip_subnet_cidr_id,omitempty"`

	// **参数解释**：负载均衡器的IPv4私网IP地址。  **取值范围**：不涉及
	VipAddress *string `json:"vip_address,omitempty"`

	// **参数解释**：负载均衡器的IPv4对应的port ID。 [创建弹性负载均衡时，会自动为负载均衡创建一个port并关联一个默认的安全组，这个安全组对所有流量不生效。 ](tag:dt,hcso_dt)  **取值范围**：不涉及
	VipPortId *string `json:"vip_port_id,omitempty"`

	// **参数解释**：负载均衡的标签列表。
	Tags *[]Tag `json:"tags,omitempty"`

	// **参数解释**：负载均衡器的创建时间。  **取值范围**： 格式：yyyy-MM-dd'T'HH:mm:ss'Z'
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释**：负载均衡器的更新时间。  **取值范围**： 格式：yyyy-MM-dd'T'HH:mm:ss'Z'
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释**：是否独享型LB。  **取值范围**： - false：共享型。 - true：独享型。
	Guaranteed *bool `json:"guaranteed,omitempty"`

	// **参数解释**：负载均衡器所在VPC ID。  **取值范围**：不涉及
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释**：负载均衡器绑定的EIP。 注：该字段与publicips一致。
	Eips *[]EipInfo `json:"eips,omitempty"`

	// **参数解释**：双栈类型负载均衡器的IPv6地址。  **取值范围**：不涉及  [不支持IPv6，请勿使用。](tag:dt)
	Ipv6VipAddress *string `json:"ipv6_vip_address,omitempty"`

	// **参数解释**：双栈类型负载均衡器所在子网的IPv6网络ID，也称为该负载均衡器实例的前端子网。  **取值范围**：不涉及  [不支持IPv6，请勿使用。](tag:dt)
	Ipv6VipVirsubnetId *string `json:"ipv6_vip_virsubnet_id,omitempty"`

	// **参数解释**：双栈类型负载均衡器的IPv6对应的port ID。  **取值范围**：不涉及  [不支持IPv6，请勿使用。](tag:dt)
	Ipv6VipPortId *string `json:"ipv6_vip_port_id,omitempty"`

	// **参数解释**：负载均衡器所在的可用区列表。  **取值范围**：不涉及
	AvailabilityZoneList *[]string `json:"availability_zone_list,omitempty"`

	// **参数解释**：资源所属的企业项目ID。  **取值范围**： - \"0\"：表示资源属于default企业项目。 - UUID格式的字符串，表示非默认企业项目。  [不支持该字段，请勿使用。](tag:dt,hcso_dt)
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**：资源账单信息。  **取值范围**： - 空：按需计费。 [- 非空：包周期计费，  包周期计费billing_info字段的格式为：order_id:product_id:region_id:project_id，如： CS2107161019CDJZZ:OFFI569702121789763584:az:057ef081eb00d2732fd1c01a9be75e6f  不支持该字段，请勿使用](tag:hws_eu,g42,hk_g42,dt,hcso_dt,hcso,hk_vdf,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b)
	BillingInfo *string `json:"billing_info,omitempty"`

	// **参数解释**：网络型规格ID。 对于弹性扩缩容实例，表示上限规格。  **取值范围**：不涉及 [hsco场景下所有LB实例共享带宽，该字段无效，请勿使用。](tag:hcso,hk_vdf,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b)
	L4FlavorId *string `json:"l4_flavor_id,omitempty"`

	// **参数解释**：四层弹性Flavor ID。  **取值范围**：不涉及  不支持该字段，请勿使用。
	L4ScaleFlavorId *string `json:"l4_scale_flavor_id,omitempty"`

	// **参数解释**：应用型ID。 对于弹性扩缩容实例，表示上限规格ID。  **取值范围**：不涉及  [hsco场景下所有LB实例共享带宽，该字段无效，请勿使用。](tag:hcso,hk_vdf,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b)
	L7FlavorId *string `json:"l7_flavor_id,omitempty"`

	// **参数解释**：七层弹性Flavor ID。  **取值范围**：不涉及  不支持该字段，请勿使用。
	L7ScaleFlavorId *string `json:"l7_scale_flavor_id,omitempty"`

	// **参数解释**：负载均衡器绑定的公网IP。只支持绑定一个公网IP。  注：该字段与eips一致。
	Publicips *[]PublicIpInfo `json:"publicips,omitempty"`

	// **参数解释**：负载均衡器绑定的global eip。  [不支持该字段，请勿使用。](tag:hws_eu,g42,hk_g42,dt,hcso_dt,hk_vdf,fcs,ctc,ocb,hws_ocb)
	GlobalEips *[]GlobalEipInfo `json:"global_eips,omitempty"`

	// **参数解释**：下联面子网的网络ID列表。  **取值范围**：不涉及
	ElbVirsubnetIds *[]string `json:"elb_virsubnet_ids,omitempty"`

	// **参数解释**：下联面子网类型。  **取值范围**： - ipv4：ipv4 - dualstack：双栈
	ElbVirsubnetType *RecycleLoadBalancerElbVirsubnetType `json:"elb_virsubnet_type,omitempty"`

	// **参数解释**：是否启用IP类型后端转发。 [开启IP类型后端转发后，后端服务器组不仅支持添加云上VPC内的服务器，还支持添加其他VPC、其他公有云、云下数据中心的服务器。](tag:hws,hws_hk,ocb,ctc,hcs,g42,tm,cmcc,hk_g42,hws_ocb,dt,hcso_dt,hws_eu) [开启IP类型后端转发后，后端服务器组不仅支持添加云上VPC内的服务器，还支持添加其他VPC、云下数据中心的服务器。](tag:fcs)  **取值范围**： - true：开启。 - false：不开启。  [荷兰region不支持该字段，请勿使用。](tag:dt)
	IpTargetEnable *bool `json:"ip_target_enable,omitempty"`

	// **参数解释**：负载均衡器的冻结场景。 [若负载均衡器有多个冻结场景，用逗号分隔。  **取值范围**： - POLICE：公安冻结场景。 - ILLEGAL：违规冻结场景。 - VERIFY：客户未实名认证冻结场景。 - PARTNER：合作伙伴冻结（合作伙伴冻结子客户资源）。 - AREAR：欠费冻结场景。](tag:hws,hws_hk)  [不支持该字段，请勿使用。](tag:hws_eu,g42,hk_g42,dt,hcso_dt,ocb,hws_ocb)
	FrozenScene *string `json:"frozen_scene,omitempty"`

	Ipv6Bandwidth *BandwidthRef `json:"ipv6_bandwidth,omitempty"`

	// **参数解释**：是否开启删除保护。  **取值范围**： - false：不开启。 - true：开启。  [不支持该字段，请勿使用。](tag:hws_eu,g42,hk_g42)  [荷兰region不支持该字段，请勿使用。](tag:dt)
	DeletionProtectionEnable *bool `json:"deletion_protection_enable,omitempty"`

	Autoscaling *AutoscalingRef `json:"autoscaling,omitempty"`

	// **参数解释**：LB所属AZ组。  **取值范围**：不涉及
	PublicBorderGroup *string `json:"public_border_group,omitempty"`

	// **参数解释**：负载均衡器实例的计费模式。  **取值范围**： - flavor：按规格计费 - lcu：按使用量计费 - 空值：若是共享型表示免费实例。若是独享型则与flavor模式一致，都是按规格计费。
	ChargeMode *string `json:"charge_mode,omitempty"`

	// **参数解释**：WAF故障时的流量处理策略。  **取值范围**：discard:丢弃，forward: 转发到后端。  [不支持该字段，请勿使用。](tag:hws_eu,hws_test,hcs,hcs_sm,hcso,hk_vdf,fcs,fcs_vm,mix,hcso_g42,hcso_g42_b,hcso_dt,dt,ocb,ctc,cmcc,tm,ct,sbc,g42,hws_ocb,hk_sbc,hk_tm,hk_g42)
	WafFailureAction *string `json:"waf_failure_action,omitempty"`

	// **参数解释**：修改保护状态。  **取值范围**： - nonProtection: 不保护。 - consoleProtection: 控制台修改保护。
	ProtectionStatus *RecycleLoadBalancerProtectionStatus `json:"protection_status,omitempty"`

	// **参数解释**：设置保护的原因。作为protection_status的转态设置的原因。  **取值范围**：除<和>外通用Unicode字符集字符，最大255个字符。
	ProtectionReason *string `json:"protection_reason,omitempty"`

	// **参数解释**：LB所绑定的logtank对应的group id。  **取值范围**：不涉及
	LogGroupId *string `json:"log_group_id,omitempty"`

	// **参数解释**：LB所绑定的logtank对应的topic id。  **取值范围**：不涉及
	LogTopicId *string `json:"log_topic_id,omitempty"`
}

func (o RecycleLoadBalancer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecycleLoadBalancer struct{}"
	}

	return strings.Join([]string{"RecycleLoadBalancer", string(data)}, " ")
}

type RecycleLoadBalancerElbVirsubnetType struct {
	value string
}

type RecycleLoadBalancerElbVirsubnetTypeEnum struct {
	IPV4      RecycleLoadBalancerElbVirsubnetType
	DUALSTACK RecycleLoadBalancerElbVirsubnetType
}

func GetRecycleLoadBalancerElbVirsubnetTypeEnum() RecycleLoadBalancerElbVirsubnetTypeEnum {
	return RecycleLoadBalancerElbVirsubnetTypeEnum{
		IPV4: RecycleLoadBalancerElbVirsubnetType{
			value: "ipv4",
		},
		DUALSTACK: RecycleLoadBalancerElbVirsubnetType{
			value: "dualstack",
		},
	}
}

func (c RecycleLoadBalancerElbVirsubnetType) Value() string {
	return c.value
}

func (c RecycleLoadBalancerElbVirsubnetType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecycleLoadBalancerElbVirsubnetType) UnmarshalJSON(b []byte) error {
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

type RecycleLoadBalancerProtectionStatus struct {
	value string
}

type RecycleLoadBalancerProtectionStatusEnum struct {
	NON_PROTECTION     RecycleLoadBalancerProtectionStatus
	CONSOLE_PROTECTION RecycleLoadBalancerProtectionStatus
}

func GetRecycleLoadBalancerProtectionStatusEnum() RecycleLoadBalancerProtectionStatusEnum {
	return RecycleLoadBalancerProtectionStatusEnum{
		NON_PROTECTION: RecycleLoadBalancerProtectionStatus{
			value: "nonProtection",
		},
		CONSOLE_PROTECTION: RecycleLoadBalancerProtectionStatus{
			value: "consoleProtection",
		},
	}
}

func (c RecycleLoadBalancerProtectionStatus) Value() string {
	return c.value
}

func (c RecycleLoadBalancerProtectionStatus) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *RecycleLoadBalancerProtectionStatus) UnmarshalJSON(b []byte) error {
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
