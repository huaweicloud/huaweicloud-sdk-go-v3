package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

//
type AttachServerVolumeOption struct {

	// 磁盘挂载点。  > 说明：  - 新增加的磁盘挂载点不能和已有的磁盘挂载点相同。  - 对于采用XEN虚拟化类型的弹性云服务器，device为必选参数；系统盘挂载点请指定/dev/sda；数据盘挂载点请按英文字母顺序依次指定，如/dev/sdb，/dev/sdc，如果指定了以“/dev/vd”开头的挂载点，系统默认改为“/dev/sd”。  - 对于采用KVM虚拟化类型的弹性云服务器，系统盘挂载点请指定/dev/vda；数据盘挂载点可不用指定，也可按英文字母顺序依次指定，如/dev/vdb，/dev/vdc，如果指定了以“/dev/sd”开头的挂载点，系统默认改为“/dev/vd”。
	Device *string `json:"device,omitempty" xml:"device"`

	// 待挂载磁盘的磁盘ID，UUID格式。
	VolumeId string `json:"volumeId" xml:"volumeId"`

	// 云硬盘类型。  该字段于dry_run为true并且volumeId不存在时有效且为必选字段。
	VolumeType *string `json:"volume_type,omitempty" xml:"volume_type"`

	// - true: 表示云硬盘的设备类型为SCSI类型，即允许ECS操作系统直接访问底层存储介质。支持SCSI锁命令 - false: 表示云硬盘的设备类型为VBD (虚拟块存储设备 , Virtual Block Device)类型，VBD只能支持简单的SCSI读写命令。 该字段于dry_run为true并且volumeId不存在时有效且为必选字段。
	Hwpassthrough *string `json:"hw:passthrough,omitempty" xml:"hw:passthrough"`
}

func (o AttachServerVolumeOption) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AttachServerVolumeOption struct{}"
	}

	return strings.Join([]string{"AttachServerVolumeOption", string(data)}, " ")
}
