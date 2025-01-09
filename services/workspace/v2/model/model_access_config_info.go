package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// AccessConfigInfo 接入配置信息。
type AccessConfigInfo struct {

	// 站点Id。
	SiteId string `json:"site_id"`

	// 备份模式是否开启。 - ON:开启。 - OFF:关闭。
	BackupMode AccessConfigInfoBackupMode `json:"backup_mode"`

	// 备份信息。
	BackupInfo *[]BackupInfo `json:"backup_info,omitempty"`
}

func (o AccessConfigInfo) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "AccessConfigInfo struct{}"
	}

	return strings.Join([]string{"AccessConfigInfo", string(data)}, " ")
}

type AccessConfigInfoBackupMode struct {
	value string
}

type AccessConfigInfoBackupModeEnum struct {
	ON  AccessConfigInfoBackupMode
	OFF AccessConfigInfoBackupMode
}

func GetAccessConfigInfoBackupModeEnum() AccessConfigInfoBackupModeEnum {
	return AccessConfigInfoBackupModeEnum{
		ON: AccessConfigInfoBackupMode{
			value: "ON",
		},
		OFF: AccessConfigInfoBackupMode{
			value: "OFF",
		},
	}
}

func (c AccessConfigInfoBackupMode) Value() string {
	return c.value
}

func (c AccessConfigInfoBackupMode) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *AccessConfigInfoBackupMode) UnmarshalJSON(b []byte) error {
	myConverter := converter.StringConverterFactory("string")
	if myConverter == nil {
		return errors.New("unsupported StringConverter type: string")
	}

	interf, err := myConverter.CovertStringToInterface(strings.Trim(string(b[:]), "\""))
	if err != nil {
		return err
	}

	if val, ok := interf.(string); ok {
		c.value = val
		return nil
	} else {
		return errors.New("convert enum data to string error")
	}
}
