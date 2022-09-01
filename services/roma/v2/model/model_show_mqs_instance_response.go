package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowMqsInstanceResponse struct {

	// 实例名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 描述信息。
	Description *string `json:"description,omitempty" xml:"description"`

	// 引擎。
	Engine *string `json:"engine,omitempty" xml:"engine"`

	// 版本。
	EngineVersion *string `json:"engine_version,omitempty" xml:"engine_version"`

	// 实例规格。
	Specification *string `json:"specification,omitempty" xml:"specification"`

	// 消息存储空间，单位：GB。
	StorageSpace *int32 `json:"storage_space,omitempty" xml:"storage_space"`

	// 最大分区数。不同规格的ROMA Connect实例的最大分区数不相同。
	PartitionNum *string `json:"partition_num,omitempty" xml:"partition_num"`

	// 已使用的消息存储空间，单位：GB。
	UsedStorageSpace *int32 `json:"used_storage_space,omitempty" xml:"used_storage_space"`

	// 实例连接IP地址。
	ConnectAddress *string `json:"connect_address,omitempty" xml:"connect_address"`

	// 实例连接端口。
	Port *int32 `json:"port,omitempty" xml:"port"`

	// 实例的状态。   - CREATING: 申请实例后，在实例状态进入运行中之前的状态。   - RUNNING: 实例正常运行状态。在这个状态的实例可以运行您的业务。
	Status *string `json:"status,omitempty" xml:"status"`

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id"`

	// 资源规格标识。  - dms.instance.kafka.cluster.c3.mini：Kafka实例的基准带宽为100MByte/秒。  - dms.instance.kafka.cluster.c3.small.2：Kafka实例的基准带宽为300MByte/秒。  - dms.instance.kafka.cluster.c3.middle.2：Kafka实例的基准带宽为600MByte/秒。  - dms.instance.kafka.cluster.c3.high.2：Kafka实例的基准带宽为1200MByte/秒。
	ResourceSpecCode *string `json:"resource_spec_code,omitempty" xml:"resource_spec_code"`

	// 付费模式，1表示按需计费，0表示包周期计费。
	ChargingMode *int32 `json:"charging_mode,omitempty" xml:"charging_mode"`

	// VPC ID。
	VpcId *string `json:"vpc_id,omitempty" xml:"vpc_id"`

	// VPC的名称。
	VpcName *string `json:"vpc_name,omitempty" xml:"vpc_name"`

	// 完成创建时间。  格式为时间戳，指从格林威治时间 1970年01月01日00时00分00秒起至指定时间的偏差总毫秒数。
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 用户ID。
	UserId *string `json:"user_id,omitempty" xml:"user_id"`

	// 用户名。
	UserName *string `json:"user_name,omitempty" xml:"user_name"`

	// 订单ID，只有在包周期计费时才会有order_id值，其他计费方式order_id值为空。
	OrderId *string `json:"order_id,omitempty" xml:"order_id"`

	// 维护时间窗开始时间，格式为HH:mm:ss。
	MaintainBegin *string `json:"maintain_begin,omitempty" xml:"maintain_begin"`

	// 维护时间窗结束时间，格式为HH:mm:ss。
	MaintainEnd *string `json:"maintain_end,omitempty" xml:"maintain_end"`

	// 实例是否开启公网访问功能。  - true：开启  - false：未开启
	EnablePublicip *bool `json:"enable_publicip,omitempty" xml:"enable_publicip"`

	// Kafka实例的KafkaManager连接地址。
	ManagementConnectAddress *string `json:"management_connect_address,omitempty" xml:"management_connect_address"`

	// 是否开启安全认证。 - true：开启 - false：未开启
	SslEnable *bool `json:"ssl_enable,omitempty" xml:"ssl_enable"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 实例扩容时用于区分老实例与新实例。  - true：新创建的实例，允许磁盘动态扩容不需要重启。  - false：老实例
	IsLogicalVolume *bool `json:"is_logical_volume,omitempty" xml:"is_logical_volume"`

	// 实例扩容磁盘次数，如果超过20次则无法扩容磁盘。
	ExtendTimes *int32 `json:"extend_times,omitempty" xml:"extend_times"`

	// 是否打开kafka自动创建topic功能。  - true：开启  - false：关闭
	EnableAutoTopic *bool `json:"enable_auto_topic,omitempty" xml:"enable_auto_topic"`

	// 实例类型：集群，cluster。
	Type *ShowMqsInstanceResponseType `json:"type,omitempty" xml:"type"`

	// 产品标识。
	ProductId *string `json:"product_id,omitempty" xml:"product_id"`

	// 安全组ID。
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id"`

	// 租户安全组名称。
	SecurityGroupName *string `json:"security_group_name,omitempty" xml:"security_group_name"`

	// 子网ID。
	SubnetId *string `json:"subnet_id,omitempty" xml:"subnet_id"`

	// 子网名称。
	SubnetName *string `json:"subnet_name,omitempty" xml:"subnet_name"`

	// 子网网段。
	SubnetCidr *string `json:"subnet_cidr,omitempty" xml:"subnet_cidr"`

	// 实例节点所在的可用区，返回“可用区ID”。
	AvailableZones *[]string `json:"available_zones,omitempty" xml:"available_zones"`

	// 总共消息存储空间，单位：GB。
	TotalStorageSpace *int32 `json:"total_storage_space,omitempty" xml:"total_storage_space"`

	// 实例公网连接IP地址。当实例开启了公网访问，实例才包含该参数。
	PublicConnectAddress *string `json:"public_connect_address,omitempty" xml:"public_connect_address"`

	// 存储资源ID。
	StorageResourceId *string `json:"storage_resource_id,omitempty" xml:"storage_resource_id"`

	// IO规格。
	StorageSpecCode *string `json:"storage_spec_code,omitempty" xml:"storage_spec_code"`

	// 服务类型。
	ServiceType *string `json:"service_type,omitempty" xml:"service_type"`

	// 存储类型。
	StorageType *string `json:"storage_type,omitempty" xml:"storage_type"`

	// 消息老化策略。
	RetentionPolicy *ShowMqsInstanceResponseRetentionPolicy `json:"retention_policy,omitempty" xml:"retention_policy"`

	// Kafka公网开启状态。
	KafkaPublicStatus *string `json:"kafka_public_status,omitempty" xml:"kafka_public_status"`

	// 公网带宽。
	PublicBandwidth *int32 `json:"public_bandwidth,omitempty" xml:"public_bandwidth"`

	// 登录Kafka Manager的用户名。
	KafkaManagerUser *string `json:"kafka_manager_user,omitempty" xml:"kafka_manager_user"`

	// 是否开启消息收集功能。
	EnableLogCollection *bool `json:"enable_log_collection,omitempty" xml:"enable_log_collection"`

	// 跨VPC访问信息。
	CrossVpcInfo *string `json:"cross_vpc_info,omitempty" xml:"cross_vpc_info"`

	// 是否开启ipv6。
	Ipv6Enable *bool `json:"ipv6_enable,omitempty" xml:"ipv6_enable"`

	// IPv6的连接地址。
	Ipv6ConnectAddresses *[]string `json:"ipv6_connect_addresses,omitempty" xml:"ipv6_connect_addresses"`

	// 是否开启转储。
	ConnectorEnable *bool `json:"connector_enable,omitempty" xml:"connector_enable"`

	// 转储任务ID。
	ConnectorId *string `json:"connector_id,omitempty" xml:"connector_id"`

	// 是否开启MQS connector。
	MqsConnectorEnable *bool `json:"mqs_connector_enable,omitempty" xml:"mqs_connector_enable"`

	// MQS connector地址。
	MqsConnectorAddress *string `json:"mqs_connector_address,omitempty" xml:"mqs_connector_address"`

	// 是否开启插件功能。
	PluginEnable *bool `json:"plugin_enable,omitempty" xml:"plugin_enable"`

	// 是否开启Kafka rest功能。
	RestEnable *bool `json:"rest_enable,omitempty" xml:"rest_enable"`

	// Kafka rest地址。
	RestConnectAddress *string `json:"rest_connect_address,omitempty" xml:"rest_connect_address"`

	// 是否开启消息查询功能。
	MessageQueryInstEnable *bool `json:"message_query_inst_enable,omitempty" xml:"message_query_inst_enable"`

	// 是否开启VPC明文访问。
	VpcClientPlain *bool `json:"vpc_client_plain,omitempty" xml:"vpc_client_plain"`

	// Kafka实例支持的特性功能。
	SupportFeatures *string `json:"support_features,omitempty" xml:"support_features"`

	// 是否开启消息轨迹功能。
	TraceEnable *bool `json:"trace_enable,omitempty" xml:"trace_enable"`

	// 租户侧连接地址。
	PodConnectAddress *string `json:"pod_connect_address,omitempty" xml:"pod_connect_address"`

	// 是否开启磁盘加密。
	DiskEncrypted *bool `json:"disk_encrypted,omitempty" xml:"disk_encrypted"`

	// Kafka实例私有连接地址。
	KafkaPrivateConnectAddress *string `json:"kafka_private_connect_address,omitempty" xml:"kafka_private_connect_address"`

	// 云监控版本。
	CesVersion *string `json:"ces_version,omitempty" xml:"ces_version"`

	// 节点数量。
	NodeNum *int32 `json:"node_num,omitempty" xml:"node_num"`

	// 公网连接地址。
	PublicipAddress *string `json:"publicip_address,omitempty" xml:"publicip_address"`

	// 监听信息。
	Listeners *interface{} `json:"listeners,omitempty" xml:"listeners"`

	// 是否开启公网访问。用于区分何时开启的公网访问。
	PublicAccessEnabled *string `json:"public_access_enabled,omitempty" xml:"public_access_enabled"`

	// 公网访问带宽。
	PublicBoundwidth *int32 `json:"public_boundwidth,omitempty" xml:"public_boundwidth"`

	// 认证用户名。
	AccessUser *string `json:"access_user,omitempty" xml:"access_user"`

	// 是否开启代理。
	AgentEnable    *bool `json:"agent_enable,omitempty" xml:"agent_enable"`
	HttpStatusCode int   `json:"-"`
}

func (o ShowMqsInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowMqsInstanceResponse struct{}"
	}

	return strings.Join([]string{"ShowMqsInstanceResponse", string(data)}, " ")
}

type ShowMqsInstanceResponseType struct {
	value string
}

type ShowMqsInstanceResponseTypeEnum struct {
	SINGLE  ShowMqsInstanceResponseType
	CLUSTER ShowMqsInstanceResponseType
}

func GetShowMqsInstanceResponseTypeEnum() ShowMqsInstanceResponseTypeEnum {
	return ShowMqsInstanceResponseTypeEnum{
		SINGLE: ShowMqsInstanceResponseType{
			value: "single",
		},
		CLUSTER: ShowMqsInstanceResponseType{
			value: "cluster",
		},
	}
}

func (c ShowMqsInstanceResponseType) Value() string {
	return c.value
}

func (c ShowMqsInstanceResponseType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMqsInstanceResponseType) UnmarshalJSON(b []byte) error {
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

type ShowMqsInstanceResponseRetentionPolicy struct {
	value string
}

type ShowMqsInstanceResponseRetentionPolicyEnum struct {
	TIME_BASE      ShowMqsInstanceResponseRetentionPolicy
	PRODUCE_REJECT ShowMqsInstanceResponseRetentionPolicy
}

func GetShowMqsInstanceResponseRetentionPolicyEnum() ShowMqsInstanceResponseRetentionPolicyEnum {
	return ShowMqsInstanceResponseRetentionPolicyEnum{
		TIME_BASE: ShowMqsInstanceResponseRetentionPolicy{
			value: "time_base",
		},
		PRODUCE_REJECT: ShowMqsInstanceResponseRetentionPolicy{
			value: "produce_reject",
		},
	}
}

func (c ShowMqsInstanceResponseRetentionPolicy) Value() string {
	return c.value
}

func (c ShowMqsInstanceResponseRetentionPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowMqsInstanceResponseRetentionPolicy) UnmarshalJSON(b []byte) error {
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
