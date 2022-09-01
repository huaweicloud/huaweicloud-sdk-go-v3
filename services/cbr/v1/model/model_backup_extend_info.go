package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

type BackupExtendInfo struct {

	// 是否是自动生成的备份副本
	AutoTrigger *bool `json:"auto_trigger,omitempty" xml:"auto_trigger"`

	// 是否系统盘备份
	Bootable *bool `json:"bootable,omitempty" xml:"bootable"`

	// 是否是增备
	Incremental *bool `json:"incremental,omitempty" xml:"incremental"`

	// 卷备份副本的快照id
	SnapshotId *string `json:"snapshot_id,omitempty" xml:"snapshot_id"`

	// 是否支持lazyloading快速恢复
	SupportLld *bool `json:"support_lld,omitempty" xml:"support_lld"`

	// 备份支持恢复的方式，当前取值包含na,snapshot和backup。如果该字段取值为snapshot，代表备份此时已经支持创建整机镜像；如果该字段取值为backup，备份支持通过云服务器上硬盘的备份进行恢复；如果该字段取值为na，备份不支持恢复。
	SupportedRestoreMode *BackupExtendInfoSupportedRestoreMode `json:"supported_restore_mode,omitempty" xml:"supported_restore_mode"`

	// 备份注册镜像ID列表
	OsImagesData *[]ImageData `json:"os_images_data,omitempty" xml:"os_images_data"`

	// 整机备份是否包含系统盘
	ContainSystemDisk *bool `json:"contain_system_disk,omitempty" xml:"contain_system_disk"`

	// 是否加密
	Encrypted *bool `json:"encrypted,omitempty" xml:"encrypted"`

	// 是否是系统盘
	SystemDisk *bool `json:"system_disk,omitempty" xml:"system_disk"`

	// 备份类型是否为多AZ
	IsMultiAz *bool `json:"is_multi_az,omitempty" xml:"is_multi_az"`
}

func (o BackupExtendInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupExtendInfo struct{}"
	}

	return strings.Join([]string{"BackupExtendInfo", string(data)}, " ")
}

type BackupExtendInfoSupportedRestoreMode struct {
	value string
}

type BackupExtendInfoSupportedRestoreModeEnum struct {
	NA       BackupExtendInfoSupportedRestoreMode
	BACKUP   BackupExtendInfoSupportedRestoreMode
	SNAPSHOT BackupExtendInfoSupportedRestoreMode
}

func GetBackupExtendInfoSupportedRestoreModeEnum() BackupExtendInfoSupportedRestoreModeEnum {
	return BackupExtendInfoSupportedRestoreModeEnum{
		NA: BackupExtendInfoSupportedRestoreMode{
			value: "na",
		},
		BACKUP: BackupExtendInfoSupportedRestoreMode{
			value: " backup",
		},
		SNAPSHOT: BackupExtendInfoSupportedRestoreMode{
			value: " snapshot",
		},
	}
}

func (c BackupExtendInfoSupportedRestoreMode) Value() string {
	return c.value
}

func (c BackupExtendInfoSupportedRestoreMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupExtendInfoSupportedRestoreMode) UnmarshalJSON(b []byte) error {
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
