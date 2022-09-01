package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowInstanceResponse struct {

	// 实例名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 引擎。
	Engine *string `json:"engine,omitempty" xml:"engine"`

	// 状态。
	Status *string `json:"status,omitempty" xml:"status"`

	// 消息描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 实例类型：集群，cluster。
	Type *ShowInstanceResponseType `json:"type,omitempty" xml:"type"`

	// 实例规格。
	Specification *string `json:"specification,omitempty" xml:"specification"`

	// 版本。
	EngineVersion *string `json:"engine_version,omitempty" xml:"engine_version"`

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id"`

	// 付费模式，1表示按需计费，0表示包年/包月计费。
	ChargingMode *int32 `json:"charging_mode,omitempty" xml:"charging_mode"`

	// 私有云ID。
	VpcId *string `json:"vpc_id,omitempty" xml:"vpc_id"`

	// 私有云名称。
	VpcName *string `json:"vpc_name,omitempty" xml:"vpc_name"`

	// 完成创建时间。 格式为时间戳，指从格林威治时间1970年01月01日00时00分00秒起至指定时间的偏差总毫秒数。
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

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

	// 子网路由。
	SubnetCidr *string `json:"subnet_cidr,omitempty" xml:"subnet_cidr"`

	// IO未售罄的可用区列表。
	AvailableZones *[]string `json:"available_zones,omitempty" xml:"available_zones"`

	// 用户ID。
	UserId *string `json:"user_id,omitempty" xml:"user_id"`

	// 用户名。
	UserName *string `json:"user_name,omitempty" xml:"user_name"`

	// 维护时间窗开始时间，格式为HH:mm:ss。
	MaintainBegin *string `json:"maintain_begin,omitempty" xml:"maintain_begin"`

	// 维护时间窗结束时间，格式为HH:mm:ss。
	MaintainEnd *string `json:"maintain_end,omitempty" xml:"maintain_end"`

	// 是否开启消息收集功能。
	EnableLogCollection *bool `json:"enable_log_collection,omitempty" xml:"enable_log_collection"`

	// 存储空间，单位：GB。
	StorageSpace *int32 `json:"storage_space,omitempty" xml:"storage_space"`

	// 已用消息存储空间，单位：GB。
	UsedStorageSpace *int32 `json:"used_storage_space,omitempty" xml:"used_storage_space"`

	// 是否开启公网。
	EnablePublicip *bool `json:"enable_publicip,omitempty" xml:"enable_publicip"`

	// 实例绑定的弹性IP地址的ID。 以英文逗号隔开多个弹性IP地址的ID。 如果开启了公网访问功能（即enable_publicip为true），该字段为必选。
	PublicipId *string `json:"publicip_id,omitempty" xml:"publicip_id"`

	// 公网IP地址。
	PublicipAddress *string `json:"publicip_address,omitempty" xml:"publicip_address"`

	// 是否开启SSL。
	SslEnable *bool `json:"ssl_enable,omitempty" xml:"ssl_enable"`

	// 跨VPC访问信息。
	CrossVpcInfo *string `json:"cross_vpc_info,omitempty" xml:"cross_vpc_info"`

	// 存储资源ID。
	StorageResourceId *string `json:"storage_resource_id,omitempty" xml:"storage_resource_id"`

	// 存储规格代码。
	StorageSpecCode *string `json:"storage_spec_code,omitempty" xml:"storage_spec_code"`

	// 服务类型。
	ServiceType *string `json:"service_type,omitempty" xml:"service_type"`

	// 存储类型。
	StorageType *string `json:"storage_type,omitempty" xml:"storage_type"`

	// 扩展时间。
	ExtendTimes *int64 `json:"extend_times,omitempty" xml:"extend_times"`

	// 是否开启IPv6。
	Ipv6Enable *bool `json:"ipv6_enable,omitempty" xml:"ipv6_enable"`

	// 实例支持的特性功能。
	SupportFeatures *string `json:"support_features,omitempty" xml:"support_features"`

	// 是否开启磁盘加密。
	DiskEncrypted *bool `json:"disk_encrypted,omitempty" xml:"disk_encrypted"`

	// 云监控版本。
	CesVersion *string `json:"ces_version,omitempty" xml:"ces_version"`

	// 节点数。
	NodeNum *int32 `json:"node_num,omitempty" xml:"node_num"`

	// 是否启用新规格计费。
	NewSpecBillingEnable *bool `json:"new_spec_billing_enable,omitempty" xml:"new_spec_billing_enable"`

	// 是否开启访问控制列表。
	EnableAcl *bool `json:"enable_acl,omitempty" xml:"enable_acl"`

	// 节点数。
	BrokerNum *int32 `json:"broker_num,omitempty" xml:"broker_num"`

	// 元数据地址。
	NamesrvAddress *string `json:"namesrv_address,omitempty" xml:"namesrv_address"`

	// 业务数据地址。
	BrokerAddress *string `json:"broker_address,omitempty" xml:"broker_address"`

	// 公网元数据地址。
	PublicNamesrvAddress *string `json:"public_namesrv_address,omitempty" xml:"public_namesrv_address"`

	// 公网业务数据地址。
	PublicBrokerAddress *string `json:"public_broker_address,omitempty" xml:"public_broker_address"`

	// 标签列表。
	Tags *[]TagEntity `json:"tags,omitempty" xml:"tags"`

	// 总存储空间。
	TotalStorageSpace *int32 `json:"total_storage_space,omitempty" xml:"total_storage_space"`

	// 资源规格。
	ResourceSpecCode *string `json:"resource_spec_code,omitempty" xml:"resource_spec_code"`
	HttpStatusCode   int     `json:"-"`
}

func (o ShowInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceResponse", string(data)}, " ")
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
