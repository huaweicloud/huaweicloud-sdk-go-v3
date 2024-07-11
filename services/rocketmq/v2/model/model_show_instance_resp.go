package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type ShowInstanceResp struct {

	// 实例名称。
	Name *string `json:"name,omitempty"`

	// 引擎。
	Engine *string `json:"engine,omitempty"`

	// 状态。
	Status *string `json:"status,omitempty"`

	// 消息描述。
	Description *string `json:"description,omitempty"`

	// 实例类型：集群，cluster。
	Type *ShowInstanceRespType `json:"type,omitempty"`

	// 实例规格。
	Specification *string `json:"specification,omitempty"`

	// 版本。
	EngineVersion *string `json:"engine_version,omitempty"`

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// [付费模式，1表示按需计费。](tag:hws_eu,g42,hk_g42,tm,sbc,hk_sbc,hk_tm)[付费模式，1表示按需计费，0表示包年/包月计费。](tag:hws,hws_eu,hws_hk,ctc) [计费模式，参数暂未使用。](tag:ocb,hws_ocb,hcs,fcs)
	ChargingMode *int32 `json:"charging_mode,omitempty"`

	// 私有云ID。
	VpcId *string `json:"vpc_id,omitempty"`

	// 私有云名称。
	VpcName *string `json:"vpc_name,omitempty"`

	// 完成创建时间。  格式为时间戳，指从格林威治时间1970年01月01日00时00分00秒起至指定时间的偏差总毫秒数。
	CreatedAt *string `json:"created_at,omitempty"`

	// 产品标识。
	ProductId *string `json:"product_id,omitempty"`

	// 安全组ID。
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// 租户安全组名称。
	SecurityGroupName *string `json:"security_group_name,omitempty"`

	// 子网ID。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 子网名称。
	SubnetName *string `json:"subnet_name,omitempty"`

	// 子网路由（仅RocketMQ 5.x版本会显示此字段）。
	SubnetCidr *string `json:"subnet_cidr,omitempty"`

	// 可用区ID列表。
	AvailableZones *[]string `json:"available_zones,omitempty"`

	// 可用区名称列表。
	AvailableZoneNames *[]string `json:"available_zone_names,omitempty"`

	// 用户ID。
	UserId *string `json:"user_id,omitempty"`

	// 用户名。
	UserName *string `json:"user_name,omitempty"`

	// 维护时间窗开始时间，格式为HH:mm:ss。
	MaintainBegin *string `json:"maintain_begin,omitempty"`

	// 维护时间窗结束时间，格式为HH:mm:ss。
	MaintainEnd *string `json:"maintain_end,omitempty"`

	// 是否开启消息收集功能。
	EnableLogCollection *bool `json:"enable_log_collection,omitempty"`

	// 存储空间，单位：GB。
	StorageSpace *int32 `json:"storage_space,omitempty"`

	// 已用消息存储空间，单位：GB。
	UsedStorageSpace *int32 `json:"used_storage_space,omitempty"`

	// 是否开启公网。
	EnablePublicip *bool `json:"enable_publicip,omitempty"`

	// 实例绑定的弹性IP地址的ID。 以英文逗号隔开多个弹性IP地址的ID。 如果开启了公网访问功能（即enable_publicip为true），该字段为必选。
	PublicipId *string `json:"publicip_id,omitempty"`

	// 公网IP地址。
	PublicipAddress *string `json:"publicip_address,omitempty"`

	// 是否开启SSL。
	SslEnable *bool `json:"ssl_enable,omitempty"`

	// 跨VPC访问信息。
	CrossVpcInfo *string `json:"cross_vpc_info,omitempty"`

	// 存储资源ID。
	StorageResourceId *string `json:"storage_resource_id,omitempty"`

	// 存储规格代码。
	StorageSpecCode *string `json:"storage_spec_code,omitempty"`

	// 服务类型。
	ServiceType *string `json:"service_type,omitempty"`

	// 存储类型。
	StorageType *string `json:"storage_type,omitempty"`

	// 扩展时间。
	ExtendTimes *int64 `json:"extend_times,omitempty"`

	// 是否开启IPv6。
	Ipv6Enable *bool `json:"ipv6_enable,omitempty"`

	// 实例支持的特性功能。
	SupportFeatures *string `json:"support_features,omitempty"`

	// 是否开启磁盘加密。
	DiskEncrypted *bool `json:"disk_encrypted,omitempty"`

	// 云监控版本。
	CesVersion *string `json:"ces_version,omitempty"`

	// 节点数。
	NodeNum *int32 `json:"node_num,omitempty"`

	// 是否启用新规格计费。
	NewSpecBillingEnable *bool `json:"new_spec_billing_enable,omitempty"`

	// 是否开启访问控制列表。
	EnableAcl *bool `json:"enable_acl,omitempty"`

	// 节点数（仅RocketMQ 4.8.0版本会显示此字段）。
	BrokerNum *int32 `json:"broker_num,omitempty"`

	// 元数据地址。
	NamesrvAddress *string `json:"namesrv_address,omitempty"`

	// 业务数据地址。
	BrokerAddress *string `json:"broker_address,omitempty"`

	// 公网元数据地址。
	PublicNamesrvAddress *string `json:"public_namesrv_address,omitempty"`

	// 公网业务数据地址。
	PublicBrokerAddress *string `json:"public_broker_address,omitempty"`

	// grpc连接地址（仅RocketMQ 5.x版本会显示此字段）。
	GrpcAddress *string `json:"grpc_address,omitempty"`

	// 公网grpc连接地址（仅RocketMQ 5.x版本会显示此字段）。
	PublicGrpcAddress *string `json:"public_grpc_address,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 标签列表。
	Tags *[]TagEntity `json:"tags,omitempty"`

	// 总存储空间。
	TotalStorageSpace *int32 `json:"total_storage_space,omitempty"`

	// 资源规格。
	ResourceSpecCode *string `json:"resource_spec_code,omitempty"`
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
