package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// ShowInstanceResponse Response Object
type ShowInstanceResponse struct {

	// 认证用户名，只能由英文字母、数字、中划线组成，长度为4~64的字符。
	AccessUser *string `json:"access_user,omitempty"`

	// 代理个数。
	BrokerNum *ShowInstanceResponseBrokerNum `json:"broker_num,omitempty"`

	// 实例名称。
	Name *string `json:"name,omitempty"`

	// 消息引擎。
	Engine *string `json:"engine,omitempty"`

	// 消息引擎版本。
	EngineVersion *string `json:"engine_version,omitempty"`

	// 实例规格。   - RabbitMQ实例单机返回vm规格。   - RabbitMQ实例集群返回vm规格和节点数。
	Specification *string `json:"specification,omitempty"`

	// 消息存储空间，单位：GB。
	StorageSpace *int32 `json:"storage_space,omitempty"`

	// 已使用的消息存储空间，单位：GB。
	UsedStorageSpace *int32 `json:"used_storage_space,omitempty"`

	// 实例是否开启域名访问功能。 - true：开启 - false：未开启
	DnsEnable *bool `json:"dns_enable,omitempty"`

	// 实例内网连接IP地址。
	ConnectAddress *string `json:"connect_address,omitempty"`

	// 实例内网连接域名。
	ConnectDomainName *string `json:"connect_domain_name,omitempty"`

	// 实例公网连接IP地址。
	PublicConnectAddress *string `json:"public_connect_address,omitempty"`

	// 实例公网连接域名。
	PublicConnectDomainName *string `json:"public_connect_domain_name,omitempty"`

	// 实例连接端口。
	Port *int32 `json:"port,omitempty"`

	// 实例的状态。
	Status *string `json:"status,omitempty"`

	// 实例描述。
	Description *string `json:"description,omitempty"`

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 资源规格标识。   [- dms.instance.rabbitmq.single.c3.2u4g：RabbitMQ单机，vm规格2u4g   - dms.instance.rabbitmq.single.c3.4u8g：RabbitMQ单机，vm规格4u8g   - dms.instance.rabbitmq.single.c3.8u16g：RabbitMQ单机，vm规格8u16g   - dms.instance.rabbitmq.single.c3.16u32g：RabbitMQ单机，vm规格16u32g   - dms.instance.rabbitmq.cluster.c3.4u8g.3：RabbitMQ集群，vm规格4u8g，3个节点   - dms.instance.rabbitmq.cluster.c3.4u8g.5：RabbitMQ集群，vm规格4u8g，5个节点   - dms.instance.rabbitmq.cluster.c3.4u8g.7：RabbitMQ集群，vm规格4u8g，7个节点   - dms.instance.rabbitmq.cluster.c3.8u16g.3：RabbitMQ集群，vm规格8u16g，3个节点   - dms.instance.rabbitmq.cluster.c3.8u16g.5：RabbitMQ集群，vm规格8u16g，5个节点   - dms.instance.rabbitmq.cluster.c3.8u16g.7：RabbitMQ集群，vm规格8u16g，7个节点   - dms.instance.rabbitmq.cluster.c3.16u32g.3：RabbitMQ集群，vm规格16u32g，3个节点   - dms.instance.rabbitmq.cluster.c3.16u32g.5：RabbitMQ集群，vm规格16u32g，5个节点   - dms.instance.rabbitmq.cluster.c3.16u32g.7：RabbitMQ集群，vm规格16u32g，7个节点](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,hk_tm,sbc)
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// 付费模式，1表示按需计费，0表示包年/包月计费。
	ChargingMode *int32 `json:"charging_mode,omitempty"`

	// VPC ID。
	VpcId *string `json:"vpc_id,omitempty"`

	// VPC的名称。
	VpcName *string `json:"vpc_name,omitempty"`

	// 完成创建时间。格式为时间戳，指从格林威治时间 1970年01月01日00时00分00秒起至指定时间的偏差总毫秒数。
	CreatedAt *string `json:"created_at,omitempty"`

	// 用户ID。
	UserId *string `json:"user_id,omitempty"`

	// 用户名。
	UserName *string `json:"user_name,omitempty"`

	// 订单ID，只有在包周期计费时才会有order_id值，其他计费方式order_id值为空。
	OrderId *string `json:"order_id,omitempty"`

	// 维护时间窗开始时间，格式为HH:mm:ss。
	MaintainBegin *string `json:"maintain_begin,omitempty"`

	// 维护时间窗结束时间，格式为HH:mm:ss。
	MaintainEnd *string `json:"maintain_end,omitempty"`

	// RabbitMQ实例是否开启公网访问功能。   - true：开启   - false：未开启
	EnablePublicip *bool `json:"enable_publicip,omitempty"`

	// RabbitMQ实例绑定的弹性IP地址。  如果未开启公网访问功能，该字段值为null。
	PublicipAddress *string `json:"publicip_address,omitempty"`

	// RabbitMQ实例绑定的弹性IP地址的ID。  如果未开启公网访问功能，该字段值为null。
	PublicipId *string `json:"publicip_id,omitempty"`

	// RabbitMQ实例的管理地址。
	ManagementConnectAddress *string `json:"management_connect_address,omitempty"`

	// RabbitMQ实例的管理域名。
	ManagementConnectDomainName *string `json:"management_connect_domain_name,omitempty"`

	// RabbitMQ实例的公网管理地址。
	PublicManagementConnectAddress *string `json:"public_management_connect_address,omitempty"`

	// RabbitMQ实例的公网管理域名。
	PublicManagementConnectDomainName *string `json:"public_management_connect_domain_name,omitempty"`

	// 是否开启安全认证。   - true：开启   - false：未开启
	SslEnable *bool `json:"ssl_enable,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 实例扩容时用于区分老实例与新实例。 - true：新创建的实例，允许磁盘动态扩容不需要重启。 - false：老实例
	IsLogicalVolume *bool `json:"is_logical_volume,omitempty"`

	// 实例扩容磁盘次数，如果超过20次则无法扩容磁盘。
	ExtendTimes *int32 `json:"extend_times,omitempty"`

	// 实例类型：集群，cluster。
	Type *ShowInstanceResponseType `json:"type,omitempty"`

	// 产品标识。
	ProductId *string `json:"product_id,omitempty"`

	// 安全组ID。
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// 租户安全组名称。
	SecurityGroupName *string `json:"security_group_name,omitempty"`

	// 子网ID。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 实例节点所在的可用区，返回“可用区ID”。
	AvailableZones *[]string `json:"available_zones,omitempty"`

	// 实例节点所在的可用区，返回“可用区名称”。
	AvailableZoneNames *[]string `json:"available_zone_names,omitempty"`

	// 总共消息存储空间，单位：GB。
	TotalStorageSpace *int32 `json:"total_storage_space,omitempty"`

	// 存储资源ID。
	StorageResourceId *string `json:"storage_resource_id,omitempty"`

	// IO规格。
	StorageSpecCode *string `json:"storage_spec_code,omitempty"`

	// 是否开启ipv6。
	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	// IPv6的连接地址。
	Ipv6ConnectAddresses *[]string `json:"ipv6_connect_addresses,omitempty"`

	// 标签列表。
	Tags           *[]TagEntity `json:"tags,omitempty"`
	HttpStatusCode int          `json:"-"`
}

func (o ShowInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceResponse", string(data)}, " ")
}

type ShowInstanceResponseBrokerNum struct {
	value int32
}

type ShowInstanceResponseBrokerNumEnum struct {
	E_1 ShowInstanceResponseBrokerNum
	E_3 ShowInstanceResponseBrokerNum
	E_5 ShowInstanceResponseBrokerNum
	E_7 ShowInstanceResponseBrokerNum
}

func GetShowInstanceResponseBrokerNumEnum() ShowInstanceResponseBrokerNumEnum {
	return ShowInstanceResponseBrokerNumEnum{
		E_1: ShowInstanceResponseBrokerNum{
			value: 1,
		}, E_3: ShowInstanceResponseBrokerNum{
			value: 3,
		}, E_5: ShowInstanceResponseBrokerNum{
			value: 5,
		}, E_7: ShowInstanceResponseBrokerNum{
			value: 7,
		},
	}
}

func (c ShowInstanceResponseBrokerNum) Value() int32 {
	return c.value
}

func (c ShowInstanceResponseBrokerNum) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceResponseBrokerNum) UnmarshalJSON(b []byte) error {
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

type ShowInstanceResponseType struct {
	value string
}

type ShowInstanceResponseTypeEnum struct {
	SINGLE  ShowInstanceResponseType
	CLUSTER ShowInstanceResponseType
}

func GetShowInstanceResponseTypeEnum() ShowInstanceResponseTypeEnum {
	return ShowInstanceResponseTypeEnum{
		SINGLE: ShowInstanceResponseType{
			value: "single",
		},
		CLUSTER: ShowInstanceResponseType{
			value: "cluster",
		},
	}
}

func (c ShowInstanceResponseType) Value() string {
	return c.value
}

func (c ShowInstanceResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceResponseType) UnmarshalJSON(b []byte) error {
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
