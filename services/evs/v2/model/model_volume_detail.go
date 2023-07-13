package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// VolumeDetail 云硬盘详情。
type VolumeDetail struct {

	// 云硬盘的ID。
	Id string `json:"id"`

	// 云硬盘URI自描述信息。请参见 [links参数说明](https://support.huaweicloud.com/api-evs/evs_04_2006.html#evs_04_2006__evs_04_2010_li1077125119136)。
	Links []Link `json:"links"`

	// 云硬盘名称。
	Name string `json:"name"`

	// 云硬盘状态，请参见[云硬盘状态](https://support.huaweicloud.com/api-evs/evs_04_0040.html)。
	Status string `json:"status"`

	// 云硬盘的挂载信息，请参见•[attachments参数说明](https://support.huaweicloud.com/api-evs/evs_04_2006.html#evs_04_2006__evs_04_2010_li12430153610291)。
	Attachments []Attachment `json:"attachments"`

	// 云硬盘所属的AZ信息。
	AvailabilityZone string `json:"availability_zone"`

	// 预留属性。
	OsVolHostAttrhost string `json:"os-vol-host-attr:host"`

	// 源云硬盘ID，如果是从源云硬盘创建，则有值。  当前云硬盘服务不支持该字段。
	SourceVolid *string `json:"source_volid,omitempty"`

	// 快照ID，如果是从快照创建，则有值。
	SnapshotId string `json:"snapshot_id"`

	// 云硬盘描述。
	Description string `json:"description"`

	// 云硬盘创建时间。 时间格式：UTC YYYY-MM-DDTHH:MM:SS.XXXXXX
	CreatedAt string `json:"created_at"`

	// 云硬盘所属的租户ID。租户ID就是项目ID。
	OsVolTenantAttrtenantId string `json:"os-vol-tenant-attr:tenant_id"`

	// 云硬盘镜像的元数据。 > 说明： >  > 关于“volume_image_metadata”字段的详细说明，具体请参见：\"[查询镜像详情](https://support.huaweicloud.com/api-ims/ims_03_0703.html)\"。
	VolumeImageMetadata map[string]interface{} `json:"volume_image_metadata"`

	// 云硬盘类型。 目前支持\"SATA\"，\"SAS\"，\"GPSSD\"，\"SSD\"，\"ESSD\"，\"GPSSD2\",\"ESSD2\"七种。  - \"SATA\"为普通IO云硬盘 - \"SAS\"为高IO云硬盘 - \"GPSSD\"为通用型SSD云硬盘 - \"SSD\"为超高IO云硬盘 - \"ESSD\"为极速IO云硬盘 - \"GPSSD2\"为通用型SSD V2云硬盘 - \"ESSD2\"为极速型SSD V2云硬盘
	VolumeType string `json:"volume_type"`

	// 云硬盘大小，单位为GB。
	Size int32 `json:"size"`

	// 预留属性。
	ConsistencygroupId *string `json:"consistencygroup_id,omitempty"`

	// 是否为启动云硬盘。 true：表示为启动云硬盘。 false：表示为非启动云硬盘。
	Bootable string `json:"bootable"`

	// 云硬盘的元数据。  __system__cmkid metadata中的加密cmkid字段，与__system__encrypted配合表示需要加密，cmkid长度固定为36个字节。 >  > 请求获取密钥ID的方法请参考：\"[查询密钥列表](https://support.huaweicloud.com/api-dew/ListKeys.html)\"。  __system__encrypted metadata中的表示加密功能的字段，0代表不加密，1代表加密。 不指定该字段时，云硬盘的加密属性与数据源保持一致，如果不是从数据源创建的场景，则默认不加密。  full_clone 从快照创建云硬盘时的创建方式。 * 0表示使用链接克隆方式。 * 1表示使用全量克隆方式。  hw:passthrough * true表示云硬盘的设备类型为SCSI类型，即允许ECS操作系统直接访问底层存储介质。支持SCSI锁命令。 * false表示云硬盘的设备类型为VBD (虚拟块存储设备 , Virtual Block Device)类型，即为默认类型，VBD只能支持简单的SCSI读写命令。 * 该字段不存在时，云硬盘默认为VBD类型。  orderID metadata中的表示云硬盘计费类型的字段。 当该字段有值时，表示该云硬盘的计费类型为包周期计费，否则计费类型为按需计费。
	Metadata map[string]interface{} `json:"metadata"`

	// 云硬盘更新时间。 时间格式：UTC YYYY-MM-DDTHH:MM:SS.XXXXXX
	UpdatedAt string `json:"updated_at"`

	// 当前云硬盘服务不支持该字段。
	Encrypted *bool `json:"encrypted,omitempty"`

	// 预留属性。
	ReplicationStatus string `json:"replication_status"`

	// 预留属性。
	OsVolumeReplicationextendedStatus string `json:"os-volume-replication:extended_status"`

	// 预留属性。
	OsVolMigStatusAttrmigstat string `json:"os-vol-mig-status-attr:migstat"`

	// 预留属性。
	OsVolMigStatusAttrnameId string `json:"os-vol-mig-status-attr:name_id"`

	// 是否为共享云硬盘。true为共享盘，false为普通云硬盘。 该字段已经废弃，请使用multiattach。
	Shareable string `json:"shareable"`

	// 预留属性。
	UserId string `json:"user_id"`

	// 服务类型，结果为EVS、DSS、DESS。
	ServiceType string `json:"service_type"`

	// 是否为共享云硬盘。
	Multiattach bool `json:"multiattach"`

	// 云硬盘所属的专属存储池ID。
	DedicatedStorageId *string `json:"dedicated_storage_id,omitempty"`

	// 云硬盘所属的专属存储池的名称。
	DedicatedStorageName *string `json:"dedicated_storage_name,omitempty"`

	// 云硬盘的标签。 如果云硬盘有标签，则会有该字段，否则该字段为空。
	Tags map[string]string `json:"tags"`

	// 云硬盘挂载时的唯一标识。
	Wwn *string `json:"wwn,omitempty"`

	// 云硬盘上绑定的企业项目ID。 > 说明： >  > 关于企业项目ID的获取及企业项目特性的详细信息，请参见：\"[企业管理用户指南](https://support.huaweicloud.com/usermanual-em/zh-cn_topic_0123692049.html)\"。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 云硬盘序列号。 只有SCSI类型的非双活卷才会返回该字段，用于与虚拟机中的盘做对应关系。
	SerialNumber *string `json:"serial_number,omitempty"`

	Iops *Iops `json:"iops,omitempty"`

	Throughput *Throughput `json:"throughput,omitempty"`
}

func (o VolumeDetail) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "VolumeDetail struct{}"
	}

	return strings.Join([]string{"VolumeDetail", string(data)}, " ")
}
