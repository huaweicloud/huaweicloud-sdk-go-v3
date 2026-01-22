package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ShowInstanceResp struct {

	// **参数解释**： 认证用户名。 **取值范围**： 不涉及。
	AccessUser *string `json:"access_user,omitempty"`

	// **参数解释**： 代理个数。 **取值范围**： - 1 - 3 - 5 - 7
	BrokerNum *ShowInstanceRespBrokerNum `json:"broker_num,omitempty"`

	// **参数解释**： 实例名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 消息引擎类型。 **取值范围**： rabbitmq：RabbitMQ引擎。
	Engine *string `json:"engine,omitempty"`

	// **参数解释**： 消息引擎版本。 **取值范围**： - 3.8.35 [- 3.12.13](tag:srg) [- AMQP-0-9-1](tag:hws,hws_hk,hws_eu)
	EngineVersion *string `json:"engine_version,omitempty"`

	// **参数解释**： 实例规格。 **取值范围**： - 单机实例：返回vm规格。 - 集群实例：返回vm规格和节点数。
	Specification *string `json:"specification,omitempty"`

	// **参数解释**： 消息存储空间，单位：GB。 **取值范围**： 不涉及。
	StorageSpace *int32 `json:"storage_space,omitempty"`

	// **参数解释**： 已使用的消息存储空间，单位：GB。 **取值范围**： 不涉及。
	UsedStorageSpace *int32 `json:"used_storage_space,omitempty"`

	// **参数解释**： 实例是否开启域名访问功能。 **取值范围**： - true：开启 - false：未开启
	DnsEnable *bool `json:"dns_enable,omitempty"`

	// **参数解释**： 实例内网连接IP地址。 **取值范围**： 不涉及。
	ConnectAddress *string `json:"connect_address,omitempty"`

	// **参数解释**： 实例内网连接域名。 **取值范围**： 不涉及。
	ConnectDomainName *string `json:"connect_domain_name,omitempty"`

	// **参数解释**： 实例公网连接IP地址。 **取值范围**： 不涉及。
	PublicConnectAddress *string `json:"public_connect_address,omitempty"`

	// **参数解释**： 实例公网连接域名。 **取值范围**： 不涉及。
	PublicConnectDomainName *string `json:"public_connect_domain_name,omitempty"`

	// **参数解释**： 实例连接端口。 **取值范围**： 不涉及。
	Port *int32 `json:"port,omitempty"`

	// **参数解释**： 实例状态。 **取值范围**： [详细状态说明请参考[实例状态说明](rabbitmq-api-180514012.xml)](tag:hws,hws_eu,hws_hk,cmcc,ctc,sbc,hk_sbc,g42,hk_g42,tm,hk_tm,ax)[详细状态说明请参考[实例状态说明](kafka-api-180514012.xml)](tag:hcs)。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 实例描述。 **取值范围**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 实例ID。 **取值范围**： 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**： 资源规格标识。 **取值范围**： -[ dms.instance.rabbitmq.single.c3.2u4g：RabbitMQ单机，vm规格2u4g - dms.instance.rabbitmq.single.c3.4u8g：RabbitMQ单机，vm规格4u8g - dms.instance.rabbitmq.single.c3.8u16g：RabbitMQ单机，vm规格8u16g - dms.instance.rabbitmq.single.c3.16u32g：RabbitMQ单机，vm规格16u32g - dms.instance.rabbitmq.cluster.c3.4u8g.3：RabbitMQ集群，vm规格4u8g，3个节点 - dms.instance.rabbitmq.cluster.c3.4u8g.5：RabbitMQ集群，vm规格4u8g，5个节点 - dms.instance.rabbitmq.cluster.c3.4u8g.7：RabbitMQ集群，vm规格4u8g，7个节点 - dms.instance.rabbitmq.cluster.c3.8u16g.3：RabbitMQ集群，vm规格8u16g，3个节点 - dms.instance.rabbitmq.cluster.c3.8u16g.5：RabbitMQ集群，vm规格8u16g，5个节点 - dms.instance.rabbitmq.cluster.c3.8u16g.7：RabbitMQ集群，vm规格8u16g，7个节点 - dms.instance.rabbitmq.cluster.c3.16u32g.3：RabbitMQ集群，vm规格16u32g，3个节点 - dms.instance.rabbitmq.cluster.c3.16u32g.5：RabbitMQ集群，vm规格16u32g，5个节点 - dms.instance.rabbitmq.cluster.c3.16u32g.7：RabbitMQ集群，vm规格16u32g，7个节点](tag:hws,hws_eu,hws_hk,ctc,g42,hk_g42,tm,hk_tm,sbc,ax,hk_sbc)
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// **参数解释**： 付费模式。 **取值范围**： - 1：按需计费。 - 0：包年/包月计费。
	ChargingMode *int32 `json:"charging_mode,omitempty"`

	// **参数解释**： VPC ID。 **取值范围**： 不涉及。
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释**： VPC的名称。 **取值范围**： 不涉及。
	VpcName *string `json:"vpc_name,omitempty"`

	// **参数解释**： 完成创建时间。格式为时间戳，指从格林威治时间 1970年01月01日00时00分00秒起至指定时间的偏差总毫秒数。 **取值范围**： 不涉及。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释**： 用户ID。 **取值范围**： 不涉及。
	UserId *string `json:"user_id,omitempty"`

	// **参数解释**： 用户名。 **取值范围**： 不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**： 订单ID，只有在包周期计费时才会有order_id值，其他计费方式order_id值为空。 **取值范围**： 不涉及。
	OrderId *string `json:"order_id,omitempty"`

	// **参数解释**： 维护时间窗开始时间，格式为HH:mm:ss。 **取值范围**： 不涉及。
	MaintainBegin *string `json:"maintain_begin,omitempty"`

	// **参数解释**： 维护时间窗结束时间，格式为HH:mm:ss。 **取值范围**： 不涉及。
	MaintainEnd *string `json:"maintain_end,omitempty"`

	// **参数解释**： RabbitMQ实例是否开启公网访问功能。 **取值范围**： - true：开启 - false：未开启
	EnablePublicip *bool `json:"enable_publicip,omitempty"`

	// **参数解释**： RabbitMQ实例绑定的弹性IP地址。  如果未开启公网访问功能，该字段值为null。 **取值范围**： - true：开启 - false：未开启
	PublicipAddress *string `json:"publicip_address,omitempty"`

	// **参数解释**： RabbitMQ实例绑定的弹性IP地址的ID。  如果未开启公网访问功能，该字段值为null。 **取值范围**： 不涉及。
	PublicipId *string `json:"publicip_id,omitempty"`

	// **参数解释**： RabbitMQ实例的管理地址。 **取值范围**： 不涉及。
	ManagementConnectAddress *string `json:"management_connect_address,omitempty"`

	// **参数解释**： RabbitMQ实例的管理域名。 **取值范围**： 不涉及。
	ManagementConnectDomainName *string `json:"management_connect_domain_name,omitempty"`

	// **参数解释**： RabbitMQ实例的公网管理地址。 **取值范围**： 不涉及。
	PublicManagementConnectAddress *string `json:"public_management_connect_address,omitempty"`

	// **参数解释**： RabbitMQ实例的公网管理域名。 **取值范围**： 不涉及。
	PublicManagementConnectDomainName *string `json:"public_management_connect_domain_name,omitempty"`

	// **参数解释**： 是否开启安全认证。 **取值范围**： - true：开启 - false：未开启
	SslEnable *bool `json:"ssl_enable,omitempty"`

	// **参数解释**： 企业项目ID。 **取值范围**： 不涉及。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 实例扩容时用于区分老实例与新实例。 **取值范围**： - true：新创建的实例，允许磁盘动态扩容不需要重启。 - false：特别老的实例不支持磁盘扩容。
	IsLogicalVolume *bool `json:"is_logical_volume,omitempty"`

	// **参数解释**： 实例扩容磁盘次数。 **取值范围**： 不涉及。
	ExtendTimes *int32 `json:"extend_times,omitempty"`

	// **参数解释**： 实例类型。 **取值范围**： - single：单机。 - cluster：集群。
	Type *ShowInstanceRespType `json:"type,omitempty"`

	// **参数解释**： 产品标识。 **取值范围**： 不涉及。
	ProductId *string `json:"product_id,omitempty"`

	// **参数解释**： 安全组ID。 **取值范围**： 不涉及。
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// **参数解释**： 租户安全组名称。 **取值范围**： 不涉及。
	SecurityGroupName *string `json:"security_group_name,omitempty"`

	// **参数解释**： 子网ID。 **取值范围**： 不涉及。
	SubnetId *string `json:"subnet_id,omitempty"`

	// **参数解释**： 实例节点所在的可用区ID。
	AvailableZones *[]string `json:"available_zones,omitempty"`

	// **参数解释**： 实例节点所在的可用区名称。
	AvailableZoneNames *[]string `json:"available_zone_names,omitempty"`

	// **参数解释**： 总共消息存储空间，单位：GB。 **取值范围**： 不涉及。
	TotalStorageSpace *int32 `json:"total_storage_space,omitempty"`

	// **参数解释**： 存储资源ID。 **取值范围**： 不涉及。
	StorageResourceId *string `json:"storage_resource_id,omitempty"`

	// **参数解释**： IO规格。 **取值范围**： 不涉及。
	StorageSpecCode *string `json:"storage_spec_code,omitempty"`

	// **参数解释**： 是否开启IPv6。 **取值范围**： - true：开启。 - false：不开启。
	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	// **参数解释**： IPv6的连接地址。
	Ipv6ConnectAddresses *[]string `json:"ipv6_connect_addresses,omitempty"`

	// **参数解释**： 标签列表。
	Tags *[]TagEntity `json:"tags,omitempty"`

	// **参数解释**： 服务类型。 **取值范围**： advanced：服务类型。
	ServiceType *string `json:"service_type,omitempty"`

	// **参数解释**： 存储类型。 **取值范围**： hec：存储类型。
	StorageType *string `json:"storage_type,omitempty"`
}

func (o ShowInstanceResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceResp struct{}"
	}

	return strings.Join([]string{"ShowInstanceResp", string(data)}, " ")
}

type ShowInstanceRespBrokerNum struct {
	value int32
}

type ShowInstanceRespBrokerNumEnum struct {
	E_1 ShowInstanceRespBrokerNum
	E_3 ShowInstanceRespBrokerNum
	E_5 ShowInstanceRespBrokerNum
	E_7 ShowInstanceRespBrokerNum
}

func GetShowInstanceRespBrokerNumEnum() ShowInstanceRespBrokerNumEnum {
	return ShowInstanceRespBrokerNumEnum{
		E_1: ShowInstanceRespBrokerNum{
			value: 1,
		}, E_3: ShowInstanceRespBrokerNum{
			value: 3,
		}, E_5: ShowInstanceRespBrokerNum{
			value: 5,
		}, E_7: ShowInstanceRespBrokerNum{
			value: 7,
		},
	}
}

func (c ShowInstanceRespBrokerNum) Value() int32 {
	return c.value
}

func (c ShowInstanceRespBrokerNum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceRespBrokerNum) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("int32")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: int32")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(int32); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to int32 error")
	}
}

type ShowInstanceRespType struct {
	value string
}

type ShowInstanceRespTypeEnum struct {
	SINGLE  ShowInstanceRespType
	CLUSTER ShowInstanceRespType
}

func GetShowInstanceRespTypeEnum() ShowInstanceRespTypeEnum {
	return ShowInstanceRespTypeEnum{
		SINGLE: ShowInstanceRespType{
			value: "single",
		},
		CLUSTER: ShowInstanceRespType{
			value: "cluster",
		},
	}
}

func (c ShowInstanceRespType) Value() string {
	return c.value
}

func (c ShowInstanceRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceRespType) UnmarshalJSON(b []byte) error {
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
