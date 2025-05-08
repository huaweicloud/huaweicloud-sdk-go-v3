package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type InstanceDetail struct {

	// **参数解释**： 实例名称。 **取值范围**： 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释**： 消息引擎。 **取值范围**： 不涉及。
	Engine *string `json:"engine,omitempty"`

	// **参数解释**： 状态。 **取值范围**： 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释**： 消息描述。 **取值范围**： 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释**： 实例类型。 **取值范围**： - single：单机。 - cluster：集群。
	Type *InstanceDetailType `json:"type,omitempty"`

	// **参数解释**： 实例规格。 **取值范围**： 不涉及。
	Specification *string `json:"specification,omitempty"`

	// **参数解释**： 实例版本。 **取值范围**： 不涉及。
	EngineVersion *string `json:"engine_version,omitempty"`

	// **参数解释**： 实例ID。 **取值范围**： 不涉及。
	InstanceId *string `json:"instance_id,omitempty"`

	// **参数解释**： 付费模式。 **取值范围**： [1表示按需计费。](tag:hws_eu,g42,hk_g42,tm,sbc,hk_sbc,hk_tm)[1表示按需计费，0表示包年/包月计费。](tag:hws,hws_eu,hws_hk,ctc) [计费模式，参数暂未使用。](tag:ocb,hws_ocb,hcs,fcs)
	ChargingMode *int32 `json:"charging_mode,omitempty"`

	// **参数解释**： 私有云ID。 **取值范围**： 不涉及。
	VpcId *string `json:"vpc_id,omitempty"`

	// **参数解释**： 私有云名称。 **取值范围**： 不涉及。
	VpcName *string `json:"vpc_name,omitempty"`

	// **参数解释**： 完成创建时间。  格式为时间戳，指从格林威治时间1970年01月01日00时00分00秒起至指定时间的偏差总毫秒数。 **取值范围**： 不涉及。
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释**： 产品标识。 **取值范围**： 不涉及。
	ProductId *string `json:"product_id,omitempty"`

	// **参数解释**： 安全组ID。 **取值范围**： 不涉及。
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// **参数解释**： 安全组名称。 **取值范围**： 不涉及。
	SecurityGroupName *string `json:"security_group_name,omitempty"`

	// **参数解释**： 子网ID。 **取值范围**： 不涉及。
	SubnetId *string `json:"subnet_id,omitempty"`

	// **参数解释**： 子网路由（仅RocketMQ 5.x版本会显示此字段）。 **取值范围**： 不涉及。
	SubnetCidr *string `json:"subnet_cidr,omitempty"`

	// **参数解释**： 可用区ID列表。 **取值范围**： 不涉及。
	AvailableZones *[]string `json:"available_zones,omitempty"`

	// **参数解释**： 可用区名称列表。
	AvailableZoneNames *[]string `json:"available_zone_names,omitempty"`

	// **参数解释**： 用户ID。 **取值范围**： 不涉及。
	UserId *string `json:"user_id,omitempty"`

	// **参数解释**： 用户名。 **取值范围**： 不涉及。
	UserName *string `json:"user_name,omitempty"`

	// **参数解释**： 维护时间窗开始时间，格式为HH:mm:ss。 **取值范围**： 不涉及。
	MaintainBegin *string `json:"maintain_begin,omitempty"`

	// **参数解释**： 维护时间窗结束时间，格式为HH:mm:ss。 **取值范围**： 不涉及。
	MaintainEnd *string `json:"maintain_end,omitempty"`

	// **参数解释**： 是否开启消息收集功能。 **取值范围**： - true：开启。 - false：不开启。
	EnableLogCollection *bool `json:"enable_log_collection,omitempty"`

	// **参数解释**： 存储空间，单位：GB。 **取值范围**： 不涉及。
	StorageSpace *int32 `json:"storage_space,omitempty"`

	// **参数解释**： 已用消息存储空间，单位：GB。 **取值范围**： 不涉及。
	UsedStorageSpace *int32 `json:"used_storage_space,omitempty"`

	// **参数解释**： 是否开启公网。 **取值范围**： - true：开启。 - false：不开启。
	EnablePublicip *bool `json:"enable_publicip,omitempty"`

	// **参数解释**： 实例绑定的弹性IP地址的ID。  以英文逗号隔开多个弹性IP地址的ID。  如果开启了公网访问功能（即enable_publicip为true），该字段为必选。 **取值范围**： 不涉及。
	PublicipId *string `json:"publicip_id,omitempty"`

	// **参数解释**： 公网IP地址。 **取值范围**： 不涉及。
	PublicipAddress *string `json:"publicip_address,omitempty"`

	// **参数解释**： 是否开启SSL。 **取值范围**： - true：开启。 - false：未开启。
	SslEnable *bool `json:"ssl_enable,omitempty"`

	// **参数解释**： 跨VPC访问信息。 **取值范围**： 不涉及。
	CrossVpcInfo *string `json:"cross_vpc_info,omitempty"`

	// **参数解释**： 存储资源ID。 **取值范围**： 不涉及。
	StorageResourceId *string `json:"storage_resource_id,omitempty"`

	// **参数解释**： 存储规格代码。 **取值范围**： 不涉及。
	StorageSpecCode *string `json:"storage_spec_code,omitempty"`

	// **参数解释**： 服务类型。 **取值范围**： 不涉及。
	ServiceType *string `json:"service_type,omitempty"`

	// **参数解释**： 存储类型。 **取值范围**： 不涉及。
	StorageType *string `json:"storage_type,omitempty"`

	// **参数解释**： 扩展时间。 **取值范围**： 不涉及。
	ExtendTimes *int64 `json:"extend_times,omitempty"`

	// **参数解释**： 是否开启IPv6。 **取值范围**： - true：开启。 - false：未开启。
	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	// **参数解释**： 实例支持的特性功能。 **取值范围**： 不涉及。
	SupportFeatures *string `json:"support_features,omitempty"`

	// **参数解释**： 是否开启磁盘加密。 **取值范围**： - true：开启。 - false：不开启。
	DiskEncrypted *bool `json:"disk_encrypted,omitempty"`

	// **参数解释**： 云监控版本。 **取值范围**： 不涉及。
	CesVersion *string `json:"ces_version,omitempty"`

	// **参数解释**： 节点数。 **取值范围**： 不涉及。
	NodeNum *int32 `json:"node_num,omitempty"`

	// **参数解释**： 是否启用新规格计费。 **取值范围**： - true：开启。 - false：未开启。
	NewSpecBillingEnable *bool `json:"new_spec_billing_enable,omitempty"`

	// **参数解释**： 是否开启访问控制列表。 **取值范围**： - true：开启。 - false：未开启。
	EnableAcl *bool `json:"enable_acl,omitempty"`

	// **参数解释**： Broker节点数（仅RocketMQ 4.8.0版本会显示此字段）。 **取值范围**： 不涉及。
	BrokerNum *int32 `json:"broker_num,omitempty"`

	// **参数解释**： 实例是否开启域名访问功能。 **取值范围**： - true：开启。 - false：未开启。
	DnsEnable *bool `json:"dns_enable,omitempty"`

	// **参数解释**： 元数据地址。 **取值范围**： 不涉及。
	NamesrvAddress *string `json:"namesrv_address,omitempty"`

	// **参数解释**： 元数据域名。 **取值范围**： 不涉及。
	NamesrvDomainName *string `json:"namesrv_domain_name,omitempty"`

	// **参数解释**： 业务数据地址。 **取值范围**： 不涉及。
	BrokerAddress *string `json:"broker_address,omitempty"`

	// **参数解释**： 公网元数据地址。 **取值范围**： 不涉及。
	PublicNamesrvAddress *string `json:"public_namesrv_address,omitempty"`

	// **参数解释**： 公网元数据域名。 **取值范围**： 不涉及。
	PublicNamesrvDomainName *string `json:"public_namesrv_domain_name,omitempty"`

	// **参数解释**： 公网业务数据地址。 **取值范围**： 不涉及。
	PublicBrokerAddress *string `json:"public_broker_address,omitempty"`

	// **参数解释**： grpc连接地址（仅RocketMQ 5.x版本会显示此字段）。 **取值范围**： 不涉及。
	GrpcAddress *string `json:"grpc_address,omitempty"`

	// **参数解释**： grpc连接域名（仅RocketMQ 5.x版本会显示此字段）。 **取值范围**： 不涉及。
	GrpcDomainName *string `json:"grpc_domain_name,omitempty"`

	// **参数解释**： 公网grpc连接地址（仅RocketMQ 5.x版本会显示此字段）。 **取值范围**： 不涉及。
	PublicGrpcAddress *string `json:"public_grpc_address,omitempty"`

	// **参数解释**： 公网grpc连接域名（仅RocketMQ 5.x版本会显示此字段）。 **取值范围**： 不涉及。
	PublicGrpcDomainName *string `json:"public_grpc_domain_name,omitempty"`

	// **参数解释**： 企业项目ID。 **取值范围**： 不涉及。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释**： 标签列表。 **取值范围**： 不涉及。
	Tags *[]TagEntity `json:"tags,omitempty"`

	// **参数解释**： 总存储空间。 **取值范围**： 不涉及。
	TotalStorageSpace *int32 `json:"total_storage_space,omitempty"`

	// **参数解释**： 资源规格。 **取值范围**： 不涉及。
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`

	// **参数解释**： 生产TPS占比。 **取值范围**： 不涉及。
	ProducePortion *int32 `json:"produce_portion,omitempty"`

	// **参数解释**： 消费TPS占比。 **取值范围**： 不涉及。
	ConsumePortion *int32 `json:"consume_portion,omitempty"`

	// **参数解释**： 是否为容灾实例。 **取值范围**： 不涉及。
	DrEnable *bool `json:"dr_enable,omitempty"`

	// **参数解释**： 配置ssl是否需要重启。 **取值范围**： 不涉及。
	ConfigSslNeedRestartProcess *bool `json:"config_ssl_need_restart_process,omitempty"`

	// **参数解释**： 实例使用的安全协议。 **取值范围**： 不涉及。
	TlsMode *string `json:"tls_mode,omitempty"`
}

func (o InstanceDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "InstanceDetail struct{}"
	}

	return strings.Join([]string{"InstanceDetail", string(data)}, " ")
}

type InstanceDetailType struct {
	value string
}

type InstanceDetailTypeEnum struct {
	SINGLE  InstanceDetailType
	CLUSTER InstanceDetailType
}

func GetInstanceDetailTypeEnum() InstanceDetailTypeEnum {
	return InstanceDetailTypeEnum{
		SINGLE: InstanceDetailType{
			value: "single",
		},
		CLUSTER: InstanceDetailType{
			value: "cluster",
		},
	}
}

func (c InstanceDetailType) Value() string {
	return c.value
}

func (c InstanceDetailType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *InstanceDetailType) UnmarshalJSON(b []byte) error {
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
