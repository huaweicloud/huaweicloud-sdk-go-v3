package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// 边缘实例对象。
type Instance struct {

	// 边缘实例ID。
	Id *string `json:"id,omitempty" xml:"id"`

	// 边缘实例名称。
	Name *string `json:"name,omitempty" xml:"name"`

	// 边缘实例状态。 取值范围： ACTIVE、BUILD、DELETED、ERROR、HARD_REBOOT、MIGRATING、PAUSED、REBOOT、REBUILD、RESIZE、REVERT_RESIZE、SHUTOFF、SHELVED、SHELVED_OFFLOADED、SOFT_DELETED、SUSPENDED、VERIFY_RESIZE
	Status *string `json:"status,omitempty" xml:"status"`

	// 边缘实例修改时间。 UTC时间，格式：yyyy-mm-ddTss:ss:ssZ，例如：2021-04-25T03:21:39Z
	Updated *string `json:"updated,omitempty" xml:"updated"`

	// 边缘实例所在主机的主机ID。
	HostId *string `json:"hostId,omitempty" xml:"hostId"`

	// 边缘实例对应的网络地址信息，详情请参见表addresses字段数据结构说明。
	Addresses map[string][]InstanceAddress `json:"addresses,omitempty" xml:"addresses"`

	// 边缘实例创建时间。 时间格式：yyyy-mm-ddTss:ss:ssZ，例如：2021-04-25T02:46:23Z
	Created *string `json:"created,omitempty" xml:"created"`

	// 边缘实例标签。 主要用来存储边缘业务ID。
	Tags *[]string `json:"tags,omitempty" xml:"tags"`

	// 边缘实例是否为锁定状态。  - true：锁定 - false：未锁定
	Locked *bool `json:"locked,omitempty" xml:"locked"`

	// 边缘实例的描述信息。
	Description *string `json:"description,omitempty" xml:"description"`

	// 边缘实例所属租户ID，即项目ID，和project_id表示相同的概念，格式为UUID。
	TenantId *string `json:"tenant_id,omitempty" xml:"tenant_id"`

	// 边缘实例系统标签。
	SysTags *[]SysTags `json:"sys_tags,omitempty" xml:"sys_tags"`

	Flavor *FlavorInstance `json:"flavor,omitempty" xml:"flavor"`

	// 边缘实例元数据。
	Metadata map[string]string `json:"metadata,omitempty" xml:"metadata"`

	// 边缘实例所属安全组列表。
	SecurityGroups *[]InstanceSecurityGroup `json:"security_groups,omitempty" xml:"security_groups"`

	// 边缘实例进度。
	Progress *int32 `json:"progress,omitempty" xml:"progress"`

	// 扩展属性，边缘实例电源状态。
	OSEXTSTSpowerState *int32 `json:"OS-EXT-STS:power_state,omitempty" xml:"OS-EXT-STS:power_state"`

	// 扩展属性，边缘实例当前状态。
	OSEXTSTSvmState *string `json:"OS-EXT-STS:vm_state,omitempty" xml:"OS-EXT-STS:vm_state"`

	// 边缘实例任务状态。
	OSEXTSTStaskState *string `json:"OS-EXT-STS:task_state,omitempty" xml:"OS-EXT-STS:task_state"`

	// 扩展属性， diskConfig的类型。  - MANUAL，镜像空间不会扩展。 - AUTO，系统盘镜像空间会自动扩展为与flavor大小一致。
	OSDCFdiskConfig *string `json:"OS-DCF:diskConfig,omitempty" xml:"OS-DCF:diskConfig"`

	// 扩展属性，边缘实例所在可用区名称。
	OSEXTAZavailabilityZone *string `json:"OS-EXT-AZ:availability_zone,omitempty" xml:"OS-EXT-AZ:availability_zone"`

	// 边缘实例启动时间。 时间格式例如：2019-05-22T03:23:59.000000
	OSSRVUSGlaunchedAt *string `json:"OS-SRV-USG:launched_at,omitempty" xml:"OS-SRV-USG:launched_at"`

	// 边缘实例删除时间。 时间格式例如：2019-05-22T03:23:59.000000
	OSSRVUSGterminatedAt *string `json:"OS-SRV-USG:terminated_at,omitempty" xml:"OS-SRV-USG:terminated_at"`

	// 边缘实例系统盘的设备名称。
	OSEXTSRVATTRrootDeviceName *string `json:"OS-EXT-SRV-ATTR:root_device_name,omitempty" xml:"OS-EXT-SRV-ATTR:root_device_name"`

	// 若使用AMI格式镜像，则表示ramdisk image的UUID；否则，留空。
	OSEXTSRVATTRramdiskId *string `json:"OS-EXT-SRV-ATTR:ramdisk_id,omitempty" xml:"OS-EXT-SRV-ATTR:ramdisk_id"`

	// 若使用AMI格式的镜像，则表示kernel image的UUID；否则，留空。
	OSEXTSRVATTRkernelId *string `json:"OS-EXT-SRV-ATTR:kernel_id,omitempty" xml:"OS-EXT-SRV-ATTR:kernel_id"`

	// 批量创建场景，边缘实例的启动顺序。
	OSEXTSRVATTRlaunchIndex *int32 `json:"OS-EXT-SRV-ATTR:launch_index,omitempty" xml:"OS-EXT-SRV-ATTR:launch_index"`

	// 批量创建场景，边缘实例的预留ID。
	OSEXTSRVATTRreservationId *string `json:"OS-EXT-SRV-ATTR:reservation_id,omitempty" xml:"OS-EXT-SRV-ATTR:reservation_id"`

	// 边缘实例的主机名。
	OSEXTSRVATTRhostname *string `json:"OS-EXT-SRV-ATTR:hostname,omitempty" xml:"OS-EXT-SRV-ATTR:hostname"`

	// 创建边缘实例时指定的user_data。
	OSEXTSRVATTRuserData *string `json:"OS-EXT-SRV-ATTR:user_data,omitempty" xml:"OS-EXT-SRV-ATTR:user_data"`

	// 边缘实例所在主机的主机名称。
	OSEXTSRVATTRhost *string `json:"OS-EXT-SRV-ATTR:host,omitempty" xml:"OS-EXT-SRV-ATTR:host"`

	// 扩展属性，边缘实例所在虚拟化主机名。
	OSEXTSRVATTRhypervisorHostname *string `json:"OS-EXT-SRV-ATTR:hypervisor_hostname,omitempty" xml:"OS-EXT-SRV-ATTR:hypervisor_hostname"`

	// 挂载到边缘实例上的磁盘。
	OsExtendedVolumesvolumesAttached *[]VolumesAttached `json:"os-extended-volumes:volumes_attached,omitempty" xml:"os-extended-volumes:volumes_attached"`

	Geolocation *GeoLocation `json:"geolocation,omitempty" xml:"geolocation"`

	// 边缘实例所属边缘业务的ID。
	EdgecloudId *string `json:"edgecloud_id,omitempty" xml:"edgecloud_id"`

	// 边缘实例所属边缘业务的名称
	EdgecloudName *string `json:"edgecloud_name,omitempty" xml:"edgecloud_name"`

	// 帐号ID。
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id"`

	// 使用的密钥对名称。
	KeyName *string `json:"key_name,omitempty" xml:"key_name"`

	// 扩展属性，边缘实例别名。
	OSEXTSRVATTRinstanceName *string `json:"OS-EXT-SRV-ATTR:instance_name,omitempty" xml:"OS-EXT-SRV-ATTR:instance_name"`
}

func (o Instance) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "Instance struct{}"
	}

	return strings.Join([]string{"Instance", string(data)}, " ")
}
