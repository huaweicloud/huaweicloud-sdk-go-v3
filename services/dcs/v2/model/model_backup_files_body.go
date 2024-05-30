package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"errors"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"

	"strings"
)

// BackupFilesBody 备份文件列表的结构体
type BackupFilesBody struct {

	// 数据来源，当前仅支持OBS桶方式，取值为：self_build_obs。
	FileSource *BackupFilesBodyFileSource `json:"file_source,omitempty"`

	// OBS桶名。
	BucketName *string `json:"bucket_name,omitempty"`

	// 导入的备份文件文件列表。
	Files []Files `json:"files"`

	// 备份记录ID，数据来源为备份记录时必须填写
	BackupId *string `json:"backup_id,omitempty"`
}

func (o BackupFilesBody) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "BackupFilesBody struct{}"
	}

	return strings.Join([]string{"BackupFilesBody", string(data)}, " ")
}

type BackupFilesBodyFileSource struct {
	value string
}

type BackupFilesBodyFileSourceEnum struct {
	SELF_BUILD_OBS BackupFilesBodyFileSource
	BACKUP_RECORD  BackupFilesBodyFileSource
}

func GetBackupFilesBodyFileSourceEnum() BackupFilesBodyFileSourceEnum {
	return BackupFilesBodyFileSourceEnum{
		SELF_BUILD_OBS: BackupFilesBodyFileSource{
			value: "self_build_obs",
		},
		BACKUP_RECORD: BackupFilesBodyFileSource{
			value: "backup_record",
		},
	}
}

func (c BackupFilesBodyFileSource) Value() string {
	return c.value
}

func (c BackupFilesBodyFileSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupFilesBodyFileSource) UnmarshalJSON(b []byte) error {
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
