package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// Response Object
type ShowInstanceResponse struct {

	// VPC的名称。
	VpcName *string `json:"vpc_name,omitempty"`

	// 付费模式，0表示按需计费，1表示包年/包月计费。
	ChargingMode *int32 `json:"charging_mode,omitempty"`

	// VPC ID
	VpcId *string `json:"vpc_id,omitempty"`

	// 用户名。
	UserName *string `json:"user_name,omitempty"`

	// 完成创建时间。格式为：2017-03-31T12:24:46.297Z
	CreatedAt *string `json:"created_at,omitempty"`

	// 实例描述。
	Description *string `json:"description,omitempty"`

	// 安全组ID。
	SecurityGroupId *string `json:"security_group_id,omitempty"`

	// 租户安全组名称。
	SecurityGroupName *string `json:"security_group_name,omitempty"`

	// 总内存，单位：MB。
	MaxMemory *int32 `json:"max_memory,omitempty"`

	// 已使用的内存，单位：MB。
	UsedMemory *int32 `json:"used_memory,omitempty"`

	// 缓存实例的容量（G Byte）。
	Capacity *int32 `json:"capacity,omitempty"`

	// 单机小规格的缓存容量。
	CapacityMinor *string `json:"capacity_minor,omitempty"`

	// 维护时间窗开始时间，为UTC时间，格式为HH:mm:ss
	MaintainBegin *string `json:"maintain_begin,omitempty"`

	// 维护时间窗结束时间，为UTC时间，格式为HH:mm:ss
	MaintainEnd *string `json:"maintain_end,omitempty"`

	// 缓存实例的引擎类型。
	Engine *string `json:"engine,omitempty"`

	// 是否允许免密码访问缓存实例。 - true：该实例无需密码即可访问。 - false：该实例必须通过密码认证才能访问。
	NoPasswordAccess *string `json:"no_password_access,omitempty"`

	// 连接缓存实例的IP地址。如果是集群实例，返回多个IP地址，使用逗号分隔。如：192.168.0.1，192.168.0.2。
	Ip *string `json:"ip,omitempty"`

	InstanceBackupPolicy *InstanceBackupPolicy `json:"instance_backup_policy,omitempty"`

	// 实例所在的可用区。返回“可用区Code”
	AzCodes *[]string `json:"az_codes,omitempty"`

	// 通过密码认证访问缓存实例的认证用户名。
	AccessUser *string `json:"access_user,omitempty"`

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 缓存的端口。
	Port *int32 `json:"port,omitempty"`

	// 用户id。
	UserId *string `json:"user_id,omitempty"`

	// 实例名称。
	Name *string `json:"name,omitempty"`

	// 产品规格编码
	SpecCode *string `json:"spec_code,omitempty"`

	// 子网ID。
	SubnetId *string `json:"subnet_id,omitempty"`

	// 子网名称。
	SubnetName *string `json:"subnet_name,omitempty"`

	// 子网网段。
	SubnetCidr *string `json:"subnet_cidr,omitempty"`

	// 缓存版本。
	EngineVersion *string `json:"engine_version,omitempty"`

	// 订单ID。
	OrderId *string `json:"order_id,omitempty"`

	// 缓存实例的状态。详细状态说明见[缓存实例状态说明](https://support.huaweicloud.com/api-dcs/dcs-api-0312047.html)
	Status *string `json:"status,omitempty"`

	// 实例的域名。
	DomainName *string `json:"domain_name,omitempty"`

	// 实例的只读域名，只有主备实例有该字段。
	ReadonlyDomainName *string `json:"readonly_domain_name,omitempty"`

	// Redis缓存实例是否开启公网访问功能。 - true：开启 - false：不开启
	EnablePublicip *bool `json:"enable_publicip,omitempty"`

	// Redis缓存实例绑定的弹性IP地址的id。 如果未开启公网访问功能，该字段值为null。
	PublicipId *string `json:"publicip_id,omitempty"`

	// Redis缓存实例绑定的弹性IP地址。 如果未开启公网访问功能，该字段值为null。
	PublicipAddress *string `json:"publicip_address,omitempty"`

	// Redis缓存实例开启公网访问功能时，是否选择支持ssl。 - true：开启 - false：不开启
	EnableSsl *bool `json:"enable_ssl,omitempty"`

	// 实例是否存在升级任务。 - true：存在 - false：不存在
	ServiceUpgrade *bool `json:"service_upgrade,omitempty"`

	// 升级任务的ID。 - 当service_upgrade为true时，为升级任务的ID。 - 当service_upgrade为false时，该参数为空。
	ServiceTaskId *string `json:"service_task_id,omitempty"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 集群实例的后端服务地址。
	BackendAddrs *string `json:"backend_addrs,omitempty"`

	Features *Features `json:"features,omitempty"`

	DomainNameInfo *DomainNameInfo `json:"domain_name_info,omitempty"`

	// 是否开启客户端ip透传。
	TransparentClientIpEnable *bool `json:"transparent_client_ip_enable,omitempty"`

	// 实例子状态。
	SubStatus *string `json:"sub_status,omitempty"`

	// 实例标签键值。
	Tags *[]ResourceTag `json:"tags,omitempty"`

	// 实例CPU类型，通常为x86_64或aarch64
	CpuType *string `json:"cpu_type,omitempty"`

	// 企业项目名称。
	EnterpriseProjectName *string `json:"enterprise_project_name,omitempty"`

	// 更新时间，形如2022-07-06T09:32:16.502Z
	UpdateAt *string `json:"update_at,omitempty"`

	// 版本类型：社区版、企业版
	ProductType *ShowInstanceResponseProductType `json:"product_type,omitempty"`

	// 存储类型：内存存储
	StorageType *ShowInstanceResponseStorageType `json:"storage_type,omitempty"`

	// 启动时间，形如2022-07-06T09:32:16.502Z
	LaunchedAt *string `json:"launched_at,omitempty"`

	// 缓存类型：单机类型，主备类型，主备读写分离，Proxy集群类型，原生集群类型
	CacheMode *ShowInstanceResponseCacheMode `json:"cache_mode,omitempty"`

	// 是否支持慢日志
	SupportSlowLogFlag *string `json:"support_slow_log_flag,omitempty"`

	// 数据库数量
	DbNumber *int32 `json:"db_number,omitempty"`

	// 副本数
	ReplicaCount *int32 `json:"replica_count,omitempty"`

	// 集群实例分片个数
	ShardingCount *int32 `json:"sharding_count,omitempty"`

	BandwidthInfo  *BandwidthInfo `json:"bandwidth_info,omitempty"`
	HttpStatusCode int            `json:"-"`
}

func (o ShowInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceResponse", string(data)}, " ")
}

type ShowInstanceResponseProductType struct {
	value string
}

type ShowInstanceResponseProductTypeEnum struct {
	GENERIC    ShowInstanceResponseProductType
	ENTERPRISE ShowInstanceResponseProductType
}

func GetShowInstanceResponseProductTypeEnum() ShowInstanceResponseProductTypeEnum {
	return ShowInstanceResponseProductTypeEnum{
		GENERIC: ShowInstanceResponseProductType{
			value: "generic",
		},
		ENTERPRISE: ShowInstanceResponseProductType{
			value: "enterprise",
		},
	}
}

func (c ShowInstanceResponseProductType) Value() string {
	return c.value
}

func (c ShowInstanceResponseProductType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceResponseProductType) UnmarshalJSON(b []byte) error {
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

type ShowInstanceResponseStorageType struct {
	value string
}

type ShowInstanceResponseStorageTypeEnum struct {
	DRAM ShowInstanceResponseStorageType
}

func GetShowInstanceResponseStorageTypeEnum() ShowInstanceResponseStorageTypeEnum {
	return ShowInstanceResponseStorageTypeEnum{
		DRAM: ShowInstanceResponseStorageType{
			value: "DRAM",
		},
	}
}

func (c ShowInstanceResponseStorageType) Value() string {
	return c.value
}

func (c ShowInstanceResponseStorageType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceResponseStorageType) UnmarshalJSON(b []byte) error {
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

type ShowInstanceResponseCacheMode struct {
	value string
}

type ShowInstanceResponseCacheModeEnum struct {
	SINGLE      ShowInstanceResponseCacheMode
	HA          ShowInstanceResponseCacheMode
	HA_RW_SPLIT ShowInstanceResponseCacheMode
	PROXY       ShowInstanceResponseCacheMode
	CLUSTER     ShowInstanceResponseCacheMode
}

func GetShowInstanceResponseCacheModeEnum() ShowInstanceResponseCacheModeEnum {
	return ShowInstanceResponseCacheModeEnum{
		SINGLE: ShowInstanceResponseCacheMode{
			value: "single",
		},
		HA: ShowInstanceResponseCacheMode{
			value: "ha",
		},
		HA_RW_SPLIT: ShowInstanceResponseCacheMode{
			value: "ha_rw_split",
		},
		PROXY: ShowInstanceResponseCacheMode{
			value: "proxy",
		},
		CLUSTER: ShowInstanceResponseCacheMode{
			value: "cluster",
		},
	}
}

func (c ShowInstanceResponseCacheMode) Value() string {
	return c.value
}

func (c ShowInstanceResponseCacheMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *ShowInstanceResponseCacheMode) UnmarshalJSON(b []byte) error {
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
