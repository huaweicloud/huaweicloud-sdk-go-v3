package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type InstanceCreateReq struct {

	// 实例描述。支持除>和<以外的字符，长度为0~255。
	Description *string `json:"description,omitempty"`

	// 维护时间窗开始时间。时间格式为 xx:00:00，xx取值为02,06,10,14,18,22。  在这个时间段内，运维人员可以对该实例的节点进行维护操作。维护期间，业务可以正常使用，可能会发生闪断。维护操作通常几个月一次。
	MaintainBegin *string `json:"maintain_begin,omitempty"`

	// 维护时间窗结束时间。时间格式为 xx:00:00，与维护时间窗开始时间相差4个小时。  在这个时间段内，运维人员可以对该实例的节点进行维护操作。维护期间，业务可以正常使用，可能会发生闪断。维护操作通常几个月一次。
	MaintainEnd *string `json:"maintain_end,omitempty"`

	// 实例名称。  中英文字符开头，只能由中英文字符、数字、中划线、下划线组成，长度为3~64。  > 中文字符必须为UTF-8或者unicode编码。
	InstanceName *string `json:"instance_name,omitempty"`

	// 实例编号，不填写自动生成
	InstanceId *string `json:"instance_id,omitempty"`

	// 实例规格： - BASIC：基础版实例 - PROFESSIONAL：专业版实例 - ENTERPRISE：企业版实例 - PLATINUM：铂金版实例 - BASIC_IPV6：基础版IPV6实例 - PROFESSIONAL_IPV6：专业版IPV6实例 - ENTERPRISE_IPV6：企业版IPV6实例 - PLATINUM_IPV6：铂金版IPV6实例 - PLATINUM_X2：铂金版 x2实例 - PLATINUM_X3：铂金版 x3实例 - PLATINUM_X4：铂金版 x4实例 - PLATINUM_X5：铂金版 x5实例 - PLATINUM_X6：铂金版 x6实例 - PLATINUM_X7：铂金版 x7实例 - PLATINUM_X8：铂金版 x8实例  当前仅部分region支持铂金版 x2、铂金版 x3、铂金版 x4、铂金版 x5、铂金版 x6、铂金版 x7、铂金版 x8
	SpecId *InstanceCreateReqSpecId `json:"spec_id,omitempty"`

	// 虚拟私有云ID。  获取方法如下：   - 方法1：登录虚拟私有云服务的控制台界面，在虚拟私有云的详情页面查找VPC ID。   - 方法2：通过虚拟私有云服务的API接口查询，具体方法请参见《虚拟私有云服务API参考》的“查询VPC列表”章节。
	VpcId *string `json:"vpc_id,omitempty"`

	// 子网的网络ID。  获取方法如下： - 方法1：登录虚拟私有云服务的控制台界面，单击VPC下的子网，进入子网详情页面，查找网络ID。 - 方法2：通过虚拟私有云服务的API接口查询，具体方法请参见《虚拟私有云服务API参考》的“查询子网列表”章节。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 指定实例所属的安全组。  获取方法如下： - 方法1：登录虚拟私有云服务的控制台界面，在安全组的详情页面查找安全组ID。 - 方法2：通过虚拟私有云服务的API接口查询，具体方法请参见《虚拟私有云服务API参考》的“查询安全组列表”章节。
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// 弹性公网IP ID。  实例需要开启公网访问，且loadbalancer_provider为lvs时需要填写，绑定后使用者可以通过该入口从公网访问APIG实例中的API等资源  获取方法：登录虚拟私有云服务的控制台界面，在弹性公网IP的详情页面查找弹性公网IP ID。
	EipId *string `json:"eip_id,omitempty"`

	// 企业项目ID，企业账号必填。  获取方法如下： - 方法1：登录企业项目管理界面，在项目管理详情页面查找项目ID。 - 方法2：通过企业项目管理的API接口查询，具体方法请参见《企业管理API参考》的“查询企业项目列表”章节。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 可用区列表。  可用区指在同一地域下，电力、网络隔离的物理区域，可用区之内内网互通，不同可用区之间物理隔离。选择多个AZ部署可以有效提升可靠性。  获取方法：通过文档中实例管理的可用区列表接口查询。
	AvailableZoneIds *[]string `json:"available_zone_ids,omitempty"`

	// 出公网带宽  实例需要开启出公网功能时需要填写，绑定后使用者可以利用该出口访问公网上的互联网资源
	BandwidthSize *int32 `json:"bandwidth_size,omitempty"`

	// 出公网带宽计费类型，实例需要开启出公网功能时需要填写： - bandwidth：按带宽计费 - traffic：按流量计费
	BandwidthChargingMode *InstanceCreateReqBandwidthChargingMode `json:"bandwidth_charging_mode,omitempty"`

	// 公网访问是否支持IPv6。  当前仅部分region部分可用区支持IPv6
	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	// 实例使用的负载均衡器类型 - 该字段废弃，由Region支持的负载均衡模式决定使用elb还是lvs，详情参考产品介绍下的约束与限制。 - lvs Linux虚拟服务器 - elb 弹性负载均衡，elb仅部分region支持
	LoadbalancerProvider *InstanceCreateReqLoadbalancerProvider `json:"loadbalancer_provider,omitempty"`

	// 标签列表。  一个实例默认最多支持创建20个标签
	Tags *[]TmsKeyValue `json:"tags,omitempty"`

	// 终端节点服务的名称。  支持英文、数字、中划线、下划线，0~16个字符。  如果您不填写该参数，系统生成的终端节点服务的名称为{region}.apig.{service_id}。 如果您填写该参数，系统生成的终端节点服务的名称为{region}.{vpcep_service_name}.{service_id}。 实例创建完成后，可以在实例管理->终端节点管理页面修改该名称。
	VpcepServiceName *string `json:"vpcep_service_name,omitempty"`

	// 入公网带宽  实例需要开启入公网功能，且loadbalancer_provider为elb时需要填写，绑定后使用者可以通过该入口从公网访问APIG实例中的API等资源
	IngressBandwidthSize *int32 `json:"ingress_bandwidth_size,omitempty"`

	// 入公网带宽计费类型，实例需要开启入公网功能，且loadbalancer_provider为elb时需要填写： - bandwidth：按带宽计费 - traffic：按流量计费
	IngressBandwidthChargingMode *InstanceCreateReqIngressBandwidthChargingMode `json:"ingress_bandwidth_charging_mode,omitempty"`
}

func (o InstanceCreateReq) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceCreateReq struct{}"
	}

	return strings.Join([]string{"InstanceCreateReq", string(data)}, " ")
}

type InstanceCreateReqSpecId struct {
	value string
}

type InstanceCreateReqSpecIdEnum struct {
	BASIC             InstanceCreateReqSpecId
	PROFESSIONAL      InstanceCreateReqSpecId
	ENTERPRISE        InstanceCreateReqSpecId
	PLATINUM          InstanceCreateReqSpecId
	BASIC_IPV6        InstanceCreateReqSpecId
	PROFESSIONAL_IPV6 InstanceCreateReqSpecId
	ENTERPRISE_IPV6   InstanceCreateReqSpecId
	PLATINUM_IPV6     InstanceCreateReqSpecId
	PLATINUM_X2       InstanceCreateReqSpecId
	PLATINUM_X3       InstanceCreateReqSpecId
	PLATINUM_X4       InstanceCreateReqSpecId
	PLATINUM_X5       InstanceCreateReqSpecId
	PLATINUM_X6       InstanceCreateReqSpecId
	PLATINUM_X7       InstanceCreateReqSpecId
	PLATINUM_X8       InstanceCreateReqSpecId
}

func GetInstanceCreateReqSpecIdEnum() InstanceCreateReqSpecIdEnum {
	return InstanceCreateReqSpecIdEnum{
		BASIC: InstanceCreateReqSpecId{
			value: "BASIC",
		},
		PROFESSIONAL: InstanceCreateReqSpecId{
			value: "PROFESSIONAL",
		},
		ENTERPRISE: InstanceCreateReqSpecId{
			value: "ENTERPRISE",
		},
		PLATINUM: InstanceCreateReqSpecId{
			value: "PLATINUM",
		},
		BASIC_IPV6: InstanceCreateReqSpecId{
			value: "BASIC_IPV6",
		},
		PROFESSIONAL_IPV6: InstanceCreateReqSpecId{
			value: "PROFESSIONAL_IPV6",
		},
		ENTERPRISE_IPV6: InstanceCreateReqSpecId{
			value: "ENTERPRISE_IPV6",
		},
		PLATINUM_IPV6: InstanceCreateReqSpecId{
			value: "PLATINUM_IPV6",
		},
		PLATINUM_X2: InstanceCreateReqSpecId{
			value: "PLATINUM_X2",
		},
		PLATINUM_X3: InstanceCreateReqSpecId{
			value: "PLATINUM_X3",
		},
		PLATINUM_X4: InstanceCreateReqSpecId{
			value: "PLATINUM_X4",
		},
		PLATINUM_X5: InstanceCreateReqSpecId{
			value: "PLATINUM_X5",
		},
		PLATINUM_X6: InstanceCreateReqSpecId{
			value: "PLATINUM_X6",
		},
		PLATINUM_X7: InstanceCreateReqSpecId{
			value: "PLATINUM_X7",
		},
		PLATINUM_X8: InstanceCreateReqSpecId{
			value: "PLATINUM_X8",
		},
	}
}

func (c InstanceCreateReqSpecId) Value() string {
	return c.value
}

func (c InstanceCreateReqSpecId) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceCreateReqSpecId) UnmarshalJSON(b []byte) error {
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

type InstanceCreateReqBandwidthChargingMode struct {
	value string
}

type InstanceCreateReqBandwidthChargingModeEnum struct {
	BANDWIDTH InstanceCreateReqBandwidthChargingMode
	TRAFFIC   InstanceCreateReqBandwidthChargingMode
}

func GetInstanceCreateReqBandwidthChargingModeEnum() InstanceCreateReqBandwidthChargingModeEnum {
	return InstanceCreateReqBandwidthChargingModeEnum{
		BANDWIDTH: InstanceCreateReqBandwidthChargingMode{
			value: "bandwidth",
		},
		TRAFFIC: InstanceCreateReqBandwidthChargingMode{
			value: "traffic",
		},
	}
}

func (c InstanceCreateReqBandwidthChargingMode) Value() string {
	return c.value
}

func (c InstanceCreateReqBandwidthChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceCreateReqBandwidthChargingMode) UnmarshalJSON(b []byte) error {
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

type InstanceCreateReqLoadbalancerProvider struct {
	value string
}

type InstanceCreateReqLoadbalancerProviderEnum struct {
	LVS InstanceCreateReqLoadbalancerProvider
	ELB InstanceCreateReqLoadbalancerProvider
}

func GetInstanceCreateReqLoadbalancerProviderEnum() InstanceCreateReqLoadbalancerProviderEnum {
	return InstanceCreateReqLoadbalancerProviderEnum{
		LVS: InstanceCreateReqLoadbalancerProvider{
			value: "lvs",
		},
		ELB: InstanceCreateReqLoadbalancerProvider{
			value: "elb",
		},
	}
}

func (c InstanceCreateReqLoadbalancerProvider) Value() string {
	return c.value
}

func (c InstanceCreateReqLoadbalancerProvider) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceCreateReqLoadbalancerProvider) UnmarshalJSON(b []byte) error {
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

type InstanceCreateReqIngressBandwidthChargingMode struct {
	value string
}

type InstanceCreateReqIngressBandwidthChargingModeEnum struct {
	BANDWIDTH InstanceCreateReqIngressBandwidthChargingMode
	TRAFFIC   InstanceCreateReqIngressBandwidthChargingMode
}

func GetInstanceCreateReqIngressBandwidthChargingModeEnum() InstanceCreateReqIngressBandwidthChargingModeEnum {
	return InstanceCreateReqIngressBandwidthChargingModeEnum{
		BANDWIDTH: InstanceCreateReqIngressBandwidthChargingMode{
			value: "bandwidth",
		},
		TRAFFIC: InstanceCreateReqIngressBandwidthChargingMode{
			value: "traffic",
		},
	}
}

func (c InstanceCreateReqIngressBandwidthChargingMode) Value() string {
	return c.value
}

func (c InstanceCreateReqIngressBandwidthChargingMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceCreateReqIngressBandwidthChargingMode) UnmarshalJSON(b []byte) error {
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
