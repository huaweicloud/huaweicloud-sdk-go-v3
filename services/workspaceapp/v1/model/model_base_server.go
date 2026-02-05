package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/sdktime"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type BaseServer struct {

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

	// 服务器维护状态： - true : 维护态的实例。 - false: 非维护态的实例。
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
}

func (o BaseServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BaseServer struct{}"
	}

	return strings.Join([]string{"BaseServer", string(data)}, " ")
}
