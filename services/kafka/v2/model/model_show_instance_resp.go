package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ShowInstanceResp struct {

	// **参数解释**： 实例名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 引擎。 **取值范围**： kafka
	Engine *string `json:"engine,omitempty"`

	// **参数解释**： Kafka的版本。 **取值范围**： [- 1.1.0](tag:hws,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,hk_tm,sbc,hk_sbc,cmcc,ax,srg) [- 2.3.0](tag:g42,tm,hk_g42,ctc,hk_tm,dt,cmcc,ocb,hws_ocb) - 2.7 [- 3.x](tag:hws,hws_hk,dt,sbc,hk_sbc,hcs,fcs,ctc,tm,hk_tm,hws_eu,ax,cmcc,srg)
	EngineVersion *string `json:"engine_version,omitempty"`

	// **参数解释**： 实例描述。 **取值范围**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 实例规格。 **取值范围**： 不涉及。
	Specification *string `json:"specification,omitempty"`

	// **参数解释**： 消息存储空间，单位：GB。 **取值范围**： [- Kafka实例规格为c6.2u4g.cluster时，存储空间取值范围300GB ~ 300000GB。 - Kafka实例规格为c6.4u8g.cluster时，存储空间取值范围300GB ~ 600000GB。 - Kafka实例规格为c6.8u16g.cluster时，存储空间取值范围300GB ~ 1500000GB。 - Kafka实例规格为c6.12u24g.cluster时，存储空间取值范围300GB ~ 1500000GB。 - Kafka实例规格为c6.16u32g.cluster时，存储空间取值范围300GB ~ 1500000GB。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,hk_tm,dt,ax,cmcc,sbc,hk_sbc,srg) [- Kafka实例规格为kafka.2u4g.cluster.small时，存储空间取值范围300GB~300000GB。](tag:hws,hws_hk,hws_eu,dt,ax) [- Kafka实例规格为kafka.2u8g.cluster时，存储空间取值范围300GB~300000GB。](tag:fcs) [- Kafka实例规格为kafka.4u16g.cluster时，存储空间取值范围300GB~600000GB。 - Kafka实例规格为kafka.8u32g.cluster时，存储空间取值范围300GB~1500000GB。 - Kafka实例规格为kafka.16u64g.cluster时，存储空间取值范围300GB~1500000GB。 - Kafka实例规格为kafka.32u128g.cluster时，存储空间取值范围300GB~1500000GB。](tag:hcs,fcs)
	StorageSpace *int32 `json:"storage_space,omitempty"`

	// **参数解释**： Kafka实例的分区数量。 **取值范围**： 不涉及。
	PartitionNum *string `json:"partition_num,omitempty"`

	// **参数解释**： 已使用的消息存储空间，单位：GB。 **取值范围**： 不涉及。
	UsedStorageSpace *int32 `json:"used_storage_space,omitempty"`

	// **参数解释**： 实例是否开启域名访问功能。 **取值范围**： - true：开启 - false：未开启
	DnsEnable *bool `json:"dns_enable,omitempty"`

	// **参数解释**： 实例连接IP地址。 **取值范围**： 不涉及。
	ConnectAddress *string `json:"connect_address,omitempty"`

	// **参数解释**： 实例连接端口。 **取值范围**： 不涉及。
	Port *int32 `json:"port,omitempty"`

	// **参数解释**： 实例的状态。详细状态说明请参考[实例状态说明](kafka-api-180514012.xml)。 **取值范围**： 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 实例ID。 **取值范围**： 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**： 资源规格标识。 **取值范围**： [- dms.instance.kafka.cluster.c3.mini：Kafka实例的基准带宽为100MByte/秒。 - dms.instance.kafka.cluster.c3.small.2：Kafka实例的基准带宽为300MByte/秒。 - dms.instance.kafka.cluster.c3.middle.2：Kafka实例的基准带宽为600MByte/秒。 - dms.instance.kafka.cluster.c3.high.2：Kafka实例的基准带宽为1200MByte/秒。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,hk_tm,dt,ax)
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// **参数解释**： 付费模式。 **取值范围**： [- 1表示按需计费。 - 0表示包年/包月计费。](tag:hws,hws_hk,ctc,cmcc,ax,hws_eu)[付费模式，暂未使用。](tag:hws_ocb,ocb) [- 1表示按需计费。](tag:dt,g42,tm,hk_g42,hk_tm,hcs,fcs,sbc,hk_sbc)
	ChargingMode *int32 `json:"charging_mode,omitempty"`

	// **参数解释**： VPC ID。 **取值范围**： 不涉及。
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释**： VPC的名称。 **取值范围**： 不涉及。
	VpcName *string `json:"vpc_name,omitempty"`

	// **参数解释**： 完成创建时间。  格式为时间戳，指从格林威治时间 1970年01月01日00时00分00秒起至指定时间的偏差总毫秒数。 **取值范围**： 不涉及。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释**： 子网名称。 **取值范围**： 不涉及。
	SubnetName *string `json:"subnet_name,omitempty"`

	// **参数解释**： 子网网段。 **取值范围**： 不涉及。
	SubnetCidr *string `json:"subnet_cidr,omitempty"`

	// **参数解释**： 用户ID。 **取值范围**： 不涉及。
	UserId *string `json:"user_id,omitempty"`

	// **参数解释**： 用户名。 **取值范围**： 不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**： 实例访问用户名。 **取值范围**： 不涉及。
	AccessUser *string `json:"access_user,omitempty"`

	// **参数解释**： 订单ID，只有在包周期计费时才会有order_id值，其他计费方式order_id值为空。 **取值范围**： 不涉及。
	OrderId *string `json:"order_id,omitempty"`

	// **参数解释**： 维护时间窗开始时间，格式为HH:mm:ss。 **取值范围**： 不涉及。
	MaintainBegin *string `json:"maintain_begin,omitempty"`

	// **参数解释**： 维护时间窗结束时间，格式为HH:mm:ss。 **取值范围**： 不涉及。
	MaintainEnd *string `json:"maintain_end,omitempty"`

	// **参数解释**： 实例是否开启公网访问功能。 **取值范围**： - true：开启 - false：未开启
	EnablePublicip *bool `json:"enable_publicip,omitempty"`

	// **参数解释**： Kafka实例的Kafka Manager连接地址。 **取值范围**： 不涉及。
	ManagementConnectAddress *string `json:"management_connect_address,omitempty"`

	// **参数解释**： 是否开启安全认证。 **取值范围**： - true：开启 - false：未开启
	SslEnable *bool `json:"ssl_enable,omitempty"`

	// **参数解释**： 是否开启broker间副本加密传输。 **取值范围**： - true：开启 - false：未开启
	BrokerSslEnable *bool `json:"broker_ssl_enable,omitempty"`

	// **参数解释**： Kafka使用的安全协议。 若实例详情中不存在port_protocols返回参数，则kafka_security_protocol同时代表内网访问、公网访问以及跨VPC访问的安全协议。 若实例详情中存在port_protocols返回参数，则kafka_security_protocol仅代表跨VPC访问的安全协议。内网访问公网访问的安全协议请参考port_protocols参数。 **取值范围**： - PLAINTEXT：既未采用SSL证书进行加密传输，也不支持账号密码认证。性能更好，安全性较低，建议在生产环境下公网访问不使用此方式。 - SASL_SSL：采用SSL证书进行加密传输，支持账号密码认证，安全性更高。 [- SASL_PLAINTEXT：明文传输，支持账号密码认证，性能更好，建议使用SCRAM-SHA-512机制。](tag:hws,hws_eu,hws_hk,ocb,hws_ocb,ctc,g42,hk_g42,tm,hk_tm,dt,ax)
	KafkaSecurityProtocol *string `json:"kafka_security_protocol,omitempty"`

	// **参数解释**： 开启SASL后使用的认证机制。
	SaslEnabledMechanisms *[]ShowInstanceRespSaslEnabledMechanisms `json:"sasl_enabled_mechanisms,omitempty"`

	// **参数解释**： 是否开启双向认证。[华为云Stack不支持此参数。](tag:hcs) **取值范围**： - true：开启 - false：未开启
	SslTwoWayEnable *bool `json:"ssl_two_way_enable,omitempty"`

	// **参数解释**： 是否开启证书替换。[华为云Stack不支持此参数。](tag:hcs) **取值范围**： - true：开启 - false：未开启
	CertReplaced *bool `json:"cert_replaced,omitempty"`

	// **参数解释**： 公网访问Kafka Manager连接地址。[华为云Stack不支持此参数。](tag:hcs) **取值范围**： 不涉及。
	PublicManagementConnectAddress *string `json:"public_management_connect_address,omitempty"`

	// **参数解释**： 企业项目ID。 **取值范围**： 不涉及。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 实例扩容时用于区分老实例与新实例。 **取值范围**： - true：新创建的实例，允许磁盘动态扩容不需要重启。 - false：老实例。
	IsLogicalVolume *bool `json:"is_logical_volume,omitempty"`

	// **参数解释**： 实例扩容磁盘次数。 **取值范围**： 不涉及。
	ExtendTimes *int32 `json:"extend_times,omitempty"`

	// **参数解释**： 是否开启自动创建Topic。 **取值范围**： - true：开启 - false：关闭
	EnableAutoTopic *bool `json:"enable_auto_topic,omitempty"`

	// **参数解释**： 实例类型。 **取值范围**： - single：单机。 - cluster：集群。
	Type *ShowInstanceRespType `json:"type,omitempty"`

	// **参数解释**： 产品标识。 **取值范围**： 不涉及。
	ProductId *string `json:"product_id,omitempty"`

	// **参数解释**： 安全组ID。 **取值范围**： 不涉及。
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// **参数解释**： 安全组名称。 **取值范围**： 不涉及。
	SecurityGroupName *string `json:"security_group_name,omitempty"`

	// **参数解释**： 子网ID。 **取值范围**： 不涉及。
	SubnetId *string `json:"subnet_id,omitempty"`

	// **参数解释**： 实例节点所在的可用区，返回“可用区ID”。
	AvailableZones *[]string `json:"available_zones,omitempty"`

	// **参数解释**： 实例节点所在的可用区名称，返回“可用区名称”。
	AvailableZoneNames *[]string `json:"available_zone_names,omitempty"`

	// **参数解释**： 总共消息存储空间，单位：GB。 **取值范围**： 不涉及。
	TotalStorageSpace *int32 `json:"total_storage_space,omitempty"`

	// **参数解释**： 实例公网连接IP地址。当实例开启了公网访问，实例才包含该参数。 **取值范围**： 不涉及。
	PublicConnectAddress *string `json:"public_connect_address,omitempty"`

	// **参数解释**： 实例公网连接域名。当实例开启了公网访问，实例才包含该参数。 **取值范围**： 不涉及。
	PublicConnectDomainName *string `json:"public_connect_domain_name,omitempty"`

	// **参数解释**： 存储资源ID。 **取值范围**： 不涉及。
	StorageResourceId *string `json:"storage_resource_id,omitempty"`

	// **参数解释**： IO规格。 **取值范围**： 不涉及。
	StorageSpecCode *string `json:"storage_spec_code,omitempty"`

	// **参数解释**： 服务类型。 **取值范围**： advanced：服务类型。
	ServiceType *string `json:"service_type,omitempty"`

	// **参数解释**： 存储类型。 **取值范围**： hec：存储类型。
	StorageType *string `json:"storage_type,omitempty"`

	// **参数解释**： 消息老化策略。 **取值范围**： - time_base：表示自动删除最老消息。 - produce_reject：表示拒绝消息写入。
	RetentionPolicy *ShowInstanceRespRetentionPolicy `json:"retention_policy,omitempty"`

	// **参数解释**： Kafka公网开启状态。 **取值范围**： - true：开启公网。 - closed：关闭公网。 - false：未使用公网。 - freezed：公网冻结。 - actived：公网解冻。
	KafkaPublicStatus *string `json:"kafka_public_status,omitempty"`

	// **参数解释**： kafka公网访问带宽。 **取值范围**： 不涉及。
	PublicBandwidth *int32 `json:"public_bandwidth,omitempty"`

	// **参数解释**： 是否开启消息收集功能。 **取值范围**： - true：开启 - false：不开启
	EnableLogCollection *bool `json:"enable_log_collection,omitempty"`

	// **参数解释**： 是否开启新证书。 **取值范围**： - true：开启 - false：不开启
	NewAuthCert *bool `json:"new_auth_cert,omitempty"`

	// **参数解释**： 跨VPC访问信息。 **取值范围**： 不涉及。
	CrossVpcInfo *interface{} `json:"cross_vpc_info,omitempty"`

	// **参数解释**： 公网跨VPC访问信息。 **取值范围**： 不涉及。
	PublicCrossVpcInfo *interface{} `json:"public_cross_vpc_info,omitempty"`

	// **参数解释**： 是否开启IPv6。 **取值范围**： - true：开启 - false：不开启
	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	// **参数解释**： IPv6的连接地址。
	Ipv6ConnectAddresses *[]string `json:"ipv6_connect_addresses,omitempty"`

	// **参数解释**： 是否开启转储。 **取值范围**： - true：开启 - false：不开启
	ConnectorEnable *bool `json:"connector_enable,omitempty"`

	// **参数解释**： connector节点数量。 **取值范围**： 2-16。
	ConnectorNodeNum *int32 `json:"connector_node_num,omitempty"`

	// **参数解释**： 转储任务ID。 **取值范围**： 不涉及。
	ConnectorId *string `json:"connector_id,omitempty"`

	// **参数解释**： 是否开启Kafka rest功能。 **取值范围**： - true：开启 - false：不开启
	RestEnable *bool `json:"rest_enable,omitempty"`

	// **参数解释**：  Kafka rest连接地址。 **取值范围**： 不涉及。
	RestConnectAddress *string `json:"rest_connect_address,omitempty"`

	// **参数解释**： Kafka公网访问带宽。待删除版本。 **取值范围**： 不涉及。
	PublicBoundwidth *int32 `json:"public_boundwidth,omitempty"`

	// **参数解释**： 是否开启消息查询功能。 **取值范围**： - true：开启 - false：不开启
	MessageQueryInstEnable *bool `json:"message_query_inst_enable,omitempty"`

	// **参数解释**： 是否开启VPC明文访问。 **取值范围**： - true：开启 - false：不开启
	VpcClientPlain *bool `json:"vpc_client_plain,omitempty"`

	// **参数解释**： Kafka实例支持的特性功能。 **取值范围**： 不涉及。
	SupportFeatures *string `json:"support_features,omitempty"`

	// **参数解释**： 是否开启消息轨迹功能。 **取值范围**： - true：开启 - false：不开启
	TraceEnable *bool `json:"trace_enable,omitempty"`

	// **参数解释**： 是否开启代理。 **取值范围**： - true：开启 - false：不开启
	AgentEnable *bool `json:"agent_enable,omitempty"`

	// **参数解释**： 租户侧连接地址。 **取值范围**： 不涉及。
	PodConnectAddress *string `json:"pod_connect_address,omitempty"`

	// **参数解释**： 是否开启磁盘加密。 **取值范围**： - true：开启 - false：不开启
	DiskEncrypted *bool `json:"disk_encrypted,omitempty"`

	// **参数解释**： 磁盘加密key，未开启磁盘加密时为空。 **取值范围**： 不涉及。
	DiskEncryptedKey *string `json:"disk_encrypted_key,omitempty"`

	// **参数解释**： Kafka实例内网连接地址。 **取值范围**： 不涉及。
	KafkaPrivateConnectAddress *string `json:"kafka_private_connect_address,omitempty"`

	// **参数解释**： Kafka实例内网连接域名。 **取值范围**： 不涉及。
	KafkaPrivateConnectDomainName *string `json:"kafka_private_connect_domain_name,omitempty"`

	// **参数解释**： 云监控版本。 **取值范围**： 不涉及。
	CesVersion *string `json:"ces_version,omitempty"`

	// **参数解释**： 区分实例什么时候开启的公网访问 **取值范围**： - true：已开启公网访问 - actived：已开启公网访问 - closed：已关闭公网访问 - false：已关闭公网访问
	PublicAccessEnabled *string `json:"public_access_enabled,omitempty"`

	// **参数解释**： 节点数。 **取值范围**： [- 1：Kafka单机实例的节点数。](tag:hws,hws_hk,hws_eu,dt,hcs,ax) - 3~50：Kafka集群实例的节点数。
	NodeNum *int32 `json:"node_num,omitempty"`

	PortProtocols *PortProtocolsEntity `json:"port_protocols,omitempty"`

	// **参数解释**： 是否开启访问控制。 **取值范围**： - true：开启 - false：不开启
	EnableAcl *bool `json:"enable_acl,omitempty"`

	// **参数解释**： 是否启用新规格计费。 **取值范围**： - true：开启 - false：不开启
	NewSpecBillingEnable *bool `json:"new_spec_billing_enable,omitempty"`

	// **参数解释**： 节点数量。 **取值范围**： 不涉及。
	BrokerNum *int32 `json:"broker_num,omitempty"`

	// **参数解释**： 标签列表。
	Tags *[]TagEntity `json:"tags,omitempty"`

	// **参数解释**：  是否为容灾实例。 **取值范围**： - true：是容灾实例。 - false：不是容灾实例。
	DrEnable *bool `json:"dr_enable,omitempty"`
}

func (o ShowInstanceResp) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceResp struct{}"
	}

	return strings.Join([]string{"ShowInstanceResp", string(data)}, " ")
}

type ShowInstanceRespSaslEnabledMechanisms struct {
	value string
}

type ShowInstanceRespSaslEnabledMechanismsEnum struct {
	PLAIN         ShowInstanceRespSaslEnabledMechanisms
	SCRAM_SHA_512 ShowInstanceRespSaslEnabledMechanisms
}

func GetShowInstanceRespSaslEnabledMechanismsEnum() ShowInstanceRespSaslEnabledMechanismsEnum {
	return ShowInstanceRespSaslEnabledMechanismsEnum{
		PLAIN: ShowInstanceRespSaslEnabledMechanisms{
			value: "PLAIN",
		},
		SCRAM_SHA_512: ShowInstanceRespSaslEnabledMechanisms{
			value: "SCRAM-SHA-512",
		},
	}
}

func (c ShowInstanceRespSaslEnabledMechanisms) Value() string {
	return c.value
}

func (c ShowInstanceRespSaslEnabledMechanisms) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceRespSaslEnabledMechanisms) UnmarshalJSON(b []byte) error {
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

func (c ShowInstanceRespRetentionPolicy) Value() string {
	return c.value
}

func (c ShowInstanceRespRetentionPolicy) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceRespRetentionPolicy) UnmarshalJSON(b []byte) error {
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
