package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// RecycleBinVolume 回收站云硬盘详情
type RecycleBinVolume struct {

	// **参数解释** 云硬盘ID。 **取值范围** 不涉及。
	Id *string `json:"id,omitempty"`

	// **参数解释** 云硬盘名称。 **取值范围** 不涉及。
	Name *string `json:"name,omitempty"`

	// **参数解释** 云硬盘描述。 **取值范围** 不涉及。
	Description *string `json:"description,omitempty"`

	// **参数解释** 云硬盘状态，具体请参见[云硬盘状态](evs_04_0040.xml)。 **取值范围** 不涉及。
	Status *string `json:"status,omitempty"`

	// **参数解释** 云硬盘的挂载信息。 **取值范围** 不涉及。
	Attachments *[]Attachment `json:"attachments,omitempty"`

	// **参数解释** 云硬盘是否共享。 **取值范围** - true：表示为共享云硬盘。 - false：表示为普通云硬盘。
	Multiattach *bool `json:"multiattach,omitempty"`

	// **参数解释** 云硬盘大小，单位为GiB。 **取值范围** 不涉及。
	Size *int32 `json:"size,omitempty"`

	// **参数解释** 云硬盘的metadata信息 ，调用方可以添加或删除metadata信息。 **取值范围** 当前部分key在EVS服务中有业务场景含义，这部分key的描述如下： - __system__cmkid    metadata中的加密cmkid字段，与__system__encrypted配合表示需要加密，cmkid长度固定为36个字节。    请求获取密钥ID的方法请参考：\"[查询密钥列表](https://support.huaweicloud.com/api-dew/ListKeys.html)\"。      - __system__encrypted    metadata中的表示加密功能的字段，0代表不加密，1代表加密。    不指定该字段时，云硬盘的加密属性与数据源保持一致，如果不是从数据源创建的场景，则默认不加密。   - hw:passthrough    - true表示云硬盘的设备类型为SCSI类型，即允许ECS操作系统直接访问底层存储介质。支持SCSI锁命令。    - false表示云硬盘的设备类型为VBD (虚拟块存储设备 , Virtual Block Device)类型，即为默认类型，VBD只能支持简单的SCSI读写命令。    - 该字段不存在时，云硬盘默认为VBD类型。
	Metadata map[string]interface{} `json:"metadata,omitempty"`

	// **参数解释** 云硬盘是否为启动盘。 **取值范围** - true：表示为启动云硬盘。 - false：表示为非启动云硬盘。
	Bootable *string `json:"bootable,omitempty"`

	// **参数解释** 云硬盘标签。 **取值范围** 不涉及。
	Tags map[string]string `json:"tags,omitempty"`

	// **参数解释** 云硬盘所属可用区。 **取值范围** 不涉及。
	AvailabilityZone *string `json:"availability_zone,omitempty"`

	// **参数解释** 云硬盘创建时间。 **取值范围** 时间格式：UTC YYYY-MM-DDTHH:MM:SS.XXXXXX
	CreatedAt *string `json:"created_at,omitempty"`

	// **参数解释** 云硬盘所属的服务类型。 **取值范围** - EVS：云硬盘。 - DSS：专属分布式存储服务。
	ServiceType *string `json:"service_type,omitempty"`

	// **参数解释** 云硬盘信息被更新的时间。 **取值范围** 时间格式：UTC YYYY-MM-DDTHH:MM:SS.XXXXXX
	UpdatedAt *string `json:"updated_at,omitempty"`

	// **参数解释** 云硬盘类型。 **取值范围** 目前支持\"SATA\"，\"SAS\"，\"GPSSD\"，\"SSD\"，\"ESSD\"，\"GPSSD2\"，\"ESSD2\"七种。   - \"SATA\"为普通IO云硬盘(已售罄) - \"SAS\"为高IO云硬盘 - \"GPSSD\"为通用型SSD云硬盘 - \"SSD\"为超高IO云硬盘 - \"ESSD\"为极速型SSD云硬盘 - \"GPSSD2\"为通用型SSD V2云硬盘 - \"ESSD2\"为极速型SSD V2云硬盘
	VolumeType *string `json:"volume_type,omitempty"`

	// **参数解释** 企业项目ID。 **取值范围** 不涉及。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// **参数解释** 预计到期清理的时间。 **取值范围** 时间格式：UTC YYYY-MM-DDTHH:MM:SS.XXXXXX
	PlanDeleteAt *string `json:"plan_delete_at,omitempty"`

	// **参数解释** 放入回收站的时间。 **取值范围** 时间格式：UTC YYYY-MM-DDTHH:MM:SS.XXXXXX
	PreDeletedAt *string `json:"pre_deleted_at,omitempty"`

	// **参数解释** 云硬盘所属的专属存储池ID。 **取值范围** 不涉及。
	DedicatedStorageId *string `json:"dedicated_storage_id,omitempty"`

	// **参数解释** 云硬盘所属的专属存储池的名称。 **取值范围** 不涉及。
	DedicatedStorageName *string `json:"dedicated_storage_name,omitempty"`

	// **参数解释** 云硬盘镜像的元数据。 关于“volume_image_metadata”字段的详细说明，具体请参见：\"[查询镜像详情](https://support.huaweicloud.com/api-ims/ims_03_0703.html)\"。  **取值范围** 不涉及。
	VolumeImageMetadata map[string]interface{} `json:"volume_image_metadata,omitempty"`

	// **参数解释** 云硬盘的唯一标识。 **取值范围** 不涉及。
	Wwn *string `json:"wwn,omitempty"`
}

func (o RecycleBinVolume) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RecycleBinVolume struct{}"
	}

	return strings.Join([]string{"RecycleBinVolume", string(data)}, " ")
}
