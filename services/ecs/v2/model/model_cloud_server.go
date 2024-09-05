package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

type CloudServer struct {

	// 云服务器唯一标识。
	Id string `json:"id"`

	// 云服务器名称。
	Name string `json:"name"`

	// 云服务器当前状态信息。
	Status string `json:"status"`

	// 云服务器所属租户ID。即项目id，与project_id表示相同的概念。
	TenantId string `json:"tenant_id"`

	// 云服务器所属用户ID。
	UserId string `json:"user_id"`

	MarketInfo *MarketModel `json:"market_info,omitempty"`

	// 可用分区
	AvailabilityZone string `json:"availability_zone"`

	// 云服务器的状态。
	VmState string `json:"vm_state"`

	// 云服务器任务状态。
	TaskState string `json:"task_state"`

	// 云服务器电源状态。
	PowerState *int32 `json:"power_state,omitempty"`

	// 云服务器创建时间。 时间格式例如：2020-05-22T07:48:53Z。
	Created string `json:"created"`

	// 云服务器是否处于回收站中
	InRecycleBin bool `json:"in_recycle_bin"`

	// 共池裸机按整机柜发放的同一批次的批创ID
	SpodId string `json:"spod_id"`

	// 云服务器上一次更新时间。时间格式例如：2020-05-22T07:48:53Z
	Updated string `json:"updated"`

	// 云服务器启动时间。时间格式例如：2020-05-22T07:48:53.000000。
	LaunchedAt *string `json:"launched_at,omitempty"`

	// 云服务器的描述信息。
	Description *string `json:"description,omitempty"`

	// 云服务器使用的密钥对名称。
	KeyName *string `json:"key_name,omitempty"`

	// 云服务器是否为锁定状态。  true：锁定 false：未锁定
	Locked *bool `json:"locked,omitempty"`

	// 云服务器系统盘的设备名称，例如当系统盘的磁盘模式是VDB时，为/dev/vda。
	RootDeviceName *string `json:"root_device_name,omitempty"`

	// 在专属主机或共享池中创建云服务器。默认为在共享池创建。值为： shared或dedicated。  shared：表示共享池。 dedicated:表示专属主机。
	Tenancy *string `json:"tenancy,omitempty"`

	// 专属主机ID。此属性仅在tenancy值为dedicated时有效，不指定此属性，系统将自动分配租户可自动放置云服务器的专属主机。
	DedicatedHostId *string `json:"dedicated_host_id,omitempty"`

	// 查询绑定某个企业项目的云服务器。 若需要查询当前用户所有企业项目绑定的云服务，请传参all_granted_eps。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 云服务器元数据。
	Metadata map[string]string `json:"metadata,omitempty"`

	// 云服务器的标签列表。
	Tags *[]string `json:"tags,omitempty"`

	// 云服务器对应的网络地址信息。
	Addresses map[string][]NetworkAddresses `json:"addresses,omitempty"`

	// 云服务器的安全组信息。
	SecurityGroups *[]SecurityGroup `json:"security_groups,omitempty"`

	// 云服务器挂载磁盘信息。
	VolumesAttached *[]VolumeAttach `json:"volumes_attached,omitempty"`

	Image *Image `json:"image,omitempty"`

	Flavor *FlavorQuasar `json:"flavor"`

	Fault *Fault `json:"fault"`

	CpuOptions *CpuOptions `json:"cpu_options,omitempty"`
}

func (o CloudServer) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CloudServer struct{}"
	}

	return strings.Join([]string{"CloudServer", string(data)}, " ")
}
