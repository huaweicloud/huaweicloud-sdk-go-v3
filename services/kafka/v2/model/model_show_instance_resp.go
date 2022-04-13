package model

import (
	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/utils"

	"code.byted.org/ti/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

type ShowInstanceResp struct {
	// 实例名称。

	Name *string `json:"name,omitempty"`
	// 引擎。

	Engine *string `json:"engine,omitempty"`
	// 版本。

	EngineVersion *string `json:"engine_version,omitempty"`
	// 实例规格。

	Specification *string `json:"specification,omitempty"`
	// 消息存储空间，单位：GB。

	StorageSpace *int32 `json:"storage_space,omitempty"`
	// Kafka实例的分区数量。

	PartitionNum *string `json:"partition_num,omitempty"`
	// 已使用的消息存储空间，单位：GB。

	UsedStorageSpace *int32 `json:"used_storage_space,omitempty"`
	// 实例连接IP地址。

	ConnectAddress *string `json:"connect_address,omitempty"`
	// 实例连接端口。

	Port *int32 `json:"port,omitempty"`
	// 实例的状态。

	Status *string `json:"status,omitempty"`
	// 实例ID。

	InstanceId *string `json:"instance_id,omitempty"`
	// 资源规格标识。   - dms.instance.kafka.cluster.c3.mini：Kafka实例的基准带宽为100MByte/秒。   - dms.instance.kafka.cluster.c3.small.2：Kafka实例的基准带宽为300MByte/秒。   - dms.instance.kafka.cluster.c3.middle.2：Kafka实例的基准带宽为600MByte/秒。   - dms.instance.kafka.cluster.c3.high.2：Kafka实例的基准带宽为1200MByte/秒。

	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`
	// 付费模式，1表示按需计费，0表示包年/包月计费。

	ChargingMode *int32 `json:"charging_mode,omitempty"`
	// VPC ID。

	VpcId *string `json:"vpc_id,omitempty"`
	// VPC的名称。

	VpcName *string `json:"vpc_name,omitempty"`
	// 完成创建时间。  格式为时间戳，指从格林威治时间 1970年01月01日00时00分00秒起至指定时间的偏差总毫秒数。

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
	// 实例是否开启公网访问功能。 - true：开启 - false：未开启

	EnablePublicip *bool `json:"enable_publicip,omitempty"`
	// Kafka实例的KafkaManager连接地址。

	ManagementConnectAddress *string `json:"management_connect_address,omitempty"`
	// 是否开启安全认证。 - true：开启 - false：未开启

	SslEnable *bool `json:"ssl_enable,omitempty"`
	// 企业项目ID。

	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`
	// 实例扩容时用于区分老实例与新实例。 - true：新创建的实例，允许磁盘动态扩容不需要重启。 - false：老实例

	IsLogicalVolume *bool `json:"is_logical_volume,omitempty"`
	// 实例扩容磁盘次数，如果超过20次则无法扩容磁盘。

	ExtendTimes *int32 `json:"extend_times,omitempty"`
	// 是否打开kafka自动创建topic功能。   - true：开启   - false：关闭

	EnableAutoTopic *bool `json:"enable_auto_topic,omitempty"`
	// 实例类型：集群，cluster。

	Type *ShowInstanceRespType `json:"type,omitempty"`
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
	// 总共消息存储空间，单位：GB。

	TotalStorageSpace *int32 `json:"total_storage_space,omitempty"`
	// 实例公网连接IP地址。当实例开启了公网访问，实例才包含该参数。

	PublicConnectAddress *string `json:"public_connect_address,omitempty"`
	// 存储资源ID。

	StorageResourceId *string `json:"storage_resource_id,omitempty"`
	// IO规格。

	StorageSpecCode *string `json:"storage_spec_code,omitempty"`
	// 服务类型。

	ServiceType *string `json:"service_type,omitempty"`
	// 存储类型。

	StorageType *string `json:"storage_type,omitempty"`
	// 消息老化策略。

	RetentionPolicy *ShowInstanceRespRetentionPolicy `json:"retention_policy,omitempty"`
	// Kafka公网开启状态。

	KafkaPublicStatus *string `json:"kafka_public_status,omitempty"`
	// 公网带宽。

	PublicBandwidth *int32 `json:"public_bandwidth,omitempty"`
	// 登录Kafka Manager的用户名。

	KafkaManagerUser *string `json:"kafka_manager_user,omitempty"`
	// 是否开启消息收集功能。

	EnableLogCollection *bool `json:"enable_log_collection,omitempty"`
	// 跨VPC访问信息。

	CrossVpcInfo *string `json:"cross_vpc_info,omitempty"`
	// 是否开启ipv6。

	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`
	// IPv6的连接地址。

	Ipv6ConnectAddresses *[]string `json:"ipv6_connect_addresses,omitempty"`
	// 是否开启转储。

	ConnectorEnable *bool `json:"connector_enable,omitempty"`
	// 转储任务ID。

	ConnectorId *string `json:"connector_id,omitempty"`
	// 是否开启Kafka rest功能。

	RestEnable *bool `json:"rest_enable,omitempty"`
	// Kafka rest连接地址。

	RestConnectAddress *string `json:"rest_connect_address,omitempty"`
	// 是否开启消息查询功能。

	MessageQueryInstEnable *bool `json:"message_query_inst_enable,omitempty"`
	// 是否开启VPC明文访问。

	VpcClientPlain *bool `json:"vpc_client_plain,omitempty"`
	// Kafka实例支持的特性功能。

	SupportFeatures *string `json:"support_features,omitempty"`
	// 是否开启消息轨迹功能。

	TraceEnable *bool `json:"trace_enable,omitempty"`
	// 租户侧连接地址。

	PodConnectAddress *string `json:"pod_connect_address,omitempty"`
	// 是否开启磁盘加密。

	DiskEncrypted *bool `json:"disk_encrypted,omitempty"`
	// Kafka实例私有连接地址。

	KafkaPrivateConnectAddress *string `json:"kafka_private_connect_address,omitempty"`
	// 云监控版本。

	CesVersion *string `json:"ces_version,omitempty"`
	// 标签列表。

	Tags *[]TagEntity `json:"tags,omitempty"`
	// 是否为容灾实例。

	DrEnable *bool `json:"dr_enable,omitempty"`
}

func (o ShowInstanceResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceResp struct{}"
	}

	return strings.Join([]string{"ShowInstanceResp", string(data)}, " ")
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

func (c ShowInstanceRespType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceRespType) UnmarshalJSON(b []byte) error {
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

type ShowInstanceRespRetentionPolicy struct {
	value string
}

type ShowInstanceRespRetentionPolicyEnum struct {
	TIME_BASE      ShowInstanceRespRetentionPolicy
	PRODUCE_REJECT ShowInstanceRespRetentionPolicy
}

func GetShowInstanceRespRetentionPolicyEnum() ShowInstanceRespRetentionPolicyEnum {
	return ShowInstanceRespRetentionPolicyEnum{
		TIME_BASE: ShowInstanceRespRetentionPolicy{
			value: "time_base",
		},
		PRODUCE_REJECT: ShowInstanceRespRetentionPolicy{
			value: "produce_reject",
		},
	}
}

func (c ShowInstanceRespRetentionPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceRespRetentionPolicy) UnmarshalJSON(b []byte) error {
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
