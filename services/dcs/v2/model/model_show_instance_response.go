package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// Response Object
type ShowInstanceResponse struct {

	// VPC的名称。
	VpcName *string `json:"vpc_name,omitempty" xml:"vpc_name"`

	// 付费模式，0表示按需计费，1表示包年/包月计费。
	ChargingMode *int32 `json:"charging_mode,omitempty" xml:"charging_mode"`

	// VPC ID
	VpcId *string `json:"vpc_id,omitempty" xml:"vpc_id"`

	// 用户名。
	UserName *string `json:"user_name,omitempty" xml:"user_name"`

	// 完成创建时间。格式为：2017-03-31T12:24:46.297Z
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at"`

	// 实例描述。
	Description *string `json:"description,omitempty" xml:"description"`

	// 安全组ID。
	SecurityGroupId *string `json:"security_group_id,omitempty" xml:"security_group_id"`

	// 租户安全组名称。
	SecurityGroupName *string `json:"security_group_name,omitempty" xml:"security_group_name"`

	// 总内存，单位：MB。
	MaxMemory *int32 `json:"max_memory,omitempty" xml:"max_memory"`

	// 已使用的内存，单位：MB。
	UsedMemory *int32 `json:"used_memory,omitempty" xml:"used_memory"`

	// 缓存实例的容量（G Byte）。
	Capacity *int32 `json:"capacity,omitempty" xml:"capacity"`

	// 单机小规格的缓存容量。
	CapacityMinor *string `json:"capacity_minor,omitempty" xml:"capacity_minor"`

	// 维护时间窗开始时间，为UTC时间，格式为HH:mm:ss
	MaintainBegin *string `json:"maintain_begin,omitempty" xml:"maintain_begin"`

	// 维护时间窗结束时间，为UTC时间，格式为HH:mm:ss
	MaintainEnd *string `json:"maintain_end,omitempty" xml:"maintain_end"`

	// 缓存实例的引擎类型。
	Engine *string `json:"engine,omitempty" xml:"engine"`

	// 是否允许免密码访问缓存实例。 - true：该实例无需密码即可访问。 - false：该实例必须通过密码认证才能访问。
	NoPasswordAccess *string `json:"no_password_access,omitempty" xml:"no_password_access"`

	// 连接缓存实例的IP地址。如果是集群实例，返回多个IP地址，使用逗号分隔。如：192.168.0.1，192.168.0.2。
	Ip *string `json:"ip,omitempty" xml:"ip"`

	InstanceBackupPolicy *InstanceBackupPolicy `json:"instance_backup_policy,omitempty" xml:"instance_backup_policy"`

	// 实例所在的可用区。返回“可用区Code”
	AzCodes *[]string `json:"az_codes,omitempty" xml:"az_codes"`

	// 通过密码认证访问缓存实例的认证用户名。
	AccessUser *string `json:"access_user,omitempty" xml:"access_user"`

	// 实例ID。
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id"`

	// 缓存的端口。
	Port *int32 `json:"port,omitempty" xml:"port"`

	// 用户id。
	UserId *string `json:"user_id,omitempty" xml:"user_id"`

	// 实例名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 产品规格编码
	SpecCode *string `json:"spec_code,omitempty" xml:"spec_code"`

	// 子网ID。
	SubnetId *string `json:"subnet_id,omitempty" xml:"subnet_id"`

	// 子网名称。
	SubnetName *string `json:"subnet_name,omitempty" xml:"subnet_name"`

	// 子网网段。
	SubnetCidr *string `json:"subnet_cidr,omitempty" xml:"subnet_cidr"`

	// 缓存版本。
	EngineVersion *string `json:"engine_version,omitempty" xml:"engine_version"`

	// 订单ID。
	OrderId *string `json:"order_id,omitempty" xml:"order_id"`

	// 缓存实例的状态。详细状态说明见[缓存实例状态说明](https://support.huaweicloud.com/api-dcs/dcs-api-0312047.html)
	Status *string `json:"status,omitempty" xml:"status"`

	// 实例的域名。
	DomainName *string `json:"domain_name,omitempty" xml:"domain_name"`

	// 实例的只读域名，只有主备实例有该字段。
	ReadonlyDomainName *string `json:"readonly_domain_name,omitempty" xml:"readonly_domain_name"`

	// Redis缓存实例是否开启公网访问功能。 - true：开启 - false：不开启
	EnablePublicip *bool `json:"enable_publicip,omitempty" xml:"enable_publicip"`

	// Redis缓存实例绑定的弹性IP地址的id。 如果未开启公网访问功能，该字段值为null。
	PublicipId *string `json:"publicip_id,omitempty" xml:"publicip_id"`

	// Redis缓存实例绑定的弹性IP地址。 如果未开启公网访问功能，该字段值为null。
	PublicipAddress *string `json:"publicip_address,omitempty" xml:"publicip_address"`

	// Redis缓存实例开启公网访问功能时，是否选择支持ssl。 - true：开启 - false：不开启
	EnableSsl *bool `json:"enable_ssl,omitempty" xml:"enable_ssl"`

	// 实例是否存在升级任务。 - true：存在 - false：不存在
	ServiceUpgrade *bool `json:"service_upgrade,omitempty" xml:"service_upgrade"`

	// 升级任务的ID。 - 当service_upgrade为true时，为升级任务的ID。 - 当service_upgrade为false时，该参数为空。
	ServiceTaskId *string `json:"service_task_id,omitempty" xml:"service_task_id"`

	// 企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty" xml:"enterprise_project_id"`

	// 集群实例的后端服务地址。
	BackendAddrs *string `json:"backend_addrs,omitempty" xml:"backend_addrs"`

	Features *Features `json:"features,omitempty" xml:"features"`

	DomainNameInfo *DomainNameInfo `json:"domain_name_info,omitempty" xml:"domain_name_info"`

	// 是否开启客户端ip透传。
	TransparentClientIpEnable *bool `json:"transparent_client_ip_enable,omitempty" xml:"transparent_client_ip_enable"`

	// 实例子状态。
	SubStatus *string `json:"sub_status,omitempty" xml:"sub_status"`

	// 实例标签键值。
	Tags *[]ResourceTag `json:"tags,omitempty" xml:"tags"`

	// 实例CPU类型，通常为x86_64或aarch64
	CpuType        *string `json:"cpu_type,omitempty" xml:"cpu_type"`
	HttpStatusCode int     `json:"-"`
}

func (o ShowInstanceResponse) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "ShowInstanceResponse struct{}"
	}

	return strings.Join([]string{"ShowInstanceResponse", string(data)}, " ")
}
