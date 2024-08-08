package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type AppServer struct {

	// aps实例的唯一标识。
	Id *string `json:"id,omitempty"`

	// 服务器名称。
	Name *string `json:"name,omitempty"`

	// 计算机名称。
	MachineName *string `json:"machine_name,omitempty"`

	// 描述。
	Description *string `json:"description,omitempty"`

	// 服务器组ID。
	ServerGroupId *string `json:"server_group_id,omitempty"`

	Flavor *Flavor `json:"flavor,omitempty"`

	Status *ServerStatus `json:"status,omitempty"`

	// 服务器创建时间。
	CreateTime *sdktime.SdkTime `json:"create_time,omitempty"`

	// 更新时间。
	UpdateTime *sdktime.SdkTime `json:"update_time,omitempty"`

	// 镜像ID。
	ImageId *string `json:"image_id,omitempty"`

	// 服务器可用分区。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// 域。
	Domain *string `json:"domain,omitempty"`

	// 组织名称。
	OuName *string `json:"ou_name,omitempty"`

	// 实例的SID。
	Sid *string `json:"sid,omitempty"`

	// 实例的ID。
	InstanceId *string `json:"instance_id,omitempty"`

	// 服务器系统版本。
	OsVersion *string `json:"os_version,omitempty"`

	// 操作系统类型，当前仅支持Windows： - Linux - Windows - Other
	OsType *string `json:"os_type,omitempty"`

	// 包周期产品的订单ID。
	OrderId *string `json:"order_id,omitempty"`

	// 是否维护状态。
	MaintainStatus *bool `json:"maintain_status,omitempty"`

	// 配置弹性伸缩策略时，服务自动创建的实例。 - true : 通过弹性伸缩创建。 - false: 不是通过弹性伸缩创建。
	ScalingAutoCreate *bool `json:"scaling_auto_create,omitempty"`

	// 上一次执行job的id。
	JobId *string `json:"job_id,omitempty"`

	JobType *JobType `json:"job_type,omitempty"`

	JobStatus *JobStatus `json:"job_status,omitempty"`

	// 上一次执行job的执行时间。
	JobTime *sdktime.SdkTime `json:"job_time,omitempty"`

	// 资源池ID。
	ResourcePoolId *string `json:"resource_pool_id,omitempty"`

	// 资源池类型： - private：私有资源池。 - public: 工作资源池。
	ResourcePoolType *string `json:"resource_pool_type,omitempty"`

	// 云专属主机id。
	HostId *string `json:"host_id,omitempty"`

	// 服务器组名称。
	ServerGroupName *string `json:"server_group_name,omitempty"`

	ProductInfo *ProductInfo `json:"product_info,omitempty"`

	// 弹性云服务器元数据。  >   1. charging_mode 云服务器的计费类型。  - “0”：按需计费（即postPaid-后付费方式）。 - “1”：按包年包月计费（即prePaid-预付费方式）。\"2\"：竞价实例计费  2. metering.order_id 按“包年/包月”计费的云服务器对应的订单ID。  3. metering.product_id 按“包年/包月”计费的云服务器对应的产品ID。  4. vpc_id 云服务器所属的虚拟私有云ID。  5. EcmResStatus 云服务器的冻结状态。  - normal：云服务器正常状态（未被冻结）。 - freeze：云服务器被冻结。  > 当云服务器被冻结或者解冻后，系统默认添加该字段，且该字段必选。  6. metering.image_id 云服务器操作系统对应的镜像ID  7.  metering.imagetype 镜像类型，目前支持：  - 公共镜像（gold） - 私有镜像（private） - 共享镜像（shared）  8. metering.resourcespeccode 云服务器对应的资源规格。  9. image_name 云服务器操作系统对应的镜像名称。  10. os_bit 操作系统位数，一般取值为“32”或者“64”。  11. lockCheckEndpoint 回调URL，用于检查弹性云服务器的加锁是否有效。  - 如果有效，则云服务器保持锁定状态。 - 如果无效，解除锁定状态，删除失效的锁。  12. lockSource 弹性云服务器来自哪个服务。订单加锁（ORDER）  13. lockSourceId 弹性云服务器的加锁来自哪个ID。lockSource为“ORDER”时，lockSourceId为订单ID。  14. lockScene 弹性云服务器的加锁类型。  - 按需转包周期（TO_PERIOD_LOCK）  15. virtual_env_type  - IOS镜像创建虚拟机，\"virtual_env_type\": \"IsoImage\" 属性； - 非IOS镜像创建虚拟机，在19.5.0版本以后创建的虚拟机将不会添加virtual_env_type 属性，而在此之前的版本创建的虚拟机可能会返回\"virtual_env_type\": \"FusionCompute\"属性 。  > virtual_env_type属性不允许用户增加、删除和修改。  16. metering.resourcetype 云服务器对应的资源类型。  17. os_type 操作系统类型，取值为：Linux、Windows。  18. cascaded.instance_extrainfo 系统内部虚拟机扩展信息。  19. __support_agent_list 云服务器是否支持企业主机安全、主机监控。  - “hss”：企业主机安全 -  “ces”：主机监控  20. agency_name 委托的名称。  委托是由租户管理员在统一身份认证服务（Identity and Access Management，IAM）上创建的，可以为弹性云服务器提供访问云服务的临时凭证。
	Metadata map[string]string `json:"metadata,omitempty"`

	// 会话数量。
	SessionCount *int32 `json:"session_count,omitempty"`

	VmStatus *AppServerStatus `json:"vm_status,omitempty"`

	TaskStatus *AppServerTaskStatus `json:"task_status,omitempty"`

	// 冻结信息。
	Freeze *[]CbcFreezeInfo `json:"freeze,omitempty"`

	// vpc和子网信息。
	HostAddress *[]EcsNetWork `json:"host_address,omitempty"`

	// 企业项目ID,仅企业项目会返回
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 标签信息
	Tags *[]TmsTag `json:"tags,omitempty"`
}

func (o AppServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AppServer struct{}"
	}

	return strings.Join([]string{"AppServer", string(data)}, " ")
}
