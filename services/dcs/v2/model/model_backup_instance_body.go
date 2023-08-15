package model

import (
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/dysodeng/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BackupInstanceBody 备份实例请求体
type BackupInstanceBody struct {

	// 备份缓存实例的备注信息。
	Remark *string `json:"remark,omitempty"`

	// 备份缓存实例的格式。
	BackupFormat *BackupInstanceBodyBackupFormat `json:"backup_format,omitempty"`
}

func (o BackupInstanceBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupInstanceBody struct{}"
	}

	return strings.Join([]string{"BackupInstanceBody", string(data)}, " ")
}

type BackupInstanceBodyBackupFormat struct {
	value string
}

type BackupInstanceBodyBackupFormatEnum struct {
	AOF BackupInstanceBodyBackupFormat
	RDB BackupInstanceBodyBackupFormat
}

func GetBackupInstanceBodyBackupFormatEnum() BackupInstanceBodyBackupFormatEnum {
	return BackupInstanceBodyBackupFormatEnum{
		AOF: BackupInstanceBodyBackupFormat{
			value: "aof",
		},
		RDB: BackupInstanceBodyBackupFormat{
			value: "rdb",
		},
	}
}

func (c BackupInstanceBodyBackupFormat) Value() string {
	return c.value
}

func (c BackupInstanceBodyBackupFormat) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupInstanceBodyBackupFormat) UnmarshalJSON(b []byte) error {
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
