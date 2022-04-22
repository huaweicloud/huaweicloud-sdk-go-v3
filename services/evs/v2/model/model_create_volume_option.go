package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// 创建云硬盘的信息。
type CreateVolumeOption struct {

	// 指定要创建云硬盘的可用区。 获取方法请参见\"[获取可用区](https://apiexplorer.developer.huaweicloud.com/apiexplorer/sdk?product=EVS&api=CinderListAvailabilityZones)\"。
	AvailabilityZone string `json:"availability_zone"`

	// 备份ID，从备份创建云硬盘时为必选。
	BackupId *string `json:"backup_id,omitempty"`

	// 批量创云硬盘的个数。如果无该参数，表明只创建1个云硬盘，目前最多支持批量创建100个。 从备份创建云硬盘时，不支持批量创建，数量只能为“1”。  如果发送请求时，将参数值设置为小数，则默认取小数点前的整数。
	Count *int32 `json:"count,omitempty"`

	// 云硬盘的描述。最大支持255个字节。
	Description *string `json:"description,omitempty"`

	// 企业项目ID。创建云硬盘时，给云硬盘绑定企业项目ID。
	EnterpriseProjectId *string `json:"enterprise_project_id,omitempty"`

	// 镜像ID，指定该参数表示创建云硬盘方式为从镜像创建云硬盘。
	ImageRef *string `json:"imageRef,omitempty"`

	// 创建云硬盘的metadata信息     可选参数如下:    [\\_\\_system\\_\\_cmkid]   metadata中的加密cmkid字段，与\\_\\_system\\_\\_encrypted配合表示需要加密，cmkid长度固定为36个字节。 > 说明： >  > 请求获取密钥ID的方法请参考：\"[查询密钥列表](https://support.huaweicloud.com/api-dew/ListKeys.html)\"。   [\\_\\_system\\_\\_encrypted]   metadata中的表示加密功能的字段,0代表不加密,1代表加密。不指定该字段时,云硬盘的加密属性与数据源保持一致,如果不是从数据源创建的场景,则默认不加密。    [full_clone]   从快照创建云硬盘时，如需使用link克隆方式，请指定该字段的值为0。    [hw:passthrough]    * true表示云硬盘的设备类型为SCSI类型，即允许ECS操作系统直接访问底层存储介质。支持SCSI锁命令。   * false表示云硬盘的设备类型为VBD (虚拟块存储设备 , Virtual Block Device)类型，即为默认类型，VBD只能支持简单的SCSI读写命令。   * 该字段不存在时，云硬盘默认为VBD类型。    [create\\_for\\_volume\\_id]  * true表示接口响应中会通过volume_ids字段返回本次创建的云硬盘ID。 * false表示接口响应中不会返回本次创建的云硬盘ID。  该字段不存在时，默认为false。
	Metadata map[string]string `json:"metadata,omitempty"`

	// 是否为共享云硬盘。true为共享盘，false为普通云硬盘。
	Multiattach *bool `json:"multiattach,omitempty"`

	// 云硬盘名称。 如果为创建单个云硬盘，name为云硬盘名称。最大支持255个字节。 创建的云硬盘数量（count字段对应的值）大于1时，为区分不同云硬盘，创建过程中系统会自动在名称后加“-0000”的类似标记。例如：volume-0001、volume-0002。最大支持250个字节。
	Name *string `json:"name,omitempty"`

	// 云硬盘大小，单位为GB，其限制如下： 系统盘：1GB-1024GB 数据盘：10GB-32768GB 创建空白云硬盘和从 镜像/快照 创建云硬盘时，size为必选，且云硬盘大小不能小于 镜像/快照 大小。 从备份创建云硬盘时，size为可选，不指定size时，云硬盘大小和备份大小一致。
	Size int32 `json:"size"`

	// 快照ID，指定该参数表示创建云硬盘方式为从快照创建云硬盘
	SnapshotId *string `json:"snapshot_id,omitempty"`

	// 云硬盘类型。  目前支持\"SATA\"，\"SAS\"，\"GPSSD\"，\"SSD\"和\"ESSD\"五种。  - \"SATA\"为普通IO云硬盘(已售罄) - \"SAS\"为高IO云硬盘 - \"GPSSD\"为通用型SSD云硬盘 - \"SSD\"为超高IO云硬盘 - \"ESSD\"为极速IO云硬盘 当指定的云硬盘类型在avaliability_zone内不存在时，则创建云硬盘失败。  说明： 从快照创建云硬盘时，volume_type字段必须和快照源云硬盘保持一致。 了解不同磁盘类型的详细信息，请参见 [磁盘类型及性能介绍](https://support.huaweicloud.com/productdesc-evs/zh-cn_topic_0044524691.html)。 获取region可用的卷类型，请参见[查询云硬盘类型列表](https://apiexplorer.developer.huaweicloud.com/apiexplorer/doc?product=EVS&api=CinderListVolumeTypes)
	VolumeType CreateVolumeOptionVolumeType `json:"volume_type"`

	// 云硬盘标签信息。
	Tags map[string]string `json:"tags,omitempty"`
}

func (o CreateVolumeOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "CreateVolumeOption struct{}"
	}

	return strings.Join([]string{"CreateVolumeOption", string(data)}, " ")
}

type CreateVolumeOptionVolumeType struct {
	value string
}

type CreateVolumeOptionVolumeTypeEnum struct {
	SSD   CreateVolumeOptionVolumeType
	GPSSD CreateVolumeOptionVolumeType
	SAS   CreateVolumeOptionVolumeType
	SATA  CreateVolumeOptionVolumeType
	ESSD  CreateVolumeOptionVolumeType
}

func GetCreateVolumeOptionVolumeTypeEnum() CreateVolumeOptionVolumeTypeEnum {
	return CreateVolumeOptionVolumeTypeEnum{
		SSD: CreateVolumeOptionVolumeType{
			value: "SSD",
		},
		GPSSD: CreateVolumeOptionVolumeType{
			value: "GPSSD",
		},
		SAS: CreateVolumeOptionVolumeType{
			value: "SAS",
		},
		SATA: CreateVolumeOptionVolumeType{
			value: "SATA",
		},
		ESSD: CreateVolumeOptionVolumeType{
			value: "ESSD",
		},
	}
}

func (c CreateVolumeOptionVolumeType) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *CreateVolumeOptionVolumeType) UnmarshalJSON(b []byte) error {
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
