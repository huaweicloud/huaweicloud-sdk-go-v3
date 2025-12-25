package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// HwcEcs 云弹性云服务器
type HwcEcs struct {

	// 弹性云服务器ID，格式为UUID。
	Id string `json:"id"`

	// 弹性云服务器名称。
	Name string `json:"name"`

	// 主机安全开启状态：OPEN | CLOSE
	ProtectedStatus string `json:"protected_status"`

	// 弹性云服务器的描述信息。
	Description string `json:"description"`

	// 弹性云服务器状态。 取值范围： ACTIVE、BUILD、ERROR、HARD_REBOOT、MIGRATING、REBOOT、REBUILD、RESIZE、REVERT_RESIZE、SHUTOFF、VERIFY_RESIZE、DELETED
	Status string `json:"status"`

	// 弹性云服务器是否为锁定状态。 true：锁定 false：未锁定
	Locked bool `json:"locked"`

	// 弹性云服务器所属的企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 创建弹性云服务器的用户ID，格式为UUID。
	UserId string `json:"user_id"`

	// 弹性云服务器所属项目id，格式为UUID。
	ProjectId string `json:"project_id"`

	// 弹性云服务器所在主机的主机ID。
	HostId string `json:"host_id"`

	// 弹性云服务器所在主机的主机名称。
	HostName string `json:"host_name"`

	// 云服务器所在主机状态。 UP：服务正常 UNKNOWN：状态未知 DOWN：服务异常 MAINTENANCE：维护状态 空字符串：弹性云服务器无主机信息
	HostStatus string `json:"host_status"`

	// 弹性云服务器的网络属性。
	Addresses []HwcEcsAddress `json:"addresses"`

	// 弹性云服务器所属安全组列表。
	SecurityGroups []HwcEcsSecurityGroup `json:"security_groups"`

	// 弹性云服务器所在可用区名称。
	AvailabilityZone string `json:"availability_zone"`

	Flavor *HwcEcsFlavor `json:"flavor,omitempty"`

	// 挂载到弹性云服务器上的磁盘。
	VolumesAttached []HwcEcsVolume `json:"volumes_attached"`

	Metadata *HwcEcsMetadata `json:"metadata"`

	// 弹性云服务器最近一次更新时间，例如开机、关机、重启等操作。 时间格式例如：2019-05-22T03:30:52Z
	Updated string `json:"updated"`

	// 弹性云服务器创建时间。 时间格式例如：2019-05-22T03:19:19Z
	Created string `json:"created"`

	// 弹性云服务器使用的密钥对名称。
	KeyName *string `json:"key_name,omitempty"`

	SchedulerHints *HwcEcsSchedulerHint `json:"scheduler_hints,omitempty"`

	Hypervisor *HwcEcsHypervisor `json:"hypervisor,omitempty"`
}

func (o HwcEcs) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "HwcEcs struct{}"
	}

	return strings.Join([]string{"HwcEcs", string(data)}, " ")
}
