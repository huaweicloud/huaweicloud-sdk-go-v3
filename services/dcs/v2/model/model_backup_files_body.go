package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/converter"
	"errors"

	"strings"
)

// 备份文件列表的结构体
type BackupFilesBody struct {
	// 数据来源，当前仅支持OBS桶方式，取值为：self_build_obs。

	FileSource *BackupFilesBodyFileSource `json:"file_source,omitempty"`
	// OBS桶名。

	BucketName string `json:"bucket_name"`
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
}

func GetBackupFilesBodyFileSourceEnum() BackupFilesBodyFileSourceEnum {
	return BackupFilesBodyFileSourceEnum{
		SELF_BUILD_OBS: BackupFilesBodyFileSource{
			value: "self_build_obs",
		},
	}
}

func (c BackupFilesBodyFileSource) MarshalJSON() ([]byte, error) {
	return utils.Marshal(c.value)
}

func (c *BackupFilesBodyFileSource) UnmarshalJSON(b []byte) error {
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
